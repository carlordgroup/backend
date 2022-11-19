package auth

import (
	"carlord/data"
	"carlord/ent/account"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

type LoginCredential struct {
	Email       string `json:"email" binding:"required,email"`
	RawPassword string `json:"password" binding:"required"`
}

var identityKey = "id"

func (s *service) payload(data interface{}) jwt.MapClaims {
	if v, ok := data.(int); ok {
		return jwt.MapClaims{
			identityKey: v,
		}
	}
	return jwt.MapClaims{}
}

func (s *service) identity(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return int(claims[identityKey].(float64))
}

func (s *service) authentication(c *gin.Context) (interface{}, error) {
	var login LoginCredential
	if err := c.ShouldBindJSON(&login); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	acc, err := s.client.Account.Query().Where(account.Email(login.Email)).Only(c)
	if err != nil {
		return nil, err
	}
	accountObject := data.NewAccount(c, acc)

	if !accountObject.Verify(login.RawPassword) {
		return nil, jwt.ErrFailedAuthentication
	}
	c.Set("user", acc)
	return acc.ID, nil

}

func (s *service) initMiddleware() *jwt.GinJWTMiddleware {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "carlord",
		Key:             []byte(os.Getenv("SECRET")),
		Timeout:         time.Hour * 24,
		MaxRefresh:      time.Hour * 96,
		IdentityKey:     identityKey,
		IdentityHandler: s.identity,
		PayloadFunc:     s.payload,
		Authenticator:   s.authentication,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"error": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return authMiddleware

}

func (s *service) MustLogin() gin.HandlerFunc {
	return s.auth.MiddlewareFunc()
}
func (s *service) MustAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !ctx.MustGet("account").(*data.Account).Admin() {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "you are not an admin"})
			return
		}
		ctx.Next()
	}
}

func (s *service) GetAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		acc, err := s.client.Account.Query().WithUser().Where(account.ID(ctx.MustGet("id").(int))).Only(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		ctx.Set("account", data.NewAccount(ctx, acc))
		ctx.Next()
	}
}
