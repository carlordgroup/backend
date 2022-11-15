package auth

import (
	"carlord/ent/account"
	"encoding/base64"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
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

	bcryptPass, err := base64.StdEncoding.DecodeString(acc.Password)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(bcryptPass, []byte(login.RawPassword))
	if err != nil {
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
