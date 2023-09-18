package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"workflow-service/config"
	"workflow-service/database"
	"workflow-service/transport"
	"workflow-service/transport/model"

	firebase "firebase.google.com/go/v4"
)

func main() {
	fmt.Println("Hello from workflow service!")
	// Read app's config
	cfg, err := config.ReadCfgFromFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(
		"Trying to connect to the database.\nConnection string: %s\n",
		cfg.ConnectionString,
	)
	if err = database.OpenDb(cfg.ConnectionString); err != nil {
		log.Fatal(err)
	}

	// Dapr client init
	/*daprClient, err := dapr.NewClient()
	if err != nil {
		log.Println(err)
		daprClient.Close()
	}
	defer daprClient.Close()*/

	// Firebase init
	firebaseApp, err := firebase.NewApp(context.Background(), config.GetFirebaseConfig())
	if err != nil {
		log.Fatalf("error initializing Firebase app:\n%v\n", err)
	}
	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Printf("error initializing auth client:\n%v\n", err) // TODO: Switch to Fatalf
	}

	// Start web server
	fmt.Printf("Trying to start a server on %d port.\n", cfg.Port)
	handler := transport.Initalize(
		cfg.Port,
		model.NewServiceCollection(authClient),
	)
	fmt.Printf("Listening on port: %d\n", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", handler.Port), handler.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
