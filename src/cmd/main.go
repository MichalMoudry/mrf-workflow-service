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
	"firebase.google.com/go/v4/auth"
	dapr "github.com/dapr/go-sdk/client"
	"google.golang.org/api/option"
)

func main() {
	fmt.Println("Hello from workflow service! ʕ•ᴥ•ʔ")

	// Read app's config
	cfg, err := config.ReadCfgFromFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("App is running in '%v' mode.\n", cfg.Environment)

	// Connect to db
	if err = database.OpenDb(cfg.ConnectionString); err != nil {
		log.Fatal(err)
	}

	// Firebase init
	var firebaseAuth *auth.Client
	if cfg.RunWithFirebase {
		firebaseCredentials, err := config.CreateFirebaseCredentials()
		if err != nil {
			log.Fatal(err)
		}
		firebaseApp, err := firebase.NewApp(
			context.Background(),
			config.GetFirebaseConfig(),
			option.WithCredentialsJSON(firebaseCredentials),
		)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}
		firebaseAuth, err = firebaseApp.Auth(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}

	// Dapr client init
	var daprClient dapr.Client
	if cfg.RunWithDapr {
		daprClient, err = dapr.NewClient()
		if err != nil {
			log.Fatal(err)
		}
		defer daprClient.Close()
	}

	// Start web server
	fmt.Printf("Trying to start a server on %d port.\n", cfg.Port)
	handler := transport.Initalize(
		cfg.Port,
		model.NewServiceCollection(daprClient),
		firebaseAuth,
	)

	fmt.Printf("Listening on port: %d\n", handler.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", handler.Port), handler.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
