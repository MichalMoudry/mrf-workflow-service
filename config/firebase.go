package config

import firebase "firebase.google.com/go/v4"

// A function to construct Firebase app configuration.
func GetFirebaseConfig() *firebase.Config {
	return &firebase.Config{
		ProjectID: "ocr-microservice-project",
	}
}
