package main

import (
	"oanda-mock-api/internal/config"
	"oanda-mock-api/internal/handlers"
	"oanda-mock-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	r := gin.Default()
	// r.SetTrustedProxies(strings.Split(config.GetEnv("TRUSTED_PROXIES"), ","))
	r.Use(middleware.AuthMiddleware())

	handlers.RegisterAccountHandlers(r.Group("/accounts"))

	r.Run()
}
