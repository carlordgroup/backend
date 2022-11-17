package main

import (
	"carlord/auth"
	"carlord/card"
	"carlord/docs"
	"carlord/ent"
	"carlord/management"
	"carlord/user"
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
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

	g := r.Group("account/")
	authService := auth.New(client)
	authService.RegisterRouter(g)

	g = r.Group("user/")
	user.New(client).RegisterRouter(g, authService)

	g = r.Group("card/")
	card.New(client).RegisterRouter(g, authService)

	g = r.Group("management/")
	management.New(client).RegisterRouter(g, authService)

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/api/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run("0.0.0.0:8686")
}
