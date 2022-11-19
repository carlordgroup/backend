package booking

import (
	"carlord/ent"
	"carlord/web"
	"github.com/gin-gonic/gin"
)

type service struct {
	client *ent.Client
}

type Authenticate interface {
	MustLogin() gin.HandlerFunc
}

func New(client *ent.Client) *service {
	s := service{
		client: client,
	}
	return &s
}

func (s *service) RegisterRouter(group gin.IRouter, auth Authenticate) {
	group.Use(auth.MustLogin())
	group.GET("/", web.W(s.listBooking))
	group.POST("/", web.W(s.bookCar))
	group.DELETE("/:id", web.ID(s.cancelBookingCar))
}
