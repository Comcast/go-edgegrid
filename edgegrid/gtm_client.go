package edgegrid

import "net/http"

type GTMClient struct {
	Credentials *AuthCredentials
	HttpClient  *http.Client
}

func (c GTMClient) GetCredentials() *AuthCredentials {
	return c.Credentials
}

func (c GTMClient) GetHttpClient() *http.Client {
	return c.HttpClient
}
