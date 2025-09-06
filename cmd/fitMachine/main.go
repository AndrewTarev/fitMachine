package main

import (
	"context"
	"log"

	appconfig "fitMachine/pkg/config"
	"fitMachine/pkg/logger"
	"fitMachine/pkg/server"
)

func main() {
	ctx := context.Background()

	cfg, err := appconfig.New()
	if err != nil {
		log.Fatalf("Failed to create config: %v", err)
	}

	appLogger := logger.New(cfg)

	httpServer := server.New(ctx, cfg, appLogger)
	httpServer.SetupRoutes()

	if err = httpServer.Run(); err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}
}
