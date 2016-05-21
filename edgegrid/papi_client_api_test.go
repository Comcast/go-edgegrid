package edgegrid

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gobs/pretty"
)

func papiTestTools(code int, body string) (*httptest.Server, *PAPIClient) {
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

	client := &PAPIClient{creds, httpClient}

	return server, client
}

func TestPAPIGroups(t *testing.T) {
	server, client := papiTestTools(200, groupsJson)
	defer server.Close()

	groups, err := client.Groups()
	if err != nil {
		panic(err)
	}
	t.Log(pretty.PrettyFormat(groups))

	if len(groups) != 2 {
		t.Error("Expected 2 Groups")
	}
	for _, g := range groups {
		if g.Name == "" {
			t.Error("Expected Group Name to have a value")
		}
		if g.GroupID == "" {
			t.Error("Expected Group ID to have a value")
		}
	}
}

func TestPAPIProducts(t *testing.T) {
	server, client := papiTestTools(200, productsJson)
	defer server.Close()

	prods, err := client.Products("contractId")
	if err != nil {
		panic(err)
	}
	t.Log(pretty.PrettyFormat(prods))

	if len(prods) != 2 {
		t.Error("Expected 2 Products")
	}
	for _, p := range prods {
		if p.Name == "" {
			t.Error("Expected Product Name to have a value")
		}
		if p.ProductID == "" {
			t.Error("Expected Product ID to have a value")
		}
	}
}

func TestPAPICpCodes(t *testing.T) {
	server, client := papiTestTools(200, cpCodesJson)
	defer server.Close()

	cps, err := client.CpCodes("contractId", "groupId")
	if err != nil {
		panic(err)
	}
	t.Log(pretty.PrettyFormat(cps))

	if len(cps) != 1 {
		t.Error("Expected 1 CPCodes")
	}

	c := cps[0]
	if c.Name != "someName" {
		t.Error("Expected CPCode Name to have a value")
	}
	if c.CPCodeID != "someCodeId" {
		t.Error("Expected CpCode ID to have a value")
	}
	if c.ProductIDs[0] != "someId" {
		t.Error("Expected CPCode ProductIDs to have a value")
	}
	if c.CreatedDate != "someDate" {
		t.Error("Expected CPCode CreatedDate to have a value")
	}
}

func TestPAPICpCode(t *testing.T) {
	server, client := papiTestTools(200, cpCodesJson)
	defer server.Close()

	c, err := client.CpCode("cpCodeId", "contractId", "groupId")
	if err != nil {
		panic(err)
	}
	t.Log(pretty.PrettyFormat(c))

	if c.Name != "someName" {
		t.Error("Expected CpCode Name to have a value")
	}
	if c.CPCodeID != "someCodeId" {
		t.Error("Expected CPCode ID to have a value")
	}
	if c.ProductIDs[0] != "someId" {
		t.Error("Expected CpCode ProductIDs to have a value")
	}
	if c.CreatedDate != "someDate" {
		t.Error("Expected CpCode CreatedDate to have a value")
	}
}

func TestPAPIHostnames(t *testing.T) {
	server, client := papiTestTools(200, hostnamesJson)
	defer server.Close()

	hostnames, err := client.Hostnames("contractId", "groupId")
	if err != nil {
		panic(err)
	}
	t.Log(pretty.PrettyFormat(hostnames))

	if len(hostnames) != 1 {
		t.Error("Expected 1 Hostnames")
	}

	h := hostnames[0]

	if h.EdgeHostnameID != "hostId1" {
		t.Error("Expected Hostname.ID to have correct value")
	}
	if h.DomainPrefix != "foo.com" {
		t.Error("Expected Hostname.DomainPrefix to have correct value")
	}
	if h.DomainSuffix != "edge.net" {
		t.Error("Expected Hostname.DomainSuffix to have correct value")
	}
	if h.IPVersionBehavior != "IPV4" {
		t.Error("Expected Hostname.IPVersionBehavior to have correct value")
	}
	if h.Secure != false {
		t.Error("Expected Hostname.Secure to have correct value")
	}
	if h.EdgeHostnameDomain != "foo.com.edge.net" {
		t.Error("Expected Hostname.EdgeHostnameDomain to have correct value")
	}
}

func TestPAPIHostname(t *testing.T) {
	server, client := papiTestTools(200, hostnamesJson)
	defer server.Close()

	h, err := client.Hostname("hostId", "contractId", "groupId")
	if err != nil {
		panic(err)
	}
	t.Log(pretty.PrettyFormat(h))

	if h.EdgeHostnameID != "hostId1" {
		t.Error("Expected Hostname.ID to have correct value")
	}
	if h.DomainPrefix != "foo.com" {
		t.Error("Expected Hostname.DomainPrefix to have correct value")
	}
	if h.DomainSuffix != "edge.net" {
		t.Error("Expected Hostname.DomainSuffix to have correct value")
	}
	if h.IPVersionBehavior != "IPV4" {
		t.Error("Expected Hostname.IPVersionBehavior to have correct value")
	}
	if h.Secure != false {
		t.Error("Expected Hostname.Secure to have correct value")
	}
	if h.EdgeHostnameDomain != "foo.com.edge.net" {
		t.Error("Expected Hostname.EdgeHostnameDomain to have correct value")
	}
}

