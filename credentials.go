package gotelemetry

import (
	"net/url"
)

type Credentials struct {
	APIKey    string
	ServerURL *url.URL
}

func NewCredentials(apiKey string, serverUrl ...string) (Credentials, error) {
	server := "https://api.telemetryapp.com"

	if len(serverUrl) > 0 {
		server = serverUrl[0]
	}

	url, err := url.Parse(server)

	return Credentials{apiKey, url}, err
}
