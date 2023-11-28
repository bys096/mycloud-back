package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"os"
)

func NewGoogleAuthConf() *oauth2.Config {
	//cxt := context.Background()

	//credentialsJSON := `{"web":{"client_id":"373576796271-tf4ub0vrmvsepj0pb2fr2vnkfpmcmnsk.apps.googleusercontent.com","project_id":"oauthapiproject-387705","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"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","redirect_uris":["http://localhost:8080"],"javascript_origins":["http://localhost:9000"]}}`
	path := "./credentials/auth_credential.json"
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	config, err := google.ConfigFromJSON(b, "email profile openid")
	if err != nil {
		panic(err)
	}
	log.Println(config)

	return nil
}
