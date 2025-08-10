package app

import (
	"AuthInGo/config"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}

type Application struct {
	Config Config
}

func NewConfig() Config {

	port := config.GetString("PORT", ":8080")

	return Config{
		Addr: port,
	}
}

func NewApplication(config Config) *Application {
	return &Application{
		Config: config,
	}
}

func (app *Application) Run() error {
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server running on port", app.Config.Addr)

	return server.ListenAndServe()
}
