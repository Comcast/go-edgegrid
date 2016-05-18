package edgegrid

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gobs/pretty"
)

func gtmTestTools(code int, body string) (*httptest.Server, *GTMClient) {
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

	creds := &AuthCredentials{
		"accessToken",
		"clientToken",
		"clientSecret",
		server.URL,
	}

	client := &GTMClient{creds, httpClient}

	return server, client
}

func TestDomains(t *testing.T) {
	server, client := gtmTestTools(200, domainsJson)
	defer server.Close()

	domains, _ := client.Domains()
	t.Log(pretty.PrettyFormat(domains))

	if len(domains) != 2 {
		t.Error("Expected 2 Domains")
	}
	for _, d := range domains {
		if d.Name == "" {
			t.Error("Expected Name to have a value")
		}
		if d.Status == "" {
			t.Error("Expected Status to have a value")
		}
		if d.LastModified == "" {
			t.Error("Expected LastModified to have a value")
		}
	}
}

func TestDomain(t *testing.T) {
	server, client := gtmTestTools(200, domainJson)
	defer server.Close()

	domain, _ := client.Domain("example.akadns.net")
	t.Log(pretty.PrettyFormat(domain))

	if "example.akadns.net" != domain.Name {
		t.Error("domain.Name should be example.akadns.net")
	}
	if domain.Status == nil {
		t.Error("domain.Status should not be nil")
	}
}

func TestDomainStatus(t *testing.T) {
	server, client := gtmTestTools(200, domainStatusJson)
	defer server.Close()

	status, _ := client.DomainStatus("example.akadns.net")
	t.Log(pretty.PrettyFormat(status))

	if "DENIED" != status.PropagationStatus {
		t.Error("status.PropagationStatus should be DENIED")
	}
	if "ERROR: Some error" != status.Message {
		t.Error("status.Message should be 'ERROR: Some error'")
	}
}

func TestDomainCreate(t *testing.T) {
	server, client := gtmTestTools(201, createdDomainJson)
	defer server.Close()

	domain, _ := client.DomainCreate("example.akadns.net", "weighted")
	t.Log(pretty.PrettyFormat(domain))

	if domain.Domain == nil {
		t.Error("domain creation response should contain the created Domain")
	}
	if domain.Status == nil {
		t.Error("domain creation response should contain the creation status")
	}
}

func TestDomainCreateFailure(t *testing.T) {
	server, client := gtmTestTools(500, errorResponseJson)
	defer server.Close()

	_, err := client.DomainCreate("example.akadns.net", "weighted")
	t.Log(pretty.PrettyFormat(err))
	if err == nil {
		t.Error("Expected an error back")
	}
	var theError = err.(*AkamaiError)
	if theError.RequestBody == "" || theError.ResponseBody == "" {
		t.Error("Expected AkamaiError to contain the HTTP request and response bodies")
	}
}

func TestDomainUpdate(t *testing.T) {
	server, client := gtmTestTools(200, updatedDomainJson)
	defer server.Close()

	domain := &Domain{
		Name: "example.akadns.net",
		Type: "full",
	}

	updatedDomain, _ := client.DomainUpdate(domain)
	t.Log(pretty.PrettyFormat(updatedDomain))

	if updatedDomain.Domain == nil {
		t.Error("domain update response should contain the updated Domain")
	}
	if updatedDomain.Status == nil {
		t.Error("domain update response should contain the update status")
	}
}

func TestDataCenters(t *testing.T) {
	server, client := gtmTestTools(200, dataCentersJson)
	defer server.Close()

	dcs, _ := client.DataCenters("foo.akadns.net")
	t.Log(pretty.PrettyFormat(dcs))

	if len(dcs) != 2 {
		t.Error("Response should contain 2 DataCenters")
	}

	if dcs[0].Nickname != "dcOne" {
		t.Error("Response should return proper DataCenter item with a Nickname")
	}
}

