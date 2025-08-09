package main

import (
	"AuthInGo/app"
)

func main() {
	cfg := app.NewConfig(":3003")
	app := app.NewApplication(cfg)

	app.Run()
}
