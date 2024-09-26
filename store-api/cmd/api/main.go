package main

import (
	"log/slog"

	"github.com/Arturlima/store-api/providers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}
	slog.Info("all system offline")
}

func run() error {
	slog.Info("Starting providers")

	provider := providers.NewProvider()

	g := gin.New()

	slog.Info("Including middlewares")
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	g.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	provider.ScopedStoreController().RegisterRoutes(g)

	slog.Info("HTTP server running port 8081")
	g.Run(":8081")

	return nil
}
