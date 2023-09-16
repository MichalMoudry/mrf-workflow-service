package main

import (
	"fmt"
	"log"
	"net/http"
	"workflow-service/config"
	"workflow-service/transport"
	"workflow-service/transport/model"
)

func main() {
	fmt.Println("Hello from workflow service!")
	cfg, err := config.ReadCfgFromFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Trying to start a server on %d port.\n", cfg.Port)
	handler := transport.Initalize(
		cfg.Port,
		model.NewServiceCollection(),
	)

	err = http.ListenAndServe(fmt.Sprintf(":%d", handler.Port), handler.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
