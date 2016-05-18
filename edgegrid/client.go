package edgegrid

import (
	"net/http"
	"os"
)

type Client interface {
	GetCredentials() *AuthCredentials
	GetHttpClient() *http.Client
}

type AuthCredentials struct {
	AccessToken  string
	ClientToken  string
	ClientSecret string
	ApiHost      string
}

func GTMClientWithCreds(accessToken, clientToken, clientSecret, apiHost string) *GTMClient {
	return &GTMClient{
		&AuthCredentials{accessToken, clientToken, clientSecret, apiHost},
		&http.Client{},
	}
}

func NewGTMClient() *GTMClient {
	return &GTMClient{
		NewCredentials(),
		&http.Client{},
	}
}

func NewPAPIClient() *PAPIClient {
	return &PAPIClient{
		NewCredentials(),
		&http.Client{},
	}
}

func PAPIClientWithCreds(accessToken, clientToken, clientSecret, apiHost string) *PAPIClient {
	return &PAPIClient{
		&AuthCredentials{accessToken, clientToken, clientSecret, apiHost},
		&http.Client{},
	}
}

func NewCredentials() *AuthCredentials {
	return &AuthCredentials{
		os.Getenv("AKAMAI_EDGEGRID_ACCESS_TOKEN"),
		os.Getenv("AKAMAI_EDGEGRID_CLIENT_TOKEN"),
		os.Getenv("AKAMAI_EDGEGRID_CLIENT_SECRET"),
		os.Getenv("AKAMAI_EDGEGRID_HOST"),
	}
}

func LogRequests() bool {
	return os.Getenv("AK_LOG") != ""
}
