package auth

import (
	"carlord/ent"
	"carlord/web"
	"encoding/base64"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type service struct {
	client *ent.Client
	auth   *jwt.GinJWTMiddleware
}

func New(client *ent.Client) *service {
	s := service{
		client: client,
	}
	auth := s.initMiddleware()
	s.auth = auth

	return &s
}

func (s *service) RegisterRouter(group gin.IRouter) {
	group.POST("/login", login(s.auth))
	group.POST("/refresh", s.auth.RefreshHandler)
	group.POST("/register", web.W(s.register))
	group.GET("/promote/:id", web.ID(s.promote))
	group.GET("/", s.MustLogin(), s.GetAccount(), web.W(s.self))
}

type tokenSample struct{ Token string }

// login godoc
// @Tags auth
// @Summary login
// @Description login an account
// @Param request body LoginCredential true "email password"
// @Accept json
// @Produce json
// @Success 200 {object} tokenSample
// @Router /account/login [get]
func login(auth *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return auth.LoginHandler
}

// Self godoc
// @Tags auth
// @Summary get login user info
// @Description will respond user account info when user login
// @Accept json
// @Produce json
// @Success 200 {object} ent.Account
// @Router /account/ [post]
func (s *service) self(ctx *gin.Context) (int, any) {
	return http.StatusOK, ctx.MustGet("account")
}

// Register godoc
// @Tags auth
// @Param request body LoginCredential true "email password"
// @Summary register a user
// @Accept json
// @Produce json
// @Success 200 {object} ent.Account
// @Router /account/register [post]
func (s *service) register(ctx *gin.Context) (int, any) {
	var login LoginCredential
	// data binding
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		return http.StatusBadRequest, err
	}
	// generate encrypted password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(login.RawPassword), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusBadRequest, err
	}
	// start transaction so it can be rollback if accident happened
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	// create account record
	a, err := tx.Account.Create().
		SetEmail(login.Email).
		SetPassword(base64.StdEncoding.EncodeToString(hashedPassword)).
		Save(ctx)
	if err != nil {
		tx.Rollback()

		return http.StatusBadRequest, err
	}
	// create user record
	_, err = tx.User.Create().SetAccountID(a.ID).Save(ctx)
	if err != nil {
		tx.Rollback()
		return http.StatusBadRequest, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return http.StatusBadRequest, err
	}

	return http.StatusCreated, a
}

// Register godoc
// @Tags auth
// @Summary test api: promote
// @Accept json
// @Produce json
// @Success 200 {object} ent.Account
// @Router /account/promote/:id [get]
func (s *service) promote(ctx *gin.Context, id int) (int, any) {
	// This is a test API, for the developers getting an admin account
	data, err := s.client.Account.UpdateOneID(id).SetIsAdmin(true).Save(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, data
}
