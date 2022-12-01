package pickup

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
	group.Use(auth.MustLogin(), auth.GetAccount(), auth.MustAdmin())
	group.POST("/start/:id", web.ID(s.realizeBooking))
	group.POST("/finish/:id", web.ID(s.returnCar))
}
