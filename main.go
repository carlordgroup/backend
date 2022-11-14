package main

import (
	"carlord/ent"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()
	db, _ := os.LookupEnv("DATABASE")
	client, err := ent.Open("postgres", db)
	if err != nil {
		panic(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	count, err := client.User.Query().Count(context.Background())
	if err != nil {
		panic(err)
	}

	print(count)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:8686")
}
