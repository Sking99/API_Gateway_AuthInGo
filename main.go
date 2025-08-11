package main

import (
	"AuthInGo/app"
	config "AuthInGo/config/env"
)

func main() {
	config.Load()
	cfg := app.NewConfig()
	app := app.NewApplication(cfg)

	app.Run()
}
