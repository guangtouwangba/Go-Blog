package controller

import (
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (b *BaseController) Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
