package edgegrid

import "io/ioutil"

// Groups returns the PAPI Groups for the PAPIClient
func (c *PAPIClient) Groups() ([]GroupSummary, error) {
	groups := &Groups{}
	err := resourceRequest(c, "GET", papiGroupsEndpoint(c.GetCredentials()), nil, groups)
	if err != nil {
		return []GroupSummary{}, err
	}
	return groups.Groups.Items, err
}

// Products takes a contract ID strings returns the associated products
func (c *PAPIClient) Products(contractId string) ([]ProductSummary, error) {
	prods := &Products{}
	err := resourceRequest(c, "GET", papiProductsEndpoint(c.GetCredentials(), contractId), nil, prods)
	if err != nil {
		return []ProductSummary{}, err
	}
	return prods.Products.Items, err
}

// CpCodes takes a contract ID string and a group ID string and returns the
// associated CP codes
func (c *PAPIClient) CpCodes(contractId, groupId string) ([]CpCodeSummary, error) {
	cps := &CpCodes{}
	err := resourceRequest(c, "GET", papiCpCodesEndpoint(c.GetCredentials(), contractId, groupId), nil, cps)
	if err != nil {
		return []CpCodeSummary{}, err
	}
	return cps.CpCodes.Items, err
}

// CpCode takes a CP code ID string, a contract ID string, and a group ID string
// and returns the associated CP code summary
func (c *PAPIClient) CpCode(cpCodeId, contractId, groupId string) (*CpCodeSummary, error) {
	cps := &CpCodes{}
	err := resourceRequest(c, "GET", papiCpCodeEndpoint(c.GetCredentials(), cpCodeId, contractId, groupId), nil, cps)
	if err != nil {
		return &CpCodeSummary{}, err
	}
	return &cps.CpCodes.Items[0], err
}

// Hostnames takes a contract ID string and a group ID string and returns the
// associated host name summaries
func (c *PAPIClient) Hostnames(contractId, groupId string) ([]HostnameSummary, error) {
	hostnames := &Hostnames{}
	err := resourceRequest(c, "GET", papiHostnamesEndpoint(c.GetCredentials(), contractId, groupId), nil, hostnames)
	if err != nil {
		return []HostnameSummary{}, err
	}
	return hostnames.Hostnames.Items, err
}

// Hostname takes a host ID string, a contract ID string, and a group ID string and returns
// the associated hostname
func (c *PAPIClient) Hostname(hostId, contractId, groupId string) (HostnameSummary, error) {
	hostnames := &Hostnames{}
	err := resourceRequest(c, "GET", papiHostnameEndpoint(c.GetCredentials(), hostId, contractId, groupId), nil, hostnames)
	if err != nil {
		return HostnameSummary{}, err
	}
	return hostnames.Hostnames.Items[0], err
}

// Properties takes a contract ID string and a group ID string and returns the associated
// properties
func (c *PAPIClient) Properties(contractId, groupId string) ([]PapiPropertySummary, error) {
	props := &PapiProperties{}
	err := resourceRequest(c, "GET", papiPropertiesEndpoint(c.GetCredentials(), contractId, groupId), nil, props)
	if err != nil {
		return []PapiPropertySummary{}, err
	}
	return props.Properties.Items, err
}

// Property takes a property ID string, a contract ID string, and a group ID string and returns
// the details for the associated property
func (c *PAPIClient) Property(propId, contractId, groupId string) (PapiPropertySummary, error) {
	props := &PapiProperties{}
	err := resourceRequest(c, "GET", papiPropertyEndpoint(c.GetCredentials(), propId, contractId, groupId), nil, props)
	if err != nil {
		return PapiPropertySummary{}, err
	}
	return props.Properties.Items[0], err
}

// PropertyVersions takes a property ID string, a contract ID string, and a group ID string and
// returns the associated version summaries for the property
func (c *PAPIClient) PropertyVersions(propId, contractId, groupId string) ([]PapiPropertyVersionSummary, error) {
	versions := &PapiPropertyVersions{}
	err := resourceRequest(c, "GET", papiPropertyVersionsEndpoint(c.GetCredentials(), propId, contractId, groupId), nil, versions)
	if err != nil {
		return []PapiPropertyVersionSummary{}, err
	}
	return versions.Versions.Items, err
}

// PropertyVersion takes a version string, a property ID string, a contract ID string, and a group ID string
// and returns the associated property version summary
func (c *PAPIClient) PropertyVersion(version, propId, contractId, groupId string) (PapiPropertyVersionSummary, error) {
	versions := &PapiPropertyVersions{}
	err := resourceRequest(c, "GET", papiPropertyVersionEndpoint(c.GetCredentials(), version, propId, contractId, groupId), nil, versions)
	if err != nil {
		return PapiPropertyVersionSummary{}, err
	}
	return versions.Versions.Items[0], err
}

// PropertyVersionXml takes a version string, a property ID string, a contract ID string, and a group ID string
// and returns the the associated property version XML as a string
func (c *PAPIClient) PropertyVersionXml(version, propId, contractId, groupId string) (string, error) {
	xml, err := getXml(c, papiPropertyVersionEndpoint(c.GetCredentials(), version, propId, contractId, groupId))
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(xml.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// PropertyLatestVersion takes a property ID string, a contract ID string, and a group ID string and returns
// the property version summary for the most recent property version
func (c *PAPIClient) PropertyLatestVersion(propId, contractId, groupId string) (PapiPropertyVersionSummary, error) {
	versions := &PapiPropertyVersions{}
	err := resourceRequest(c, "GET", papiPropertyLatestVersionEndpoint(c.GetCredentials(), propId, contractId, groupId), nil, versions)
	if err != nil {
		return PapiPropertyVersionSummary{}, err
	}
	return versions.Versions.Items[0], err
}

// PropertyRules takes a property ID string, a version string, and a group ID string and returns a
// the rule summary for the associated property
func (c *PAPIClient) PropertyRules(propId, version, contractId, groupId string) (PapiPropertyRuleSummary, error) {
	rules := &PapiPropertyRules{}
	err := resourceRequest(c, "GET", papiPropertyRulesEndpoint(c.GetCredentials(), propId, version, contractId, groupId), nil, rules)
	if err != nil {
		return PapiPropertyRuleSummary{}, err
	}
	return rules.Rules, err
}

// Activations takes a property ID string, a contract ID string, and a group ID string and returns
// the associated property activations
func (c *PAPIClient) Activations(propId, contractId, groupId string) ([]PapiActivation, error) {
	acts := &PapiActivations{}
	err := resourceRequest(c, "GET", papiActivationsEndpoint(c.GetCredentials(), propId, contractId, groupId), nil, acts)
	if err != nil {
		return []PapiActivation{}, err
	}
	return acts.Activations.Items, err
}
