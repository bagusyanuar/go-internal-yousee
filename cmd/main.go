package main

import (
	"log"

	"github.com/bagusyanuar/go-internal-yousee/internal/config"
)

func main() {
	app := config.NewFiber()
	config.Bootstrap(&config.BootstrapConfig{
		App: app,
	})
	err := app.Listen(":8000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
