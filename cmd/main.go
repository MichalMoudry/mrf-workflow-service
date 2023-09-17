package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"workflow-service/config"
	"workflow-service/transport"
	"workflow-service/transport/model"

	firebase "firebase.google.com/go/v4"
	dapr "github.com/dapr/go-sdk/client"
)

func main() {
	fmt.Println("Hello from workflow service!")
	// Read app's config
	cfg, err := config.ReadCfgFromFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	// Dapr client init
	daprClient, err := dapr.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer daprClient.Close()

	// Firebase init
	_, err = firebase.NewApp(context.Background(), config.GetFirebaseConfig())
	if err != nil {
		log.Fatalf("error initializing Firebase app: %v\n", err)
	}

	// Start web server
	fmt.Printf("Trying to start a server on %d port.\n", cfg.Port)
	handler := transport.Initalize(
		cfg.Port,
		model.NewServiceCollection(),
	)
	err = http.ListenAndServe(fmt.Sprintf(":%d", handler.Port), handler.Mux)
	fmt.Printf("Listening on port: %d\n", cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}
