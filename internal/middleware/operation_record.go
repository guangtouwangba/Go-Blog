package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func RequestLogger(c *gin.Context) {
	log.Println("current request is: ", c.Request.URL.Path)
	log.Println("current request method is: ", c.Request.Method)
	log.Println("current request body is", c.Request.Body)
}
