package management

import (
	"carlord/ent"
	"carlord/web"
	"github.com/gin-gonic/gin"
)

type service struct {
	client *ent.Client
}

type Authenticate interface {
	MustAdmin() gin.HandlerFunc
}

func New(client *ent.Client) *service {
	s := service{
		client: client,
	}
	return &s
}

func (s *service) RegisterRouter(group gin.IRouter, auth Authenticate) {
	g := group.Group("/location")
	g.GET("/", web.W(s.listLocation))
	group.Use(auth.MustAdmin())
	group.POST("/", web.W(s.addLocation))
	group.POST("/:id", web.ID(s.updateLocation))
	group.DELETE("/:id", web.ID(s.deleteLocation))

	g = group.Group("/car")
	g.GET("/", web.W(s.filterCar))
	group.Use(auth.MustAdmin())
	group.POST("/", web.W(s.addCar))
	group.POST("/:id", web.ID(s.updateCar))
	group.DELETE("/:id", web.ID(s.deleteCar))

}
