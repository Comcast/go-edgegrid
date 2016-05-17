package edgegrid

const papiPath = "/papi/v0/"

func papiBase(c *AuthCredentials) string {
	return concat([]string{
		c.ApiHost,
		papiPath,
	})
}

func papiGroupsEndpoint(c *AuthCredentials) string {
	return concat([]string{
		papiBase(c),
		"groups/",
	})
}

func papiProductsEndpoint(c *AuthCredentials, contractId string) string {
	return concat([]string{
		papiBase(c),
		"products?contractId=",
		contractId,
	})
}

func papiCpCodesEndpoint(c *AuthCredentials, contractId, groupId string) string {
	return concat([]string{
		papiBase(c),
		"cpcodes/",
		papiQuery(contractId, groupId),
	})
}

func papiCpCodeEndpoint(c *AuthCredentials, cpCodeId, contractId, groupId string) string {
	return concat([]string{
		papiBase(c),
		"cpcodes/",
		cpCodeId,
		papiQuery(contractId, groupId),
	})
}

func papiQuery(contractId, groupId string) string {
	return concat([]string{
		"?contractId=",
		contractId,
		"&groupId=",
		groupId,
	})
}

func papiHostnamesEndpoint(c *AuthCredentials, contractId, groupId string) string {
	return concat([]string{
		papiBase(c),
		"edgehostnames",
		papiQuery(contractId, groupId),
	})
}

func papiHostnameEndpoint(c *AuthCredentials, hostId, contractId, groupId string) string {
	return concat([]string{
		papiBase(c),
		"edgehostnames/",
		hostId,
		papiQuery(contractId, groupId),
	})
}

func papiPropertiesBase(c *AuthCredentials) string {
	return concat([]string{
		papiBase(c),
		"properties/",
	})
}

func papiPropertiesEndpoint(c *AuthCredentials, contractId, groupId string) string {
	return concat([]string{
		papiPropertiesBase(c),
		papiQuery(contractId, groupId),
	})
}

func papiPropertyBase(c *AuthCredentials, propId string) string {
	return concat([]string{
		papiPropertiesBase(c),
		propId,
	})
}

func papiPropertyEndpoint(c *AuthCredentials, propId, contractId, groupId string) string {
	return concat([]string{
		papiPropertyBase(c, propId),
		papiQuery(contractId, groupId),
	})
}

func papiPropertyVersionsBase(c *AuthCredentials, propId, contractId, groupId string) string {
	return concat([]string{
		papiPropertyBase(c, propId),
		"/versions",
	})
}

func papiPropertyVersionsEndpoint(c *AuthCredentials, propId, contractId, groupId string) string {
	return concat([]string{
		papiPropertyVersionsBase(c, propId, contractId, groupId),
		papiQuery(contractId, groupId),
	})
}

func papiPropertyVersionBase(c *AuthCredentials, version, propId, contractId, groupId string) string {
	return concat([]string{
		papiPropertyVersionsBase(c, propId, contractId, groupId),
		"/",
		version,
	})
}

func papiPropertyVersionEndpoint(c *AuthCredentials, version, propId, contractId, groupId string) string {
	return concat([]string{
		papiPropertyVersionBase(c, version, propId, contractId, groupId),
		papiQuery(contractId, groupId),
	})
}

func papiPropertyLatestVersionEndpoint(c *AuthCredentials, propId, contractId, groupId string) string {
	return concat([]string{
		papiPropertyVersionsBase(c, propId, contractId, groupId),
		"/latest",
		papiQuery(contractId, groupId),
	})
}

func papiPropertyRulesEndpoint(c *AuthCredentials, propId, version, contractId, groupId string) string {
	return concat([]string{
		papiPropertyVersionBase(c, version, propId, contractId, groupId),
		"/rules/",
		papiQuery(contractId, groupId),
	})
}

func papiActivationsEndpoint(c *AuthCredentials, propId, contractId, groupId string) string {
	return concat([]string{
		papiPropertyBase(c, propId),
		"/activations",
		papiQuery(contractId, groupId),
	})
}
