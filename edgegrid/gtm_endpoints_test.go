package edgegrid

import "testing"

var c = &AuthCredentials{"accessToken", "clientToken", "clientSecret", "http://apibase.com"}

func TestGtmBase(t *testing.T) {
	if gtmBase(c) != "http://apibase.com/config-gtm/v1/" {
		t.Error("gtmBase should return the proper endpoint")
	}
}

func TestDomainsEndpoint(t *testing.T) {
	if domainsEndpoint(c) != "http://apibase.com/config-gtm/v1/domains" {
		t.Error("domainsEndpoint should return the proper endpoint")
	}
}

func TestDomainEndpoint(t *testing.T) {
	if domainEndpoint(c, "domain") != "http://apibase.com/config-gtm/v1/domains/domain" {
		t.Error("domainEndpoint should return the proper endpoint")
	}
}

func TestDomainStatusEndpoint(t *testing.T) {
	if domainStatusEndpoint(c, "domain") != "http://apibase.com/config-gtm/v1/domains/domain/status/current" {
		t.Error("domainStatusEndpoint should return the proper endpoint")
	}
}

func TestDcsEndpoint(t *testing.T) {
	if dcsEndpoint(c, "domain") != "http://apibase.com/config-gtm/v1/domains/domain/datacenters" {
		t.Error("dcsEndpoint should return the proper endpoint")
	}
}

func TestDcEndpoint(t *testing.T) {
	if dcEndpoint(c, "domain", 4) != "http://apibase.com/config-gtm/v1/domains/domain/datacenters/4" {
		t.Error("dcEndpoint should return the proper endpoint")
	}
}

func TestPropertiesEndpoint(t *testing.T) {
	if propertiesEndpoint(c, "domain") != "http://apibase.com/config-gtm/v1/domains/domain/properties" {
		t.Error("propertiesEndpoint should return the proper endpoint")
	}
}

func TestPropertyEndpoint(t *testing.T) {
	if propertyEndpoint(c, "domain", "property") != "http://apibase.com/config-gtm/v1/domains/domain/properties/property" {
		t.Error("propertyEndpoint should return the proper endpoint")
	}
}
