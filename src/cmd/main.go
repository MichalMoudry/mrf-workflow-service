package main

import (
	"fmt"
	"log"
	"net/http"
	"workflow-service/config"
	"workflow-service/database"
	"workflow-service/transport"
	"workflow-service/transport/model"
)

func main() {
	fmt.Println("Hello from workflow service!")
	// Read app's config
	cfg, err := config.ReadCfgFromFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

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

	// Start web server
	fmt.Printf("Trying to start a server on %d port.\n", cfg.Port)
	handler := transport.Initalize(
		cfg.Port,
		model.NewServiceCollection(),
	)
	fmt.Printf("Listening on port: %d\n", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", handler.Port), handler.Mux)
	if err != nil {
		log.Fatal(err)
	}
}