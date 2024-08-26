package fxmodule

import (
	"example.com/src/api/controllers"
	"example.com/src/api/routes"
	"example.com/src/api/server"
	"go.uber.org/fx"
)

var ApiModule = fx.Module(
	"Api Module",
	fx.Provide(server.ProvideGinServer),
	fx.Provide(controllers.NewController),
	fx.Invoke(server.ConfigureCors),
	fx.Invoke(routes.ConfigureRoutes),
	fx.Invoke(server.StartGinServer),
)
