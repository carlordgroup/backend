package main

import (
	"carlord/auth"
	"carlord/docs"
	"carlord/ent"
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
	auth.New(client, g)
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/api/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run("0.0.0.0:8686")
}
