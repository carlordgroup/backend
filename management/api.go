package management

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
	g := group.Group("/location")
	g.GET("/", web.W(s.listLocation))
	g.Use(auth.MustLogin(), auth.GetAccount(), auth.MustAdmin())
	g.POST("/:id", web.ID(s.updateLocation))
	g.POST("/", web.W(s.addLocation))
	g.DELETE("/:id", web.ID(s.deleteLocation))

	g = group.Group("/car")
	g.GET("/", web.W(s.filterCar))
	g.GET("/:id", web.ID(s.getCar))
	g.Use(auth.MustLogin(), auth.GetAccount(), auth.MustAdmin())
	g.POST("/", web.W(s.addCar))
	g.POST("/:id", web.ID(s.updateCar))
	g.DELETE("/:id", web.ID(s.deleteCar))

	g = group.Group("/offline")
	g.Use(auth.MustLogin(), auth.GetAccount(), auth.MustAdmin())
	g.POST("/user", web.W(s.createUser))
	g.POST("/card/:id", web.ID(s.createCard))
	g.POST("/book", web.W(s.bookCar))
	g.DELETE("/book/:id", web.ID(s.cancelBookingCar))

	g = group.Group("/booking")
	g.Use(auth.MustLogin(), auth.GetAccount(), auth.MustAdmin())
	g.GET("/", web.W(s.listBooking))
}
