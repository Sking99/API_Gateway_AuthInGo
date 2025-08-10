package main

import (
	"AuthInGo/app"
	"AuthInGo/config"
)

func main() {
	config.Load()
	cfg := app.NewConfig()
	app := app.NewApplication(cfg)

	app.Run()
}