func TestPAPIProperties(t *testing.T) {
	server, client := papiTestTools(200, propertiesJson)
	defer server.Close()

	ps, err := client.Properties("contractId", "groupId")
	if err != nil {
		panic(err)
	}
	t.Log(pretty.PrettyFormat(ps))

	p := ps[0]

	if p.PropertyID != "propertyId" {
		t.Error("Expected PropertySummary.PropertyID to have correct value")
	}
	if p.AccountID != "accountId" {
		t.Error("Expected PropertySummary.AccountId to have correct value")
	}
	if p.GroupID != "groupid" {
		t.Error("Expected PropertySummary.GroupID to have correct value")
	}
	if p.ContractID != "contractId" {
		t.Error("Expected PropertySummary.ContractId to have correct value")
	}
	if p.Name != "m.example.com" {
		t.Error("Expected PropertySummary.Name to have correct value")
	}
	if p.LatestVersion != 1 {
		t.Error("Expected PropertySummary.LatestVersion to have correct value")
	}
	if p.StagingVersion != 2 {
		t.Error("Expected PropertySummary.StagingVersion to have correct value")
	}
	if p.ProductionVersion != 0 {
		t.Error("Expected PropertySummary.ProductionVersion to have correct value")
	}
	if p.Note != "a note" {
		t.Error("Expected PropertySummary.Note to have correct value")
	}
}

func TestPAPIProperty(t *testing.T) {
	server, client := papiTestTools(200, propertiesJson)
	defer server.Close()

	p, err := client.Property("propId", "contractId", "groupId")
	if err != nil {
		panic(err)
	}
	t.Log(pretty.PrettyFormat(p))

	if p.PropertyID != "propertyId" {
		t.Error("Expected PropertySummary.PropertyID to have correct value")
	}
	if p.AccountID != "accountId" {
		t.Error("Expected PropertySummary.AccountID to have correct value")
	}
	if p.GroupID != "groupid" {
		t.Error("Expected PropertySummary.GroupID to have correct value")
	}
	if p.ContractID != "contractId" {
		t.Error("Expected PropertySummary.ContractID to have correct value")
	}
	if p.Name != "m.example.com" {
		t.Error("Expected PropertySummary.Name to have correct value")
	}
	if p.LatestVersion != 1 {
		t.Error("Expected PropertySummary.LatestVersion to have correct value")
	}
	if p.StagingVersion != 2 {
		t.Error("Expected PropertySummary.StagingVersion to have correct value")
	}
	if p.ProductionVersion != 0 {
		t.Error("Expected PropertySummary.ProductionVersion to have correct value")
	}
	if p.Note != "a note" {
		t.Error("Expected PropertySummary.Note to have correct value")
	}
}

func TestPAPIPropertyVersions(t *testing.T) {
	server, client := papiTestTools(200, propertyVersionsJson)
	defer server.Close()

	vs, err := client.PropertyVersions("propId", "contractId", "groupId")
	if err != nil {
		panic(err)
	}
	t.Log(pretty.PrettyFormat(vs))

	v := vs[0]

	if v.ProductID != "productId" {
		t.Error("Expected PropertyVersionSummary.ProductID to have correct value")
	}
	if v.PropertyVersion != 2 {
		t.Error("Expected PropertyVersionSummary.PropertyVersion to have correct value")
	}
	if v.UpdatedByUser != "some user" {
		t.Error("Expected PropertyVersionSummary.UpdatedByUser to have correct value")
	}
	if v.UpdatedDate != "updatedDate" {
		t.Error("Expected PropertyVersionSummary.UpdatedDate to have correct value")
	}
	if v.ProductionStatus != "INACTIVE" {
		t.Error("Expected PropertyVersionSummary.ProductionStatus to have correct value")
	}
	if v.StagingStatus != "ACTIVE" {
		t.Error("Expected PropertyVersionSummary.StagingStatus to have correct value")
	}
	if v.Etag != "123" {
		t.Error("Expected PropertyVersionSummary.Etag to have correct value")
	}
	if v.Note != "some note" {
		t.Error("Expected PropertyVersionSummary.Note to have correct value")
	}
}

