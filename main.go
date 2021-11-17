package main

import (
	"log"
	"shotnr/handler"
	"shotnr/store"

	"github.com/gin-gonic/gin"
)

func main() {
	_ = store.InitStore()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome",
		})
	})
	r.POST("/create_short_url", func(c *gin.Context) {
		c.Set("user_id", "moyinka")
		handler.CreateShortUrl(c)
	})
	r.GET("/:short_url", func(c *gin.Context) {
		handler.RedirectUrl(c)
	})

	log.Fatal(r.Run(":8080"))
}
