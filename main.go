package main

import (
	"carlord/auth"
	"carlord/booking"
	"carlord/card"
	"carlord/docs"
	"carlord/ent"
	"carlord/management"
	"carlord/pickup"
	"carlord/user"
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
)

func main() {
	db, _ := os.LookupEnv("DATABASE")
	client, err := ent.Open("postgres", db)
	if err != nil {
		panic(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	r := gin.Default()
	{
		c := cors.DefaultConfig()
		c.AllowHeaders = append(c.AllowHeaders, "Authorization")
		c.AllowAllOrigins = true
		r.Use(cors.New(c))
	}

	api := r.Group("api/")
	api.GET("", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/api/swagger/index.html")
	})

	g := api.Group("account/")
	authService := auth.New(client)
	authService.RegisterRouter(g)

	g = api.Group("user/")
	user.New(client).RegisterRouter(g, authService)

	g = api.Group("card/")
	card.New(client).RegisterRouter(g, authService)

	g = api.Group("management/")
	management.New(client).RegisterRouter(g, authService)

	g = api.Group("booking/")
	booking.New(client).RegisterRouter(g, authService)

	g = api.Group("pickup/")
	pickup.New(client).RegisterRouter(g, authService)

	docs.SwaggerInfo.BasePath = "/api"
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run("0.0.0.0:8686")

}
