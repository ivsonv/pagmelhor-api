package main

import (
	"app/configs"
	"app/modules/club"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg, _ := configs.LoadConfig(".")

	// Initialize echo server
	e := echo.New()

	// Start modules
	club.Start(e.Group("v1/club"), cfg)

	// Start the server
	go func() {
		if err := e.Start(":" + cfg.API_PORT); err != nil {
			log.Printf("Error starting server: %v", err)
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Printf("Error during server shutdown: %v", err)
	}
}
