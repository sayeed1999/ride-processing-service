package main

import (
	"fmt"
	"ride-processing-service/config"
	transition_checker "ride-processing-service/internal/features/transition-checker"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	run(r)
}

func run(r *gin.Engine) {
	config := config.LoadConfig()
	serverHost := config.Server.Host
	serverPort := config.Server.Port

	transition_checker.InitEndpoints(r)

	// Attempt 1: -
	// Dockerized the app running on http://localhost:8080.
	// The endpoints are not accessible by the host OS saying,
	// 'Connection is refused by the server'.
	// Attempt 2: -
	// Dockerized the app running on http://0.0.0.0:8080.
	// It works...

	// Ref: https://stackoverflow.com/a/71980210

	addr := fmt.Sprintf("%s:%s", serverHost, serverPort)
	r.Run(addr)
}
