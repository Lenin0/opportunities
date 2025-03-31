package main

import (
	"github.com/Lenin0/opportunities/config"
	"github.com/Lenin0/opportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	// Initializer Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization erro: %v", err)
		return
		// panic(err)
	}

	// Initializer Router
	router.Initialize()
}
