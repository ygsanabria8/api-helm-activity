package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
)

func StartGinServer(
	lifecycle fx.Lifecycle, server *gin.Engine,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				server.Use(gin.Recovery())
				err := server.Run("0.0.0.0:11000")
				if err != nil {
					log.Fatalf("error starting server: %s", err.Error())
				}
				log.Printf("Server started on port 11000")
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Printf("Shutting down server")
			return nil
		},
	})
}

func ProvideGinServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}