func TestDataCenterCreate(t *testing.T) {
	server, client := gtmTestTools(201, createdDataCenterJson)
	defer server.Close()

	dataCenter := &DataCenter{
		Nickname:  "Winterfell",
		City:      "Doune",
		Country:   "GB",
		Continent: "EU",
		Latitude:  56.185097,
		Longitude: -4.050264,
	}

	createdDataCenter, _ := client.DataCenterCreate("example.akadns.net", dataCenter)
	t.Log(pretty.PrettyFormat(createdDataCenter))

	if createdDataCenter.DataCenter == nil {
		t.Error("Expected response to contain the DataCenter resource")
	}
	if createdDataCenter.DataCenter.DataCenterID != 3133 {
		t.Error("Expected response to contain the assigned DataCenterID")
	}
	if createdDataCenter.Status == nil {
		t.Error("Ecpected response to contain the creation status")
	}
}

func TestDataCenterByNameSuccess(t *testing.T) {
	server, client := gtmTestTools(200, dataCentersJson)
	defer server.Close()

	dc, _ := client.DataCenterByName("foo.akadns.net", "dcOne")
	t.Log(pretty.PrettyFormat(dc))

	if dc.Nickname != "dcOne" {
		t.Error("DataCenterByName should return a DataCenter with a 'Nickname'")
	}
}

func TestDataCenterByNameError(t *testing.T) {
	server, client := gtmTestTools(200, `{ "items": [] }`)
	defer server.Close()

	_, err := client.DataCenterByName("foo.akadns.net", "dcOne")
	if err == nil {
		t.Error("Expected an error since the DataCenter was not found")
	}
}

func TestDataCenter(t *testing.T) {
	server, client := gtmTestTools(200, dataCenterJson)
	defer server.Close()

	dc, _ := client.DataCenter("foo.akadns.net", 1)
	t.Log(pretty.PrettyFormat(dc))

	if dc.Nickname != "dcNickname" {
		t.Error("DataCenter should return the correct DataCenter for id 1")
	}
}

func TestDataCenterUpdate(t *testing.T) {
	server, client := gtmTestTools(200, createdDataCenterJson)
	defer server.Close()

	dataCenter := &DataCenter{
		Nickname:  "Winterfell",
		City:      "Doune",
		Country:   "GB",
		Continent: "EU",
		Latitude:  56.185097,
		Longitude: -4.050264,
	}

	createdDataCenter, _ := client.DataCenterUpdate("example.akadns.net", dataCenter)
	t.Log(pretty.PrettyFormat(createdDataCenter))

	if createdDataCenter.DataCenter == nil {
		t.Error("Expected response to contain the DataCenter resource")
	}
	if createdDataCenter.DataCenter.DataCenterID != 3133 {
		t.Error("Expected response to contain the assigned DataCenterID")
	}
	if createdDataCenter.Status == nil {
		t.Error("Ecpected response to contain the creation status")
	}
}

func TestDataCenterDeleteSuccess(t *testing.T) {
	server, client := gtmTestTools(200, "")
	defer server.Close()

	err := client.DataCenterDelete("domain", 123)
	if err != nil {
		t.Error("DataCenterDelete should return true if the HTTP request succeeds")
	}
}

func TestDataCenterDeleteFail(t *testing.T) {
	server, client := gtmTestTools(500, "")
	defer server.Close()

	err := client.DataCenterDelete("domain", 123)
	if err.Error() != "HTTP status not OK" {
		t.Error("DataCenterDelete should return an error if the HTTP request returns non-successfully")
	}
}

func TestProperties(t *testing.T) {
	server, client := gtmTestTools(200, propertiesResponseJson)
	defer server.Close()

	props, _ := client.Properties("foo.akadns.net")
	t.Log(pretty.PrettyFormat(props))

	if len(props.Properties) != 5 {
		t.Error("Response should contain 5 Properties")
	}

	if props.Properties[0].Name != "newprop" {
		t.Error("Response should return proper Property item with a Name")
	}
}

