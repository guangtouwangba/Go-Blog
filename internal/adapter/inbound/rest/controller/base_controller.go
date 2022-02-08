package controller

import (
	"Go-Blog/internal/domain/dto/response"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (b *BaseController) Health(c *gin.Context) {
	r := response.Response{}
	r.Success()
	c.JSON(200, r)
}
