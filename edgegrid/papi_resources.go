package edgegrid

type Groups struct {
	Groups struct {
		Items []GroupSummary `json:"items"`
	} `json:"groups"`
}

type GroupSummary struct {
	GroupId       string   `json:"groupId"`
	Name          string   `json:"groupName"`
	ContractIds   []string `json:"contractIds"`
	ParentGroupId string   `json:"parentGroupId"`
}

type Products struct {
	Products struct {
		Items []ProductSummary `json:"items"`
	} `json:"products"`
}

type ProductSummary struct {
	ProductId string `json:"productId"`
	Name      string `json:"productName"`
}

type CpCodes struct {
	CpCodes struct {
		Items []CpCodeSummary `json:"items"`
	} `json:"cpcodes"`
}

type CpCodeSummary struct {
	CpCodeId    string   `json:"cpcodeId"`
	Name        string   `json:"cpcodeName"`
	ProductIds  []string `json:"productIds"`
	CreatedDate string   `json:"createdDate"`
}

type Hostnames struct {
	Hostnames struct {
		Items []HostnameSummary `json:"items"`
	} `json:"edgehostnames"`
}

type HostnameSummary struct {
	EdgeHostnameId     string `json:"edgeHostnameId"`
	DomainPrefix       string `json:"domainPrefix"`
	DomainSuffix       string `json:"domainSuffix"`
	IpVersionBehavior  string `json:"ipVersionBehavior"`
	Secure             bool   `json:"secure"`
	EdgeHostnameDomain string `json:"edgehostnameDomain"`
}

type PapiProperties struct {
	Properties struct {
		Items []PapiPropertySummary `json:"items"`
	} `json:"properties"`
}

type PapiPropertySummary struct {
	AccountId         string `json:"accountId"`
	ContractId        string `json:"contractId"`
	GroupId           string `json:"groupId"`
	PropertyId        string `json:"propertyId"`
	Name              string `json:"propertyName"`
	LatestVersion     int    `json:"latestVersion"`
	StagingVersion    int    `json:"stagingVersion"`
	ProductionVersion int    `json:"productionVersion"`
	Note              string `json:"note"`
}

type PapiPropertyVersions struct {
	Versions struct {
		Items []PapiPropertyVersionSummary `json:"items"`
	} `json:"versions"`
}

type PapiPropertyVersionSummary struct {
	PropertyVersion  int    `json:"propertyVersion"`
	UpdatedByUser    string `json:"updatedByUser"`
	UpdatedDate      string `json:"updatedDate"`
	ProductionStatus string `json:"productionStatus"`
	StagingStatus    string `json:"stagingStatus"`
	Etag             string `json:"etag"`
	ProductId        string `json:"productId"`
	Note             string `json:"note"`
}

type PapiPropertyRules struct {
	Rules PapiPropertyRuleSummary `json:"rules"`
}

type PapiPropertyRuleSummary struct {
	Name      string                     `json:"name"`
	Uuid      string                     `json:"uuid"`
	Behaviors []PapiPropertyRuleBehavior `json:"behaviors"`
}

type PapiPropertyRuleBehavior struct {
	Name string `json:"name"`
}

type PapiActivations struct {
	Activations struct {
		Items []PapiActivation `json:"items"`
	} `json:"activations"`
}

type PapiActivation struct {
	ActivationId    string `json:"activationId"`
	PropertyName    string `json:"propertyName"`
	PropertyId      string `json:"propertyId"`
	PropertyVersion int    `json:"propertyVersion"`
	Network         string `json:"network"`
	ActivationType  string `json:"activationType"`
	Status          string `json:"status"`
	SubmitDate      string `json:"submitDate"`
	UpdateDate      string `json:"updateDate"`
	Note            string `json:"note"`
}
