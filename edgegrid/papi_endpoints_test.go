package edgegrid

import "testing"

var a = &AuthCredentials{"accessToken", "clientToken", "clientSecret", "http://apibase.com"}

func TestPapiBase(t *testing.T) {
	if papiBase(a) != "http://apibase.com/papi/v0/" {
		t.Error("papiBase should return the proper endpoint")
	}
}

func TestGroups(t *testing.T) {
	if papiGroupsEndpoint(a) != "http://apibase.com/papi/v0/groups/" {
		t.Error("groups should return the proper endpoint")
	}
}

func TestProducts(t *testing.T) {
	if papiProductsEndpoint(a, "someId") != "http://apibase.com/papi/v0/products?contractId=someId" {
		t.Error("products should return the proper endpoint")
	}
}

func TestCpCodes(t *testing.T) {
	if papiCpCodesEndpoint(a, "contractId", "groupId") != "http://apibase.com/papi/v0/cpcodes/?contractId=contractId&groupId=groupId" {
		t.Error("cpCodes should return the proper endpoint")
	}
}

func TestCpCode(t *testing.T) {
	if papiCpCodeEndpoint(a, "cpCodeId", "contractId", "groupId") != "http://apibase.com/papi/v0/cpcodes/cpCodeId?contractId=contractId&groupId=groupId" {
		t.Error("cpCode should return the proper endpoint")
	}
}

func TestHostnames(t *testing.T) {
	if papiHostnamesEndpoint(a, "contractId", "groupId") != "http://apibase.com/papi/v0/edgehostnames?contractId=contractId&groupId=groupId" {
		t.Error("hostnames should return the proper endpoint")
	}
}

func TestHostname(t *testing.T) {
	if papiHostnameEndpoint(a, "hostId", "contractId", "groupId") != "http://apibase.com/papi/v0/edgehostnames/hostId?contractId=contractId&groupId=groupId" {
		t.Error("hostname should return the proper endpoint")
	}
}

func TestPropertiesEp(t *testing.T) {
	if papiPropertiesEndpoint(a, "contractId", "groupId") != "http://apibase.com/papi/v0/properties/?contractId=contractId&groupId=groupId" {
		t.Error("papiProperties should return the proper endpoint")
	}
}

func TestPropertyEp(t *testing.T) {
	if papiPropertyEndpoint(a, "propId", "contractId", "groupId") != "http://apibase.com/papi/v0/properties/propId?contractId=contractId&groupId=groupId" {
		t.Error("papiProperty should return the proper endpoint")
	}
}

func TestPropertyVersionsEp(t *testing.T) {
	if papiPropertyVersionsEndpoint(a, "propId", "contractId", "groupId") != "http://apibase.com/papi/v0/properties/propId/versions?contractId=contractId&groupId=groupId" {
		t.Error("papiPropertyVersions should return the proper endpoint")
	}
}

func TestPropertyVersionEp(t *testing.T) {
	if papiPropertyVersionEndpoint(a, "2", "propId", "contractId", "groupId") != "http://apibase.com/papi/v0/properties/propId/versions/2?contractId=contractId&groupId=groupId" {
		t.Error("papiPropertyVersion should return the proper endpoint")
	}
}

func TestPropertyLatestVersionEp(t *testing.T) {
	if papiPropertyLatestVersionEndpoint(a, "propId", "contractId", "groupId") != "http://apibase.com/papi/v0/properties/propId/versions/latest?contractId=contractId&groupId=groupId" {
		t.Error("papiPropertyLatestVersion should return the proper endpoint")
	}
}

func TestPropertyRulesEp(t *testing.T) {
	if papiPropertyRulesEndpoint(a, "propId", "version", "contractId", "groupId") != "http://apibase.com/papi/v0/properties/propId/versions/version/rules/?contractId=contractId&groupId=groupId" {
		t.Error("papiPropertyRules should return the proper endpoint")
	}
}

func TestPropertyHostnamesEp(t *testing.T) {
	if papiPropertyHostnamesEndpoint(a, "propId", "version", "contractId", "groupId") != "http://apibase.com/papi/v0/properties/propId/versions/version/hostnames/?contractId=contractId&groupId=groupId" {
		t.Error(papiPropertyHostnamesEndpoint(a, "propId", "version", "contractId", "groupId"))
	}
}

func TestActivationsEp(t *testing.T) {
	if papiActivationsEndpoint(a, "propId", "contractId", "groupId") != "http://apibase.com/papi/v0/properties/propId/activations?contractId=contractId&groupId=groupId" {
		t.Error("papiActivations should return the proper endpoint")
	}
}
