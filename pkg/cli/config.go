package cli

import "fmt"

type Config struct {
	KeycloakConfig KeycloakConfig
	EmbeddedServerConfig EmbeddedServerConfig
}

type KeycloakConfig struct {
	KeycloakURL string
	Realm string
	ClientID string
}

type EmbeddedServerConfig struct {
	Port         uint32
	CallbackPath string
}

func (c *EmbeddedServerConfig) GetCallbackURL() string {
	return fmt.Sprintf("http://localhost:%v/%v", c.Port, c.CallbackPath)
}
