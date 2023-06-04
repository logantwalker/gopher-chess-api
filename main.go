package main

import (
	"log"
	"os"

	"github.com/logantwalker/gopher-chess-api/router"
)

func main(){
	r := router.InitRouter()

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
		logPID := os.Getenv("devLogPID")
		os.Setenv("logPID", logPID)
	} 

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	r.Run(":" + port)
}