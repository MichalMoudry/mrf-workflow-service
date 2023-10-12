package config

import (
	"encoding/json"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
)

// A structure for encapsulating credentials for Firebase.
type GoogleCredentials struct {
	Type                string `json:"type"`
	ProjectId           string `json:"project_id"`
	PrivateKeyId        string `json:"private_key_id"`
	PrivateKey          string `json:"private_key"`
	ClientEmail         string `json:"client_email"`
	ClientId            string `json:"client_id"`
	AuthUri             string `json:"auth_uri"`
	TokenUri            string `json:"token_uri"`
	AuthProviderCertUrl string `json:"auth_provider_x509_cert_url"`
	ClientCertUrl       string `json:"client_x509_cert_url"`
	UniverseDomain      string `json:"universe_domain"`
}

// Function for obtaining Firebases configuration.
func GetFirebaseConfig() *firebase.Config {
	return &firebase.Config{
		ProjectID: "ocr-microservice-project",
	}
}

// Function for creating Firebase credentials based on environment values.
func CreateFirebaseCredentials() ([]byte, error) {
	return json.Marshal(GoogleCredentials{
		Type:                os.Getenv("FB_TYPE"),
		ProjectId:           "ocr-microservice-project",
		PrivateKeyId:        os.Getenv("FB_PRIV_KEY_ID"),
		PrivateKey:          strings.ReplaceAll(os.Getenv("FB_PRIV_KEY"), "\\n", "\n"),
		ClientEmail:         os.Getenv("FB_CLIENT_EMAIL"),
		ClientId:            os.Getenv("FB_CLIENT_ID"),
		AuthUri:             "https://accounts.google.com/o/oauth2/auth",
		TokenUri:            "https://oauth2.googleapis.com/token",
		AuthProviderCertUrl: "https://www.googleapis.com/oauth2/v1/certs",
		ClientCertUrl:       os.Getenv("FB_CLIENT_CERT_URL"),
		UniverseDomain:      "googleapis.com",
	})
}
