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
	jwt := config.NewJWT(viperConfig)
	validator := config.NewValidator()

	config.Bootstrap(&config.BootstrapConfig{
		App:       app,
		Config:    viperConfig,
		Log:       log,
		DB:        db,
		JWT:       jwt,
		Validator: validator,
	})

	port := viperConfig.GetString("APP_PORT")
	err := app.Listen(fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
