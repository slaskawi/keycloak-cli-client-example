package main

import (
	"github.com/slaskawi/keycloak-cli-client-example/pkg/cli"
	"log"
)

func main() {
	cli.CloseApp.Add(1)
	config := cli.Config{
		KeycloakConfig:       cli.KeycloakConfig{
			KeycloakURL: "http://localhost:8080",
			Realm:       "master",
			ClientID:    "cli-example",
		},
		EmbeddedServerConfig: cli.EmbeddedServerConfig{
			Port:         8081,
			CallbackPath: "sso-callback",
		},
	}

	cli.StartServer(config)
	err := cli.OpenBrowser(cli.BuildAuthorizationRequest(config))
	if err != nil {
		log.Fatalf("Could not open the browser for url %v", cli.BuildAuthorizationRequest(config))
	}

	cli.CloseApp.Wait()
}