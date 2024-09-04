package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (c *Controller) Greeting(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"greeting": "Hello From ArgoCD!"})
}

func (c *Controller) Bye(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"bye": "Bye From ArgoCD"})
}
