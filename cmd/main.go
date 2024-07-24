package main

import (
	"log"
	netHttp "net/http"

	"github.com/azizyogo/bank-app/server"
	"github.com/azizyogo/bank-app/server/http"
)

func main() {
	// Initialize the server and defer the closing of resources
	err := server.Init()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}
	defer server.Close()

	// Create the router
	router := http.NewServer()

	// Start the HTTP server
	log.Println("Server starting at :8080")
	log.Fatal(netHttp.ListenAndServe(":8080", router))
}
