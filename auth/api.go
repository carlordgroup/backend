package auth

import (
	"carlord/ent"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type service struct {
	client *ent.Client
}

func New(client *ent.Client, group *gin.Engine) {
	s := service{
		client: client,
	}
	auth := s.initMiddleware()
	group.POST("/login", auth.LoginHandler)
	group.POST("/refresh", auth.RefreshHandler)
	group.POST("/register", s.register)
	group.GET("/", auth.MiddlewareFunc(), s.self)

}

func (s *service) self(ctx *gin.Context) {
	u, err := s.client.Account.Get(ctx, ctx.MustGet("id").(int))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, u)

}

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
