package controllers

import (
	"github.com/gin-gonic/gin"
)

type IController interface {
	Health(ctx *gin.Context)
	Greeting(ctx *gin.Context)
	Bye(ctx *gin.Context)
}

type Controller struct {
}

func NewController() IController {
	return &Controller{}
}
