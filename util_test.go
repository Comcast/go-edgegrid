package edgegrid

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type testClient struct {
	Credentials *AuthCredentials
	HttpClient  *http.Client
}

func (c testClient) GetCredentials() *AuthCredentials {
	return c.Credentials
}

func (c testClient) GetHttpClient() *http.Client {
	return c.HttpClient
}

func utilTestTools(code int, body string) (*httptest.Server, *testClient) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, body)
	}))

	tr := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := &http.Client{Transport: tr}

	return server, &testClient{
		&AuthCredentials{
			"accessToken",
			"clientToken",
			"clientSecret",
			server.URL,
		},
		httpClient,
	}
}

func TestAuthenticate(t *testing.T) {
	body := []byte("hello world")
	server, client := utilTestTools(200, "hello world")
	defer server.Close()

	req, _ := http.NewRequest("GET", client.GetCredentials().ApiHost, bytes.NewBuffer(body))

	authenticatedReq := authenticate(client, req)
	header := authenticatedReq.Header.Get("Authorization")

	if header == "" {
		t.Error("authenticate should set an Authorization header on the request")
	}
}

func TestConcat(t *testing.T) {
	res := concat([]string{"I ", "am ", "a ", "sentence"})

	if res != "I am a sentence" {
		t.Error("concat failed")
	}
}

func TestUrlPathWithQuery(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://foo.com/biz/bar?someKey=someVal", nil)

	if urlPathWithQuery(req) != "/biz/bar?someKey=someVal" {
		t.Error("urlPathWithQuery should return the request URL path & query")
	}
}

func TestDoClientReq(t *testing.T) {
	server, client := utilTestTools(200, "hello world")
	defer server.Close()

	resp, err := doClientReq(client, "GET", client.GetCredentials().ApiHost, nil)
	if err != nil {
		t.Error("test doClientReq failed")
	}

	body, _ := ioutil.ReadAll(resp.Body)

	if string(body) != "hello world" {
		t.Error("test doClientReq failed")
	}
}
