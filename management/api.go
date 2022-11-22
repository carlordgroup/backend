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
	MustLogin() gin.HandlerFunc
	GetAccount() gin.HandlerFunc
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
	g.Use(auth.MustLogin(), auth.GetAccount(), auth.MustAdmin())
	g.POST("/:id", web.ID(s.updateLocation))
	g.POST("/", web.W(s.addLocation))
	g.DELETE("/:id", web.ID(s.deleteLocation))

	g = group.Group("/car")
	g.GET("/", web.W(s.filterCar))
	g.Use(auth.MustLogin(), auth.GetAccount(), auth.MustAdmin())
	g.POST("/", web.W(s.addCar))
	g.POST("/:id", web.ID(s.updateCar))
	g.DELETE("/:id", web.ID(s.deleteCar))

	g.Use(auth.MustLogin(), auth.GetAccount(), auth.MustAdmin())
	g = group.Group("/offline")
	g.POST("/user", web.W(s.createUser))
	g.POST("/card/:id", web.ID(s.createCard))
	g.POST("/book", web.W(s.bookCar))
	g.DELETE("/book/:id", web.ID(s.cancelBookingCar))

}