func TestPropertiesSorted(t *testing.T) {
	server, client := gtmTestTools(200, propertiesResponseJson)
	defer server.Close()

	props, _ := client.PropertiesSorted("foo.akadns.net")
	t.Log(pretty.PrettyFormat(props))

	if len(props.Properties) != 5 {
		t.Error("Response should contain 5 Properties")
	}

	if props.Properties[0].Name != "anotherprop" {
		t.Error("PropertiesSorted should return the correct Property in the first position")
	}

	if props.Properties[1].Name != "aprop" {
		t.Error("PropertiesSorted should return the correct Property in the second position")
	}

	if props.Properties[2].Name != "newprop" {
		t.Error("PropertiesSorted should return the correct Property in the third position")
	}

	if props.Properties[3].Name != "prop" {
		t.Error("PropertiesSorted should return the correct Property in the fourth position")
	}

	if props.Properties[4].Name != "someprop" {
		t.Error("PropertiesSorted should return the correct Property in the fifth position")
	}
}

func TestProperty(t *testing.T) {
	server, client := gtmTestTools(200, propJson)
	defer server.Close()

	prop, _ := client.Property("foo.akadns.net", "someName")
	t.Log(pretty.PrettyFormat(prop))

	if prop.Name != "someName" {
		t.Error("Property should return the correct Property for the name it's passed")
	}

	if prop.TrafficTargets[0].DataCenterID != 3131 {
		t.Error("Property should return the correct traffic targets for the property associated with the name it's passed")
	}

	if prop.LivenessTests[0].Name != "livenessTestOne" {
		t.Error("Property should return the correct liveness tests for the property associated with the name it's passed")
	}
}

func TestPropertyCreate(t *testing.T) {
	newProperty := &Property{
		Name:                 "origin",
		HandoutMode:          "normal",
		FailoverDelay:        0,
		FailbackDelay:        0,
		Type:                 "weighted-round-robin",
		ScoreAggregationType: "mean",
		TrafficTargets: []TrafficTarget{
			TrafficTarget{
				DataCenterID: 3131,
				Enabled:      true,
				Servers:      []string{"1.2.3.5"},
				Weight:       50,
			},
			TrafficTarget{
				DataCenterID: 3132,
				Enabled:      true,
				Servers:      []string{"1.2.3.4"},
				Weight:       50,
			},
		},
	}

	server, client := gtmTestTools(201, createdPropertyJson)
	defer server.Close()

	createdProperty, err := client.PropertyCreate("foo.akadns.net", newProperty)
	if err != nil {
		t.Fail()
	}
	t.Log(pretty.PrettyFormat(createdProperty))

	if createdProperty.Property == nil {
		t.Error("Property response should contain a Property resource")
	}
	if createdProperty.Status == nil {
		t.Error("Property response should contain the creation status")
	}
}

func TestPropertyCreateFailure(t *testing.T) {
	newProperty := &Property{}

	server, client := gtmTestTools(500, errorResponseJson)
	defer server.Close()

	_, err := client.PropertyCreate("foo.akadns.net", newProperty)
	if err == nil {
		t.Error("Property creation failure should result in an error")
	}
}

func TestPropertyDeleteSuccess(t *testing.T) {
	server, client := gtmTestTools(200, "prop")
	defer server.Close()

	deleted, _ := client.PropertyDelete("domain", "prop")
	if deleted != true {
		t.Error("PropertyDelete should return true if the HTTP request succeeds")
	}
}

func TestPropertyDeleteFail(t *testing.T) {
	server, client := gtmTestTools(500, "prop")
	defer server.Close()

	deleted, err := client.PropertyDelete("domain", "prop")
	if deleted != false {
		t.Error("PropertyDelete should return false if the HTTP request returns non-successfully")
	}
	if err.Error() != "HTTP status not OK" {
		t.Error("PropertyDelete should return an error if the HTTP request returns non-successfully")
	}
}
