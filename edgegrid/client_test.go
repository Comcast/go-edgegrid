package edgegrid

import (
	"os"
	"testing"
)

func TestNewCredentials(t *testing.T) {
	os.Setenv("AKAMAI_EDGEGRID_HOST", "https://apiHost")
	os.Setenv("AKAMAI_EDGEGRID_ACCESS_TOKEN", "accessToken")
	os.Setenv("AKAMAI_EDGEGRID_CLIENT_TOKEN", "clientToken")
	os.Setenv("AKAMAI_EDGEGRID_CLIENT_SECRET", "clientSecret")

	creds := NewCredentials()

	if creds.ApiHost != "https://apiHost" {
		t.Error("ClientWithCredsFromEnv should set an ApiHost")
	}

	if creds.AccessToken != "accessToken" {
		t.Error("ClientWithCredsFromEnv should set an AccessToken")
	}

	if creds.ClientToken != "clientToken" {
		t.Error("ClientWithCredsFromEnv should set a ClientToken")
	}

	if creds.ClientSecret != "clientSecret" {
		t.Error("ClientWithCredsFromEnv should set a ClientSecret")
	}
}
