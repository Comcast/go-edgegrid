package edgegrid

import "net/http"

// PAPIClient is an Akamai PAPI API client
// https://developer.akamai.com/api/luna/papi/overview.html
type PAPIClient struct {
	Credentials *AuthCredentials
	HttpClient  *http.Client
}

// GetCredentials takes a PAPIClient and returns its credentials
func (c PAPIClient) GetCredentials() *AuthCredentials {
	return c.Credentials
}

// GetHttpClient takes a PAPIClient and returns its HttpClient
func (c PAPIClient) GetHttpClient() *http.Client {
	return c.HttpClient
}
