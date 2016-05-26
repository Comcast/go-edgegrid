package edgegrid

import "io/ioutil"

// Groups returns the PAPI Groups for the PAPIClient.
func (c *PAPIClient) Groups() ([]GroupSummary, error) {
	groups := &Groups{}
	err := resourceRequest(c, "GET", papiGroupsEndpoint(c.GetCredentials()), nil, groups)
	if err != nil {
		return []GroupSummary{}, err
	}
	return groups.Groups.Items, err
}

// Products takes a contractID strings returns the associated products.
func (c *PAPIClient) Products(contractID string) ([]ProductSummary, error) {
	prods := &Products{}
	err := resourceRequest(c, "GET", papiProductsEndpoint(c.GetCredentials(), contractID), nil, prods)
	if err != nil {
		return []ProductSummary{}, err
	}
	return prods.Products.Items, err
}

// CpCodes takes a contractID string and a group ID string and returns the
// associated CP codes.
func (c *PAPIClient) CpCodes(contractID, groupID string) ([]CpCodeSummary, error) {
	cps := &CpCodes{}
	err := resourceRequest(c, "GET", papiCpCodesEndpoint(c.GetCredentials(), contractID, groupID), nil, cps)
	if err != nil {
		return []CpCodeSummary{}, err
	}
	return cps.CpCodes.Items, err
}

// CpCode takes a cpCodeID, a contractID, and a groupID
// and returns the associated CP code summary.
func (c *PAPIClient) CpCode(cpCodeID, contractID, groupID string) (*CpCodeSummary, error) {
	cps := &CpCodes{}
	err := resourceRequest(c, "GET", papiCpCodeEndpoint(c.GetCredentials(), cpCodeID, contractID, groupID), nil, cps)
	if err != nil {
		return &CpCodeSummary{}, err
	}
	return &cps.CpCodes.Items[0], err
}

// Hostnames takes a contractID and a groupID string and returns the
// associated host name summaries.
func (c *PAPIClient) Hostnames(contractID, groupID string) ([]HostnameSummary, error) {
	hostnames := &Hostnames{}
	err := resourceRequest(c, "GET", papiHostnamesEndpoint(c.GetCredentials(), contractID, groupID), nil, hostnames)
	if err != nil {
		return []HostnameSummary{}, err
	}
	return hostnames.Hostnames.Items, err
}

// Hostname takes a hostID, a contractID, and a groupID and returns
// the associated hostname.
func (c *PAPIClient) Hostname(hostID, contractID, groupID string) (HostnameSummary, error) {
	hostnames := &Hostnames{}
	err := resourceRequest(c, "GET", papiHostnameEndpoint(c.GetCredentials(), hostID, contractID, groupID), nil, hostnames)
	if err != nil {
		return HostnameSummary{}, err
	}
	return hostnames.Hostnames.Items[0], err
}

// Properties takes a contractID and a groupID and returns the associated
// properties.
func (c *PAPIClient) Properties(contractID, groupID string) ([]PapiPropertySummary, error) {
	props := &PapiProperties{}
	err := resourceRequest(c, "GET", papiPropertiesEndpoint(c.GetCredentials(), contractID, groupID), nil, props)
	if err != nil {
		return []PapiPropertySummary{}, err
	}
	return props.Properties.Items, err
}

// Property takes a propertyID, a contractID, and a groupID and returns
// the details for the associated property.
func (c *PAPIClient) Property(propID, contractID, groupID string) (PapiPropertySummary, error) {
	props := &PapiProperties{}
	err := resourceRequest(c, "GET", papiPropertyEndpoint(c.GetCredentials(), propID, contractID, groupID), nil, props)
	if err != nil {
		return PapiPropertySummary{}, err
	}
	return props.Properties.Items[0], err
}

// PropertyVersions takes a propertyID, a contractID, and a groupID and
// returns the associated version summaries for the property.
func (c *PAPIClient) PropertyVersions(propID, contractID, groupID string) ([]PapiPropertyVersionSummary, error) {
	versions := &PapiPropertyVersions{}
	err := resourceRequest(c, "GET", papiPropertyVersionsEndpoint(c.GetCredentials(), propID, contractID, groupID), nil, versions)
	if err != nil {
		return []PapiPropertyVersionSummary{}, err
	}
	return versions.Versions.Items, err
}

// PropertyVersion takes a version, a propertyID, a contractID, and a groupID
// and returns the associated property version summary.
func (c *PAPIClient) PropertyVersion(version, propID, contractID, groupID string) (PapiPropertyVersionSummary, error) {
	versions := &PapiPropertyVersions{}
	err := resourceRequest(c, "GET", papiPropertyVersionEndpoint(c.GetCredentials(), version, propID, contractID, groupID), nil, versions)
	if err != nil {
		return PapiPropertyVersionSummary{}, err
	}
	return versions.Versions.Items[0], err
}

// PropertyVersionXML takes a version, a propertyID, a contractID, and a groupID
// and returns the the associated property version XML as a string.
func (c *PAPIClient) PropertyVersionXML(version, propID, contractID, groupID string) (string, error) {
	xml, err := getXML(c, papiPropertyVersionEndpoint(c.GetCredentials(), version, propID, contractID, groupID))
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(xml.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// PropertyLatestVersion takes a propertyID, a contractID, and a groupID and returns
// the property version summary for the most recent property version.
func (c *PAPIClient) PropertyLatestVersion(propID, contractID, groupID string) (PapiPropertyVersionSummary, error) {
	versions := &PapiPropertyVersions{}
	err := resourceRequest(c, "GET", papiPropertyLatestVersionEndpoint(c.GetCredentials(), propID, contractID, groupID), nil, versions)
	if err != nil {
		return PapiPropertyVersionSummary{}, err
	}
	return versions.Versions.Items[0], err
}

// PropertyRules takes a propertyID string, a version, and a groupID and returns a
// the rule summary for the associated property.
func (c *PAPIClient) PropertyRules(propID, version, contractID, groupID string) (PapiPropertyRuleSummary, error) {
	rules := &PapiPropertyRules{}
	err := resourceRequest(c, "GET", papiPropertyRulesEndpoint(c.GetCredentials(), propID, version, contractID, groupID), nil, rules)
	if err != nil {
		return PapiPropertyRuleSummary{}, err
	}
	return rules.Rules, err
}

// Activations takes a propertyID, a contractID, and a groupID and returns
// the associated property activations.
func (c *PAPIClient) Activations(propID, contractID, groupID string) ([]PapiActivation, error) {
	acts := &PapiActivations{}
	err := resourceRequest(c, "GET", papiActivationsEndpoint(c.GetCredentials(), propID, contractID, groupID), nil, acts)
	if err != nil {
		return []PapiActivation{}, err
	}
	return acts.Activations.Items, err
}