func TestPAPIPropertyVersion(t *testing.T) {
	server, client := papiTestTools(200, propertyVersionsJson)
	defer server.Close()

	v, err := client.PropertyVersion("2", "propId", "contractId", "groupId")
	if err != nil {
		panic(err)
	}
	t.Log(pretty.PrettyFormat(v))

	if v.ProductID != "productId" {
		t.Error("Expected PropertyVersionSummary.ProductID to have correct value")
	}
	if v.PropertyVersion != 2 {
		t.Error("Expected PropertyVersionSummary.PropertyVersion to have correct value")
	}
	if v.UpdatedByUser != "some user" {
		t.Error("Expected PropertyVersionSummary.UpdatedByUser to have correct value")
	}
	if v.UpdatedDate != "updatedDate" {
		t.Error("Expected PropertyVersionSummary.UpdatedDate to have correct value")
	}
	if v.ProductionStatus != "INACTIVE" {
		t.Error("Expected PropertyVersionSummary.ProductionStatus to have correct value")
	}
	if v.StagingStatus != "ACTIVE" {
		t.Error("Expected PropertyVersionSummary.StagingStatus to have correct value")
	}
	if v.Etag != "123" {
		t.Error("Expected PropertyVersionSummary.Etag to have correct value")
	}
	if v.Note != "some note" {
		t.Error("Expected PropertyVersionSummary.Note to have correct value")
	}
}

func TestPAPIPropertyVersionXml(t *testing.T) {
	server, client := papiTestTools(200, "<fakeXml/>")
	defer server.Close()

	xml, err := client.PropertyVersionXml("2", "propId", "contractId", "groupId")
	if err != nil {
		panic(err)
	}
	t.Log(pretty.PrettyFormat(xml))

	if xml != "<fakeXml/>" {
		t.Error("Expected PropertyVersionXml to return the correct value")
	}
}

func TestPAPIPropertyLatestVersion(t *testing.T) {
	server, client := papiTestTools(200, propertyVersionsJson)
	defer server.Close()

	v, err := client.PropertyLatestVersion("propId", "contractId", "groupId")
	if err != nil {
		panic(err)
	}
	t.Log(pretty.PrettyFormat(v))

	if v.ProductID != "productId" {
		t.Error("Expected PropertyVersionSummary.ProductID to have correct value")
	}
	if v.PropertyVersion != 2 {
		t.Error("Expected PropertyVersionSummary.PropertyVersion to have correct value")
	}
	if v.UpdatedByUser != "some user" {
		t.Error("Expected PropertyVersionSummary.UpdatedByUser to have correct value")
	}
	if v.UpdatedDate != "updatedDate" {
		t.Error("Expected PropertyVersionSummary.UpdatedDate to have correct value")
	}
	if v.ProductionStatus != "INACTIVE" {
		t.Error("Expected PropertyVersionSummary.ProductionStatus to have correct value")
	}
	if v.StagingStatus != "ACTIVE" {
		t.Error("Expected PropertyVersionSummary.StagingStatus to have correct value")
	}
	if v.Etag != "123" {
		t.Error("Expected PropertyVersionSummary.Etag to have correct value")
	}
	if v.Note != "some note" {
		t.Error("Expected PropertyVersionSummary.Note to have correct value")
	}
}

func TestPAPIPropertyRules(t *testing.T) {
	server, client := papiTestTools(200, propertyRulesJson)
	defer server.Close()

	r, err := client.PropertyRules("propId", "version", "contractId", "groupId")
	if err != nil {
		panic(err)
	}
	t.Log(pretty.PrettyFormat(r))

	if r.Name != "default" {
		t.Error("Expected PropertyRuleSummary.Name to have correct value")
	}
	if r.UUID != "some-uuid" {
		t.Error("Expected PropertyRuleSummary.UUID to have correct value")
	}
}

func TestPAPIActivations(t *testing.T) {
	server, client := papiTestTools(200, activationsJson)
	defer server.Close()

	r, err := client.Activations("propId", "contractId", "groupId")
	if err != nil {
		panic(err)
	}
	t.Log(pretty.PrettyFormat(r))

	act := r[0]

	if act.PropertyName != "example.com" {
		t.Error("Expected PapiActivation.PropertyName to have correct value")
	}
	if act.ActivationID != "123" {
		t.Error("Expected PapiActivation.ActivationID to have correct value")
	}
	if act.PropertyID != "propId" {
		t.Error("Expected PapiActivation.PropertyID to have correct value")
	}
	if act.PropertyVersion != 1 {
		t.Error("Expected PapiActivation.PropertyVersion to have correct value")
	}
	if act.PropertyVersion != 1 {
		t.Error("Expected PapiActivation.PropertyVersion to have correct value")
	}
	if act.Network != "STAGING" {
		t.Error("Expected PapiActivation.Network to have correct value")
	}
	if act.ActivationType != "ACTIVATE" {
		t.Error("Expected PapiActivation.ActivationType to have correct value")
	}
	if act.Status != "ACTIVE" {
		t.Error("Expected PapiActivation.Status to have correct value")
	}
	if act.SubmitDate != "2014-03-05T02:22:12Z" {
		t.Error("Expected PapiActivation.SubmitDate to have correct value")
	}
	if act.UpdateDate != "2014-03-04T21:12:57Z" {
		t.Error("Expected PapiActivation.UpdateDate to have correct value")
	}
	if act.Note != "some note" {
		t.Error("Expected PapiActivation.Note to have correct value")
	}
}
