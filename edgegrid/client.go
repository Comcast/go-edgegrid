package edgegrid

import (
	"net/http"
	"os"
)

// Client is an interface for an Akamai API client.
type Client interface {
	GetCredentials() *AuthCredentials
	GetHTTPClient() *http.Client
}

// AuthCredentials houses various Akamai-client-specific
// data necessary to authenticate the Akamai API.
type AuthCredentials struct {
	AccessToken  string
	ClientToken  string
	ClientSecret string
	APIHost      string
}

// GTMClientWithCreds takes an access token string, a client token string,
// a client secret string, and an API host string and returns a GTMClient.
func GTMClientWithCreds(accessToken, clientToken, clientSecret, apiHost string) *GTMClient {
	return &GTMClient{
		&AuthCredentials{accessToken, clientToken, clientSecret, apiHost},
		&http.Client{},
	}
}

// NewGTMClient is a convenience method that returns a GTMClient
// a client secret string, and an API host string and returns a GTMClient
// using the AKAMAI_EDGEGRID_ACCESS_TOKEN, AKAMAI_EDGEGRID_CLIENT_TOKEN,
// AKAMAI_EDGEGRID_CLIENT_SECRET, and AKAMAI_EDGEGRID_HOST environment
// variables.
func NewGTMClient() *GTMClient {
	return &GTMClient{
		NewCredentials(),
		&http.Client{},
	}
}

// NewPAPIClient is a convenience method that returns a GTMClient
// a client secret string, and an API host string and returns a GTMClient
// using the AKAMAI_EDGEGRID_ACCESS_TOKEN, AKAMAI_EDGEGRID_CLIENT_TOKEN,
// AKAMAI_EDGEGRID_CLIENT_SECRET, and AKAMAI_EDGEGRID_HOST environment
// variables.
func NewPAPIClient() *PAPIClient {
	return &PAPIClient{
		NewCredentials(),
		&http.Client{},
	}
}

// PAPIClientWithCreds takes an access token string, a client token string,
// a client secret string, and an API host string and returns a PAPIClient.
func PAPIClientWithCreds(accessToken, clientToken, clientSecret, apiHost string) *PAPIClient {
	return &PAPIClient{
		&AuthCredentials{accessToken, clientToken, clientSecret, apiHost},
		&http.Client{},
	}
}

// NewCredentials is a convenience function that returns an &AuthCredentials{}
// using the AKAMAI_EDGEGRID_ACCESS_TOKEN, AKAMAI_EDGEGRID_CLIENT_TOKEN,
// AKAMAI_EDGEGRID_CLIENT_SECRET, and AKAMAI_EDGEGRID_HOST environment
// variables.
func NewCredentials() *AuthCredentials {
	return &AuthCredentials{
		os.Getenv("AKAMAI_EDGEGRID_ACCESS_TOKEN"),
		os.Getenv("AKAMAI_EDGEGRID_CLIENT_TOKEN"),
		os.Getenv("AKAMAI_EDGEGRID_CLIENT_SECRET"),
		os.Getenv("AKAMAI_EDGEGRID_HOST"),
	}
}

// LogRequests returns true if the AK_LOG environment variable is set;
// false if it is not.
func LogRequests() bool {
	return os.Getenv("AK_LOG") != ""
}
