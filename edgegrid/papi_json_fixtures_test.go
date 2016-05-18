package edgegrid

const (
	groupsJson = `{
		"groups": {
			"items": [{
				"groupName": "group one",
				"groupId": "1",
				"parentGroupId": "parent1",
				"contractIds": ["contract1", "contract1-2"]
			}, {
				"groupName": "group two",
				"groupId": "2",
				"parentGroupId": "parent2",
				"contractIds": ["contract2", "contract2-2"]
			}]
		}
	}`

	productsJson = `{
		"products": {
			"items": [{
				"productName": "prod one",
				"productId": "1"
			}, {
				"productName": "prod two",
				"productId": "2"
			}]
		}
	}`

	cpCodesJson = `{
		"accountId": "act_1-1TJZFB",
		"contractId": "ctr_1-1TJZFW",
		"groupId": "grp_15225",
		"cpcodes": {
			"items": [{
				"cpcodeId": "someCodeId",
				"cpcodeName": "someName",
				"productIds": ["someId"],
				"createdDate":"someDate"
			}]
		}
	}`

	hostnamesJson = `{
		"edgeHostnames": {
			"items": [{
				"edgeHostnameId": "hostId1",
				"domainPrefix": "foo.com",
				"domainSuffix": "edge.net",
				"ipVersionBehavior": "IPV4",
				"secure" : false,
				"edgeHostnameDomain": "foo.com.edge.net"
			}]
		}
	}`

	propertiesJson = `{
		"properties": {
			"items": [{
				"accountId": "accountId",
				"contractId": "contractId",
				"groupId": "groupid",
				"propertyId": "propertyId",
				"propertyName": "m.example.com",
				"latestVersion": 1,
				"stagingVersion": 2,
				"productionVersion": null,
				"note": "a note"
			}]
		}
	}`

	propertyVersionsJson = `{
		"versions": {
			"items": [{
				"propertyVersion": 2,
				"updatedByUser": "some user",
				"updatedDate": "updatedDate",
				"productionStatus": "INACTIVE",
				"stagingStatus": "ACTIVE",
				"etag": "123",
				"productId": "productId",
				"note": "some note"
			}]
		}
	}`

	propertyRulesJson = `{
		"rules": {
			"name": "default",
			"options": {
				"is_secure": false
			},
			"uuid": "some-uuid",
			"behaviors": [
			{
				"name": "origin",
				"options": {
					"type"              : "customer",
					"hostname"          : "example.com",
					"forwardhostheader" : "requesthostheader",
					"cachekeyhostname"  : "originhostname",
					"compression"       : "on",
					"tcip_enabled"      : "off",
					"http_port"         : "80"
				}
			},
			{
				"name": "cpCode",
				"options": {
					"cpcode": {
						"id"   : 12345,
						"name" : "main site"
					}
				}
			}
			]
		}
	}`

	activationsJson = `{
		"accountId": "act_1-1TJZFB",
		"contractId": "ctr_1-1TJZH5",
		"groupId": "grp_15225",
		"activations": {
			"items": [{
				"activationId": "123",
				"propertyName": "example.com",
				"propertyId": "propId",
				"propertyVersion": 1,
				"network": "STAGING",
				"status": "ACTIVE",
				"activationType": "ACTIVATE",
				"submitDate": "2014-03-05T02:22:12Z",
				"updateDate": "2014-03-04T21:12:57Z",
				"note": "some note",
				"notifyEmails": [
				"you@example.com",
				"them@example.com"
				]
			}]
		}
	}`
)
