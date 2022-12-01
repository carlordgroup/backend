package card

import (
	"carlord/ent"
	"carlord/web"
	"github.com/gin-gonic/gin"
)

type service struct {
	client *ent.Client
}

func New(client *ent.Client) *service {
	s := service{
		client: client,
	}
	return &s
}

func (s *service) RegisterRouter(group gin.IRouter, auth web.Authenticate) {
	group.Use(auth.MustLogin(), auth.GetAccount())
	group.GET("/", web.W(s.get))
	group.POST("/:id", web.W(s.post))
	group.POST("/", web.W(s.create))
	group.DELETE("/:id", web.W(s.delete))
}
