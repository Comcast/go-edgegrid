package edgegrid

import "net/http"

// GTMClient is an Akamai GTM API client
// https://developer.akamai.com/api/luna/config-gtm/overview.html
type GTMClient struct {
	Credentials *AuthCredentials
	HttpClient  *http.Client
}

// GetCredentials takes a GTMClient and returns its Credentials
func (c GTMClient) GetCredentials() *AuthCredentials {
	return c.Credentials
}

// GetHttpClient takes a GTMClient and returns its HttpClient
func (c GTMClient) GetHttpClient() *http.Client {
	return c.HttpClient
}
