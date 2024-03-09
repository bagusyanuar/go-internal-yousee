package main

import (
	"fmt"

	"github.com/bagusyanuar/go-internal-yousee/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger()
	app := config.NewFiber()
	db := config.NewDatabase(viperConfig, log)
	config.Bootstrap(&config.BootstrapConfig{
		App:    app,
		Config: viperConfig,
		Log:    log,
		DB:     db,
	})

	port := viperConfig.GetString("APP_PORT")
	err := app.Listen(fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
