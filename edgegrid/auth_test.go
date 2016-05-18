package edgegrid

import (
	"bytes"
	"net/http"
	"testing"
)

func newTestAuthParams() AuthParams {
	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)

	params := AuthParams{
		clientToken:  "clientToken",
		accessToken:  "accessToken",
		clientSecret: "clientSecret",
		nonce:        "uuid",
		timestamp:    "timestamp",
		req:          req,
	}

	return params
}

func TestNewAuthParams(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	params := NewAuthParams(req, "accessToken", "clientToken", "clientSecret")

	if params.req != req {
		t.Error("NewAuthParams failed: req")
	}

	if params.accessToken != "accessToken" {
		t.Error("NewAuthParams failed: accessToken")
	}

	if params.clientToken != "clientToken" {
		t.Error("NewAuthParams failed: clientToken")
	}

	if params.clientSecret != "clientSecret" {
		t.Error("NewAuthParams failed: clientSecret")
	}

	if params.timestamp == "" {
		t.Error("NewAuthParams failed: timestamp")
	}

	if params.nonce == "" {
		t.Error("NewAuthParams failed: nonce")
	}
}

func TestAuth(t *testing.T) {
	auth := Auth(newTestAuthParams())

	if auth != "EG1-HMAC-SHA256 client_token=clientToken;access_token=accessToken;timestamp=timestamp;nonce=uuid;signature=1Vqh8JzaFhBbyzPoyjZ5EygbH86MNJxAa6QlGAEY85A=" {
		t.Error("Auth failed")
	}
}

func TestSignRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	signed := signRequest(req, "timestamp", "secret", "auth", []string{})

	if signed != "signature=JYgCcBLZV63pALRV/zsufws39UPzIJHjPRUa0Uayi/k=" {
		t.Error("signRequest failed")
	}
}

func TestBase64Sha256(t *testing.T) {
	if base64Sha256("foo") != "LCa0a2j/xo/5m0U8HTBBNBNCLXBkg7+g+YpeiGJm564=" {
		t.Error("base64Sha256 should properly hash 'foo' string")
	}

	if base64Sha256("monarch") != "DyfVOp4Jx6moDmQCZl6cLbqaelHmTN/pKcYSE+OWn+Y=" {
		t.Error("base64Sha256 should properly hash 'monarch' string")
	}

	if base64Sha256("{}") != "RBNvo1WzZ4oRRq0W9+hknpT7T8If536DEMBg9hyq/4o=" {
		t.Error("base64Sha256 should properly hash a '{}' string")
	}

	/* TODO: this fails
	simpleJ := `{"foo:":"bar"}`

	if base64Sha256(simpleJ) != "eji/gfOD9pQzrW6QDTWz4jhVk/dqe3q11DVbi6Qe4ks=" {
		t.Error("base64Sha256 should properly hash a simple JSON string")
	}
	*/

	/* TODO: this fails
	j := `{
		"city": "Downpatrick",
		"continent": "EU",
		"country": "GB",
		"latitude": 54.367,
		"longitude": -5.582,
		"nickname": "somedc"
	}`

	if base64Sha256(str) != "az4Agj4KfiS/eH+vIB6pG8eTe2a2B6aTuiKSGP4bfbY=" {
		t.Error("base64Sha256 should properly hash a complex JSON string")
	}
	*/
}

func TestBase64HmacSha256(t *testing.T) {
	hash := base64HmacSha256("message", "secret")

	if hash != "i19IcCmVwVmMVz2x4hhmqbgl1KeU0WnXBgoDYFeWNgs=" {
		t.Error("base64HmacSha256 failed")
	}
}

func TestMakeDataToSign(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	data := makeDataToSign(req, "authHeader", []string{})

	if data != "GET\thttp\texample.com\t/foo\t\t\tauthHeader" {
		t.Error("makeDataToSign failed")
	}
}

func TestMakeDataToSignWithQuery(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/foo?baz=bim", nil)
	data := makeDataToSign(req, "authHeader", []string{})

	if data != "GET\thttp\texample.com\t/foo?baz=bim\t\t\tauthHeader" {
		t.Error("makeDataToSign failed for a request URL containing a query")
	}
}

func TestCanonicalizeHeaders(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)

	req.Header.Add("Foo", "bar")
	req.Header.Add("Baz", "   baz   zoo ")

	canonicalized := canonicalizeHeaders(req, []string{"foo", "baz"})

	if canonicalized != "foo:bar\tbaz:baz zoo\t" && canonicalized != "baz:baz zoo\tfoo:bar\t" {
		t.Error("canonicalized failed")
	}
}

func TestCanonicalizeHeadersMultipleValues(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)

	req.Header.Add("Foo", "bar")
	req.Header.Add("Foo", "biz")

	canonicalized := canonicalizeHeaders(req, []string{"foo"})

	if canonicalized != "foo:bar\t" {
		t.Error("canonicalized with multiple values failed")
	}
}

func TestCanonicalizeHeadersNoHeaders(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)

	canonicalized := canonicalizeHeaders(req, []string{"foo"})

	if canonicalized != "" {
		t.Error("canonicalized with no header values failed")
	}
}

func TestCanonicalizeHeadersNoIncludedHeaders(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)

	req.Header.Add("Fiz", "biz")

	canonicalized := canonicalizeHeaders(req, []string{"foo"})

	if canonicalized != "" {
		t.Error("canonicalized with no headers to include failed")
	}
}

func TestMakeContentHash(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	hash := makeContentHash(req)

	if hash != "" {
		t.Error("makeContentHash with GET request failed")
	}
}

func TestMakeContentHashPost(t *testing.T) {
	req, _ := http.NewRequest("POST", "http://example.com/foo", bytes.NewBufferString("foo"))
	hash := makeContentHash(req)

	if hash != "LCa0a2j/xo/5m0U8HTBBNBNCLXBkg7+g+YpeiGJm564=" {
		t.Error("makeContentHash with GET request failed")
	}
}

func TestMakeSigningKey(t *testing.T) {
	key := makeSigningKey("timestamp", "secret")

	if key != "ydMIxJIPPypuUya3KZGJ0qCRwkYcKrFn68Nyvpkf1WY=" {
		t.Error("makeSigningKey failed")
	}
}

func TestStringInSlice(t *testing.T) {
	result := stringInSlice("a", []string{"a", "b", "c"})
	resultTwo := stringInSlice("a", []string{"b", "c"})

	if result != true {
		t.Error("stringInSlice:true failed")
	}

	if resultTwo != false {
		t.Error("stringInSlice:false failed")
	}
}
