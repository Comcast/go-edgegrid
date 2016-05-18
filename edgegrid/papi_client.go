package edgegrid

import "net/http"

type PAPIClient struct {
	Credentials *AuthCredentials
	HttpClient  *http.Client
}

func (c PAPIClient) GetCredentials() *AuthCredentials {
	return c.Credentials
}

func (c PAPIClient) GetHttpClient() *http.Client {
	return c.HttpClient
}
