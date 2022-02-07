package routing

import (
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter() *gin.Engine {
	var router = gin.Default()
	log.Println("Router initialized")
	router.GET("/health-check", func(c *gin.Context) {
		log.Println("health-check")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return router
}
