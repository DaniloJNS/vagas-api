package main

import (
	"github.com/DaniloJNS/vagas-api/config"
	"github.com/DaniloJNS/vagas-api/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	// Initialize configs
	err := config.Init()

	if err != nil {
	  logger.Errorf("config initialization error: %v", err)
		return
	}
	// Start web server
	router.Initializer()
}
