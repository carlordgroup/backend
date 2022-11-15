package auth

import (
	"carlord/ent"
	"encoding/base64"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type service struct {
	client *ent.Client
}

func New(client *ent.Client, group gin.IRoutes) *jwt.GinJWTMiddleware {
	s := service{
		client: client,
	}
	auth := s.initMiddleware()
	group.POST("/login", login(auth))
	group.POST("/refresh", auth.RefreshHandler)
	group.POST("/register", s.register)
	group.GET("/", auth.MiddlewareFunc(), s.self)
	return auth
}

type token struct{ Token string }

// login godoc
// @Summary login
// @Description login an account
// @Param request body LoginCredential true "email password"
// @Accept json
// @Produce json
// @Success 200 {object} token
// @Router /account/login [get]
func login(auth *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return auth.LoginHandler
}

// Self godoc
// @Summary get login user info
// @Description will respond user account info when user login
// @Accept json
// @Produce json
// @Success 200 {object} ent.Account
// @Router /account/ [get]
func (s *service) self(ctx *gin.Context) {
	u, err := s.client.Account.Get(ctx, ctx.MustGet("id").(int))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, u)

}

// Register godoc
// @Param request body LoginCredential true "email password"
// @Summary register a user
// @Accept json
// @Produce json
// @Success 200 {object} ent.Account
// @Router /account/register [post]
func (s *service) register(ctx *gin.Context) {
	var login LoginCredential
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(login.RawPassword), bcrypt.DefaultCost)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	u, err := s.client.Account.Create().
		SetEmail(login.Email).
		SetPassword(base64.StdEncoding.EncodeToString(hashedPassword)).
		Save(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, u)
}
