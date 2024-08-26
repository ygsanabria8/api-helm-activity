package main

import (
	"example.com/src/fxmodule"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fxmodule.ApiModule,
	).Run()
}
