package edgegrid

const (
	domainsJson = `{
	  "items": [
	    {
	      "acgId": "1-2345",
	      "lastModified": "2014-03-03T16:02:45.000+0000",
	      "links": [
	        {
	          "href": "/config-gtm/v1/domains/example.akadns.net",
	          "rel": "self"
	        }
	      ],
	      "name": "example.akadns.net",
	      "status": "2014-02-20 22:56 GMT: Current configuration has been propagated to all toplevels"
	    },
	    {
	      "acgId": "1-2345",
	      "lastModified": "2013-11-09T12:04:45.000+0000",
	      "links": [
	        {
	          "href": "/config-gtm/v1/domains/demo.akadns.net",
	          "rel": "self"
	        }
	      ],
	      "name": "demo.akadns.net",
	      "status": "2014-02-20 22:56 GMT: Current configuration has been propagated to all toplevels"
	    }
	  ]
	}`

	createdDomainJson = `{
    "resource": {
        "cidrMaps": [],
        "datacenters": [],
        "defaultErrorPenalty": 75,
        "defaultSslClientCertificate": null,
        "defaultSslClientPrivateKey": null,
        "defaultTimeoutPenalty": 25,
        "emailNotificationList": [
            "our.admin@example.com"
        ],
        "geographicMaps": [],
        "lastModified": "2014-08-04T15:30:40.057+0000",
        "lastModifiedBy": "admin",
        "links": [
            {
                "href": "/config-gtm/v1/domains/example.akadns.net",
                "rel": "self"
            }
        ],
        "loadFeedback": false,
        "loadImbalancePercentage": null,
        "modificationComments": "New Traffic Management domain.",
        "name": "example.akadns.net",
        "properties": [],
        "resources": [],
        "status": {
            "changeId": "abb94136-dd94-4779-bfb0-fcfac67fb843",
            "links": [
                {
                    "href": "/config-gtm/v1/domains/example.akadns.net/status/current",
                    "rel": "self"
                }
            ],
            "message": "Change Pending",
            "passingValidation": true,
            "propagationStatus": "PENDING",
            "propagationStatusDate": "2014-08-04T15:30:40.057+0000"
        },
        "type": "weighted"
    },
    "status": {
        "changeId": "abb94136-dd94-4779-bfb0-fcfac67fb843",
        "links": [
            {
                "href": "/config-gtm/v1/domains/example.akadns.net/status/current",
                "rel": "self"
            }
        ],
        "message": "Change Pending",
        "passingValidation": true,
        "propagationStatus": "PENDING",
        "propagationStatusDate": "2014-08-04T15:30:40.057+0000"
    }
	}`

	domainJson = `{
      "links": [
          {
              "href": "/config-gtm/v1/domains/example.akadns.net",
              "rel": "self"
          },
          {
              "href": "/config-gtm/v1/domains/example.akadns.net/datacenters",
              "rel": "datacenters"
          },
          {
              "href": "/config-gtm/v1/domains/example.akadns.net/properties",
              "rel": "properties"
          },
          {
              "href": "/config-gtm/v1/domains/example.akadns.net/geographic-maps",
              "rel": "geographic-maps"
          },
          {
              "href": "/config-gtm/v1/domains/example.akadns.net/cidr-maps",
              "rel": "cidr-maps"
          },
          {
              "href": "/config-gtm/v1/domains/example.akadns.net/resources",
              "rel": "resources"
          }
      ],
      "cidrMaps": [
          {
              "assignments": [
                  {
                      "blocks": [
                          "1.3.5.9",
                          "1.2.3.0/24"
                      ],
                      "datacenterId": 3134,
                      "nickname": "Frostfangs and the Fist of First Men"
                  },
                  {
                      "blocks": [
                          "1.2.4.0/24"
                      ],
                      "datacenterId": 3133,
                      "nickname": "Winterfell"
                  }
              ],
              "defaultDatacenter": {
                  "datacenterId": 5400,
                  "nickname": "All Other CIDR Blocks"
              },
              "links": [
                  {
                      "href": "/config-gtm/v1/domains/example.akadns.net/cidr-maps/The%20North",
                      "rel": "self"
                  }
              ],
              "name": "The North"
          }
      ],
      "datacenters": [
          {
              "city": "Downpatrick",
              "cloneOf": 0,
              "continent": "EU",
              "country": "GB",
              "datacenterId": 3133,
              "defaultLoadObject": {
                  "loadObject": null,
                  "loadObjectPort": 0,
                  "loadServers": null
              },
              "latitude": 54.367,
              "links": [
                  {
                      "href": "/config-gtm/v1/domains/example.akadns.net/datacenters/3133",
                      "rel": "self"
                  }
              ],
              "longitude": -5.582,
              "nickname": "Winterfell",
              "stateOrProvince": null,
              "virtual": true
          },
          {
              "city": "Doune",
              "cloneOf": 0,
              "continent": "EU",
              "country": "GB",
              "datacenterId": 3134,
              "defaultLoadObject": {
                  "loadObject": null,
                  "loadObjectPort": 0,
                  "loadServers": null
              },
              "latitude": 56.185097,
              "links": [
                  {
                      "href": "/config-gtm/v1/domains/example.akadns.net/datacenters/3134",
                      "rel": "self"
                  }
              ],
              "longitude": -4.050264,
              "nickname": "Winterfell",
              "stateOrProvince": "Perthshire",
              "virtual": true
          },
          {
              "city": null,
              "cloneOf": 0,
              "continent": null,
              "country": null,
              "datacenterId": 5400,
              "defaultLoadObject": {
                  "loadObject": null,
                  "loadObjectPort": 0,
                  "loadServers": null
              },
              "latitude": 0.0,
              "links": [
                  {
                      "href": "/config-gtm/v1/domains/example.akadns.net/datacenters/5400",
                      "rel": "self"
                  }
              ],
              "longitude": 0.0,
              "nickname": "Default Datacenter",
              "stateOrProvince": null,
              "virtual": true
          }
      ],
      "defaultSslClientCertificate": null,
      "defaultSslClientPrivateKey": null,
      "defaultUnreachableThreshold": null,
      "emailNotificationList": [],
      "geographicMaps": [
          {
              "assignments": [
                  {
                      "countries": [
                          "GB"
                      ],
                      "datacenterId": 3133,
                      "nickname": "UK users"
                  }
              ],
              "defaultDatacenter": {
                  "datacenterId": 5400,
                  "nickname": "Default Mapping"
              },
              "links": [
                  {
                      "href": "/config-gtm/v1/domains/example.akadns.net/geographic-maps/UK%20Delivery",
                      "rel": "self"
                  }
              ],
              "name": "UK Delivery"
          }
      ],
      "lastModified": "2014-04-08T18:25:51.000+0000",
      "lastModifiedBy": "admin@example.com",
      "loadFeedback": true,
      "loadImbalancePercentage": 10.0,
      "minPingableRegionFraction": null,
      "modificationComments": "CIDRMap example",
      "name": "example.akadns.net",
      "pingInterval": null,
      "properties": [
          {
              "backupCName": null,
              "backupIp": null,
              "balanceByDownloadScore": false,
              "cname": null,
              "comments": null,
              "dynamicTTL": 300,
              "failbackDelay": 0,
              "failoverDelay": 0,
              "handoutMode": "normal",
              "healthMax": null,
              "healthMultiplier": null,
              "healthThreshold": null,
              "ipv6": false,
              "lastModified": "2014-04-08T18:25:52.000+0000",
              "links": [
                  {
                      "href": "/config-gtm/v1/domains/example.akadns.net/properties/www",
                      "rel": "self"
                  }
              ],
              "livenessTests": [
                  {
                      "disableNonstandardPortWarning": false,
                      "hostHeader": "foo.example.com",
                      "httpError3xx": true,
                      "httpError4xx": true,
                      "httpError5xx": true,
                      "links": [],
                      "name": "health-check",
                      "requestString": null,
                      "responseString": null,
                      "sslClientCertificate": null,
                      "sslClientPrivateKey": null,
                      "testInterval": 60,
                      "testObject": "/status",
                      "testObjectPassword": null,
                      "testObjectPort": 80,
                      "testObjectProtocol": "HTTP",
                      "testObjectUsername": null,
                      "testTimeout": 25.0
                  }
              ],
              "mapName": null,
              "maxUnreachablePenalty": null,
              "mxRecords": [],
              "name": "www",
              "scoreAggregationType": "mean",
              "staticTTL": 600,
              "stickinessBonusConstant": 0,
              "stickinessBonusPercentage": 0,
              "trafficTargets": [
                  {
                      "datacenterId": 5400,
                      "enabled": false,
                      "handoutCName": null,
                      "name": null,
                      "servers": [],
                      "weight": 0.0
                  },
                  {
                      "datacenterId": 3134,
                      "enabled": true,
                      "handoutCName": null,
                      "name": null,
                      "servers": [
                          "1.2.3.5"
                      ],
                      "weight": 0.0
                  },
                  {
                      "datacenterId": 3133,
                      "enabled": true,
                      "handoutCName": null,
                      "name": null,
                      "servers": [
                          "1.2.3.4"
                      ],
                      "weight": 1.0
                  }
              ],
              "type": "failover",
              "unreachableThreshold": null,
              "useComputedTargets": false
          },
          {
              "backupCName": null,
              "backupIp": null,
              "balanceByDownloadScore": false,
              "cname": null,
              "comments": null,
              "dynamicTTL": 300,
              "failbackDelay": 0,
              "failoverDelay": 0,
              "handoutMode": "normal",
              "healthMax": null,
              "healthMultiplier": null,
              "healthThreshold": null,
              "ipv6": true,
              "lastModified": "2014-04-08T18:25:52.000+0000",
              "links": [
                  {
                      "href": "/config-gtm/v1/domains/example.akadns.net/properties/mail",
                      "rel": "self"
                  }
              ],
              "livenessTests": [],
              "loadImbalancePercentage": null,
              "mapName": null,
              "maxUnreachablePenalty": null,
              "mxRecords": [],
              "name": "mail",
              "scoreAggregationType": "mean",
              "staticTTL": 600,
              "stickinessBonusConstant": 0,
              "stickinessBonusPercentage": 0,
              "trafficTargets": [
                  {
                      "datacenterId": 5400,
                      "enabled": false,
                      "handoutCName": null,
                      "name": null,
                      "servers": [],
                      "weight": 0.0
                  },
                  {
                      "datacenterId": 3134,
                      "enabled": true,
                      "handoutCName": null,
                      "name": null,
                      "servers": [
                          "2001:4878::5043:4078"
                      ],
                      "weight": 1.0
                  },
                  {
                      "datacenterId": 3133,
                      "enabled": true,
                      "handoutCName": null,
                      "name": null,
                      "servers": [
                          "2001:4878::5043:4072",
                          "2001:4878::5043:4071"
                      ],
                      "weight": 1.0
                  }
              ],
              "type": "weighted-round-robin",
              "unreachableThreshold": null,
              "useComputedTargets": false
          },
          {
              "backupCName": null,
              "backupIp": null,
              "balanceByDownloadScore": false,
              "cname": null,
              "comments": null,
              "dynamicTTL": 300,
              "failbackDelay": 0,
              "failoverDelay": 0,
              "handoutMode": "normal",
              "healthMax": null,
              "healthMultiplier": null,
              "healthThreshold": null,
              "ipv6": false,
              "lastModified": "2014-04-08T18:25:52.000+0000",
              "links": [
                  {
                      "href": "/config-gtm/v1/domains/example.akadns.net/properties/supplies",
                      "rel": "self"
                  }
              ],
              "livenessTests": [],
              "loadImbalancePercentage": null,
              "mapName": null,
              "maxUnreachablePenalty": null,
              "mxRecords": [],
              "name": "supplies",
              "scoreAggregationType": "mean",
              "staticTTL": 600,
              "stickinessBonusConstant": 0,
              "stickinessBonusPercentage": 0,
              "trafficTargets": [
                  {
                      "datacenterId": 5400,
                      "enabled": false,
                      "handoutCName": "supplies.example.com",
                      "name": null,
                      "servers": [],
                      "weight": 0.0
                  },
                  {
                      "datacenterId": 3134,
                      "enabled": true,
                      "handoutCName": "winter.supplies.example.com",
                      "name": null,
                      "servers": [],
                      "weight": 0.0
                  },
                  {
                      "datacenterId": 3133,
                      "enabled": true,
                      "handoutCName": "redcross.org",
                      "name": null,
                      "servers": [],
                      "weight": 0.0
                  }
              ],
              "type": "failover",
              "unreachableThreshold": null,
              "useComputedTargets": false
          },
          {
              "backupCName": null,
              "backupIp": null,
              "balanceByDownloadScore": false,
              "cname": null,
              "comments": null,
              "dynamicTTL": 300,
              "failbackDelay": 0,
              "failoverDelay": 0,
              "handoutMode": "normal",
              "healthMax": null,
              "healthMultiplier": null,
              "healthThreshold": null,
              "ipv6": false,
              "lastModified": "2014-04-08T18:25:52.000+0000",
              "links": [
                  {
                      "href": "/config-gtm/v1/domains/example.akadns.net/properties/shop",
                      "rel": "self"
                  }
              ],
              "livenessTests": [],
              "loadImbalancePercentage": null,
              "mapName": "UK Delivery",
              "maxUnreachablePenalty": null,
              "mxRecords": [],
              "name": "shop",
              "scoreAggregationType": "mean",
              "staticTTL": 600,
              "stickinessBonusConstant": 0,
              "stickinessBonusPercentage": 0,
              "trafficTargets": [
                  {
                      "datacenterId": 5400,
                      "enabled": true,
                      "handoutCName": "shop.example.com",
                      "name": null,
                      "servers": [],
                      "weight": 1.0
                  },
                  {
                      "datacenterId": 3134,
                      "enabled": false,
                      "handoutCName": null,
                      "name": null,
                      "servers": [],
                      "weight": 1.0
                  },
                  {
                      "datacenterId": 3133,
                      "enabled": true,
                      "handoutCName": "uk.shop.example.com",
                      "name": null,
                      "servers": [],
                      "weight": 1.0
                  }
              ],
              "type": "geographic",
              "unreachableThreshold": null,
              "useComputedTargets": false
          }
      ],
      "resources": [
          {
              "aggregationType": "latest",
              "constrainedProperty": "mail",
              "decayRate": null,
              "description": "CPU utilization",
              "hostHeader": null,
              "leaderString": null,
              "leastSquaresDecay": null,
              "links": [
                  {
                      "href": "/config-gtm/v1/domains/example.akadns.net/resources/cpu",
                      "rel": "self"
                  }
              ],
              "loadImbalancePercentage": null,
              "maxUMultiplicativeIncrement": null,
              "name": "cpu",
              "resourceInstances": [
                  {
                      "datacenterId": 3134,
                      "loadObject": "/cpu",
                      "loadObjectPort": 80,
                      "loadServers": [
                          "1.2.3.8"
                      ],
                      "useDefaultLoadObject": false
                  },
                  {
                      "datacenterId": 3133,
                      "loadObject": "/cpu",
                      "loadObjectPort": 80,
                      "loadServers": [
                          "1.2.3.7"
                      ],
                      "useDefaultLoadObject": false
                  },
                  {
                      "datacenterId": 5400,
                      "loadObject": null,
                      "loadObjectPort": 0,
                      "loadServers": [],
                      "useDefaultLoadObject": false
                  }
              ],
              "type": "XML load object via HTTP",
              "upperBound": 0
          },
          {
              "aggregationType": "latest",
              "constrainedProperty": "**",
              "decayRate": null,
              "description": "Supply levels of Arbor Gold",
              "hostHeader": null,
              "leaderString": null,
              "leastSquaresDecay": null,
              "links": [
                  {
                      "href": "/config-gtm/v1/domains/example.akadns.net/resources/arbor-gold",
                      "rel": "self"
                  }
              ],
              "loadImbalancePercentage": null,
              "maxUMultiplicativeIncrement": null,
              "name": "arbor-gold",
              "resourceInstances": [
                  {
                      "datacenterId": 3134,
                      "loadObject": "/cups",
                      "loadObjectPort": 80,
                      "loadServers": [
                          "1.2.3.8"
                      ],
                      "useDefaultLoadObject": false
                  },
                  {
                      "datacenterId": 3133,
                      "loadObject": "/cups",
                      "loadObjectPort": 80,
                      "loadServers": [
                          "1.2.3.7"
                      ],
                      "useDefaultLoadObject": false
                  },
                  {
                      "datacenterId": 5400,
                      "loadObject": null,
                      "loadObjectPort": 0,
                      "loadServers": [],
                      "useDefaultLoadObject": false
                  }
              ],
              "type": "Non-XML load object via HTTP",
              "upperBound": 0
          }
      ],
      "roundRobinPrefix": null,
      "servermonitorLivenessCount": null,
      "servermonitorLoadCount": null,
      "status": {
          "changeId": "5beb11ae-8908-4bfe-8459-e88efc4d2fdc",
          "links": [
              {
                  "href": "/config-gtm/v1/domains/example.akadns.net/status/current",
                  "rel": "self"
              }
          ],
          "message": "Change Pending",
          "passingValidation": true,
          "propagationStatus": "PENDING",
          "propagationStatusDate": "2014-04-08T18:25:51.000+0000"
      },
      "type": "full"
  }
	`
	domainStatusJson = `{
		"message": "ERROR: Some error",
		"changeId": "123",
		"propagationStatus": "DENIED",
		"propagationStatusDate": "2015-07-31T23:08:00.000+0000",
		"passingValidation": false,
		"links": []
	}`

	updatedDomainJson = `{
    "resource": {
        "links": [
            {
                "href": "/config-gtm/v1/domains/example.akadns.net",
                "rel": "self"
            },
            {
                "href": "/config-gtm/v1/domains/example.akadns.net/datacenters",
                "rel": "datacenters"
            },
            {
                "href": "/config-gtm/v1/domains/example.akadns.net/properties",
                "rel": "properties"
            },
            {
                "href": "/config-gtm/v1/domains/example.akadns.net/geographic-maps",
                "rel": "geographic-maps"
            },
            {
                "href": "/config-gtm/v1/domains/example.akadns.net/cidr-maps",
                "rel": "cidr-maps"
            },
            {
                "href": "/config-gtm/v1/domains/example.akadns.net/resources",
                "rel": "resources"
            }
        ],
        "cidrMaps": [
            {
                "assignments": [
                    {
                        "blocks": [
                            "1.3.5.9",
                            "1.2.3.0/24"
                        ],
                        "datacenterId": 3134,
                        "nickname": "Frostfangs and the Fist of First Men"
                    },
                    {
                        "blocks": [
                            "1.2.4.0/24"
                        ],
                        "datacenterId": 3133,
                        "nickname": "Winterfell"
                    }
                ],
                "defaultDatacenter": {
                    "datacenterId": 5400,
                    "nickname": "All Other CIDR Blocks"
                },
                "links": [
                    {
                        "href": "/config-gtm/v1/domains/example.akadns.net/cidr-maps/The%20North",
                        "rel": "self"
                    }
                ],
                "name": "The North"
            }
        ],
        "datacenters": [
            {
                "city": "Doune",
                "cloneOf": 0,
                "continent": "EU",
                "country": "GB",
                "datacenterId": 3133,
                "defaultLoadObject": {
                    "loadObject": null,
                    "loadObjectPort": 0,
                    "loadServers": null
                },
                "latitude": 56.185097,
                "links": [
                    {
                        "href": "/config-gtm/v1/domains/example.akadns.net/datacenters/3133",
                        "rel": "self"
                    }
                ],
                "longitude": -4.050264,
                "nickname": "Winterfell",
                "stateOrProvince": "Perthshire",
                "virtual": true
            },
            {
                "city": "Snæfellsjökull",
                "cloneOf": 0,
                "continent": "EU",
                "country": "IS",
                "datacenterId": 3134,
                "defaultLoadObject": {
                    "loadObject": null,
                    "loadObjectPort": 0,
                    "loadServers": null
                },
                "latitude": 64.808,
                "links": [
                    {
                        "href": "/config-gtm/v1/domains/example.akadns.net/datacenters/3134",
                        "rel": "self"
                    }
                ],
                "longitude": -23.776,
                "nickname": "Frostfangs",
                "stateOrProvince": null,
                "virtual": true
            },
            {
                "city": null,
                "cloneOf": 0,
                "continent": null,
                "country": null,
                "datacenterId": 5400,
                "defaultLoadObject": {
                    "loadObject": null,
                    "loadObjectPort": 0,
                    "loadServers": null
                },
                "latitude": 0.0,
                "links": [
                    {
                        "href": "/config-gtm/v1/domains/example.akadns.net/datacenters/5400",
                        "rel": "self"
                    }
                ],
                "longitude": 0.0,
                "nickname": "Default Datacenter",
                "stateOrProvince": null,
                "virtual": true
            }
        ],                    
        "defaultSslClientCertificate": null,
        "defaultSslClientPrivateKey": null,
        "defaultUnreachableThreshold": null,
        "emailNotificationList": [],
        "geographicMaps": [
            {
                "assignments": [
                    {
                        "countries": [
                            "GB"
                        ],
                        "datacenterId": 3133,
                        "nickname": "UK users"
                    }
                ],
                "defaultDatacenter": {
                    "datacenterId": 5400,
                    "nickname": "Default Mapping"
                },
                "links": [
                    {
                        "href": "/config-gtm/v1/domains/example.akadns.net/geographic-maps/UK%20Delivery",
                        "rel": "self"
                    }
                ],
                "name": "UK Delivery"
            }
        ],
        "lastModified": "2014-04-08T18:25:51.000+0000",
        "lastModifiedBy": "admin@example.com",
        "loadFeedback": true,
        "loadImbalancePercentage": 10.0,
        "minPingableRegionFraction": null,
        "modificationComments": "CIDRMap example",
        "name": "example.akadns.net",
        "pingInterval": null,
        "properties": [
            {
                "backupCName": null,
                "backupIp": null,
                "balanceByDownloadScore": false,
                "cname": null,
                "comments": null,
                "dynamicTTL": 300,
                "failbackDelay": 0,
                "failoverDelay": 0,
                "handoutMode": "normal",
                "healthMax": null,
                "healthMultiplier": null,
                "healthThreshold": null,
                "ipv6": false,
                "lastModified": "2014-04-08T18:25:52.000+0000",
                "links": [
                    {
                        "href": "/config-gtm/v1/domains/example.akadns.net/properties/www",
                        "rel": "self"
                    }
                ],
                "livenessTests": [
                    {
                        "disableNonstandardPortWarning": false,
                        "hostHeader": "foo.example.com",
                        "httpError3xx": true,
                        "httpError4xx": true,
                        "httpError5xx": true,
                        "links": [],
                        "name": "health-check",
                        "requestString": null,
                        "responseString": null,
                        "sslClientCertificate": null,
                        "sslClientPrivateKey": null,
                        "testInterval": 60,
                        "testObject": "/status",
                        "testObjectPassword": null,
                        "testObjectPort": 80,
                        "testObjectProtocol": "HTTP",
                        "testObjectUsername": null,
                        "testTimeout": 25.0
                    }
                ],
                "mapName": null,
                "maxUnreachablePenalty": null,
                "mxRecords": [],
                "name": "www",
                "scoreAggregationType": "mean",
                "staticTTL": 600,
                "stickinessBonusConstant": 0,
                "stickinessBonusPercentage": 0,
                "trafficTargets": [
                    {
                        "datacenterId": 5400,
                        "enabled": false,
                        "handoutCName": null,
                        "name": null,
                        "servers": [],
                        "weight": 0.0
                    },
                    {
                        "datacenterId": 3134,
                        "enabled": true,
                        "handoutCName": null,
                        "name": null,
                        "servers": [
                            "1.2.3.5"
                        ],
                        "weight": 0.0
                    },
                    {
                        "datacenterId": 3133,
                        "enabled": true,
                        "handoutCName": null,
                        "name": null,
                        "servers": [
                            "1.2.3.4"
                        ],
                        "weight": 1.0
                    }
                ],
                "type": "failover",
                "unreachableThreshold": null,
                "useComputedTargets": false
            },
            {
                "backupCName": null,
                "backupIp": null,
                "balanceByDownloadScore": false,
                "cname": null,
                "comments": null,
                "dynamicTTL": 300,
                "failbackDelay": 0,
                "failoverDelay": 0,
                "handoutMode": "normal",
                "healthMax": null,
                "healthMultiplier": null,
                "healthThreshold": null,
                "ipv6": true,
                "lastModified": "2014-04-08T18:25:52.000+0000",
                "links": [
                    {
                        "href": "/config-gtm/v1/domains/example.akadns.net/properties/mail",
                        "rel": "self"
                    }
                ],
                "livenessTests": [],
                "mapName": null,
                "maxUnreachablePenalty": null,
                "mxRecords": [],
                "name": "mail",
                "scoreAggregationType": "mean",
                "staticTTL": 600,
                "stickinessBonusConstant": 0,
                "stickinessBonusPercentage": 0,
                "trafficTargets": [
                    {
                        "datacenterId": 5400,
                        "enabled": false,
                        "handoutCName": null,
                        "name": null,
                        "servers": [],
                        "weight": 1.0
                    },
                    {
                        "datacenterId": 3134,
                        "enabled": true,
                        "handoutCName": null,
                        "name": null,
                        "servers": [
                            "2001:4878::5043:4078"
                        ],
                        "weight": 1.0
                    },
                    {
                        "datacenterId": 3133,
                        "enabled": true,
                        "handoutCName": null,
                        "name": null,
                        "servers": [
                            "2001:4878::5043:4072",
                            "2001:4878::5043:4071"
                        ],
                        "weight": 1.0
                    }
                ],
                "type": "weighted-round-robin",
                "unreachableThreshold": null,
                "useComputedTargets": false
            },
            {
                "backupCName": null,
                "backupIp": null,
                "balanceByDownloadScore": false,
                "cname": null,
                "comments": null,
                "dynamicTTL": 300,
                "failbackDelay": 0,
                "failoverDelay": 0,
                "handoutMode": "normal",
                "healthMax": null,
                "healthMultiplier": null,
                "healthThreshold": null,
                "ipv6": false,
                "lastModified": "2014-04-08T18:25:52.000+0000",
                "links": [
                    {
                        "href": "/config-gtm/v1/domains/example.akadns.net/properties/supplies",
                        "rel": "self"
                    }
                ],
                "livenessTests": [],
                "loadImbalancePercentage": null,
                "mapName": null,
                "maxUnreachablePenalty": null,
                "mxRecords": [],
                "name": "supplies",
                "scoreAggregationType": "mean",
                "staticTTL": 600,
                "stickinessBonusConstant": 0,
                "stickinessBonusPercentage": 0,
                "trafficTargets": [
                    {
                        "datacenterId": 5400,
                        "enabled": true,
                        "handoutCName": "supplies.example.com",
                        "name": null,
                        "servers": [],
                        "weight": 1.0
                    },
                    {
                        "datacenterId": 3134,
                        "enabled": true,
                        "handoutCName": "winter.supplies.example.com",
                        "name": null,
                        "servers": [],
                        "weight": 0.0
                    },
                    {
                        "datacenterId": 3133,
                        "enabled": true,
                        "handoutCName": "redcross.org",
                        "name": null,
                        "servers": [],
                        "weight": 0.0
                    }
                ],
                "type": "failover",
                "unreachableThreshold": null,
                "useComputedTargets": false
            },
            {
                "backupCName": null,
                "backupIp": null,
                "balanceByDownloadScore": false,
                "cname": null,
                "comments": null,
                "dynamicTTL": 300,
                "failbackDelay": 0,
                "failoverDelay": 0,
                "handoutMode": "normal",
                "healthMax": null,
                "healthMultiplier": null,
                "healthThreshold": null,
                "ipv6": false,
                "lastModified": "2014-04-08T18:25:52.000+0000",
                "links": [
                    {
                        "href": "/config-gtm/v1/domains/example.akadns.net/properties/shop",
                        "rel": "self"
                    }
                ],
                "livenessTests": [],
                "loadImbalancePercentage": null,
                "mapName": "UK Delivery",
                "maxUnreachablePenalty": null,
                "mxRecords": [],
                "name": "shop",
                "scoreAggregationType": "mean",
                "staticTTL": 600,
                "stickinessBonusConstant": 0,
                "stickinessBonusPercentage": 0,
                "trafficTargets": [
                    {
                        "datacenterId": 5400,
                        "enabled": true,
                        "handoutCName": "shop.example.com",
                        "name": null,
                        "servers": [],
                        "weight": 1.0
                    },
                    {
                        "datacenterId": 3134,
                        "enabled": false,
                        "handoutCName": null,
                        "name": null,
                        "servers": [],
                        "weight": 1.0
                    },
                    {
                        "datacenterId": 3133,
                        "enabled": true,
                        "handoutCName": "uk.shop.example.com",
                        "name": null,
                        "servers": [],
                        "weight": 1.0
                    }
                ],
                "type": "geographic",
                "unreachableThreshold": null,
                "useComputedTargets": false
            }
        ],
        "resources": [
            {
                "aggregationType": "latest",
                "constrainedProperty": "mail",
                "decayRate": null,
                "description": "CPU utilization",
                "hostHeader": null,
                "leaderString": null,
                "leastSquaresDecay": null,
                "links": [
                    {
                        "href": "/config-gtm/v1/domains/example.akadns.net/resources/cpu",
                        "rel": "self"
                    }
                ],
                "loadImbalancePercentage": null,
                "maxUMultiplicativeIncrement": null,
                "name": "cpu",
                "resourceInstances": [
                    {
                        "datacenterId": 3134,
                        "loadObject": "/cpu",
                        "loadObjectPort": 80,
                        "loadServers": [
                            "1.2.3.8"
                        ],
                        "useDefaultLoadObject": false
                    },
                    {
                        "datacenterId": 3133,
                        "loadObject": "/cpu",
                        "loadObjectPort": 80,
                        "loadServers": [
                            "1.2.3.7"
                        ],
                        "useDefaultLoadObject": false
                    },
                    {
                        "datacenterId": 5400,
                        "loadObject": null,
                        "loadObjectPort": 0,
                        "loadServers": [],
                        "useDefaultLoadObject": false
                    }
                ],
                "type": "XML load object via HTTP",
                "upperBound": 0
            },
            {
                "aggregationType": "latest",
                "constrainedProperty": "**",
                "decayRate": null,
                "description": "Supply levels of Arbor Gold",
                "hostHeader": null,
                "leaderString": null,
                "leastSquaresDecay": null,
                "links": [
                    {
                        "href": "/config-gtm/v1/domains/example.akadns.net/resources/arbor-gold",
                        "rel": "self"
                    }
                ],
                "loadImbalancePercentage": null,
                "maxUMultiplicativeIncrement": null,
                "name": "arbor-gold",
                "resourceInstances": [
                    {
                        "datacenterId": 3134,
                        "loadObject": "/cups",
                        "loadObjectPort": 80,
                        "loadServers": [
                            "1.2.3.8"
                        ],
                        "useDefaultLoadObject": false
                    },
                    {
                        "datacenterId": 3133,
                        "loadObject": "/cups",
                        "loadObjectPort": 80,
                        "loadServers": [
                            "1.2.3.7"
                        ],
                        "useDefaultLoadObject": false
                    },
                    {
                        "datacenterId": 5400,
                        "loadObject": null,
                        "loadObjectPort": 0,
                        "loadServers": [],
                        "useDefaultLoadObject": false
                    }
                ],
                "type": "Non-XML load object via HTTP",
                "upperBound": 0
            }
        ],
        "roundRobinPrefix": null,
        "servermonitorLivenessCount": null,
        "servermonitorLoadCount": null,
        "status": {
            "changeId": "5beb11ae-8908-4bfe-8459-e88efc4d2fdc",
            "links": [
                {
                    "href": "/config-gtm/v1/domains/example.akadns.net/status/current",
                    "rel": "self"
                }
            ],
            "message": "Change Pending",
            "passingValidation": true,
            "propagationStatus": "PENDING",
            "propagationStatusDate": "2014-04-08T18:25:51.000+0000"
        },
        "type": "full"
    },
    "status": {
        "changeId": "5beb11ae-8908-4bfe-8459-e88efc4d2fdc",
        "links": [
            {
                "href": "/config-gtm/v1/domains/example.akadns.net/status/current",
                "rel": "self"
            }
        ],
        "message": "Change Pending",
        "passingValidation": true,
        "propagationStatus": "PENDING",
        "propagationStatusDate": "2014-04-08T18:25:51.000+0000"
    }
  }`

	dataCenterJson = `{
		"nickname":"dcNickname",
		"city":"city",
		"stateOrProvince":"state",
		"country":"country",
		"latitude":0.0,
		"longitude":0.0,
		"cloneOf":null,
		"virtual":false,
		"defaultLoadObject":{
			"loadObject":null,
			"loadObjectPort":0,
			"loadServers":null
		},
		"continent":null,
		"cloudServerTargeting":false,
		"links":[{
			"rel":"self",
			"href":"https://some-url.com"
		}],
		"datacenterId":1
	}`

	createdDataCenterJson = `{
	  "resource": {
	    "city": "Doune",
	    "cloneOf": 0,
	    "continent": "EU",
	    "country": "GB",
	    "datacenterId": 3133,
	    "defaultLoadObject": {
	      "loadObject": null,
	      "loadObjectPort": 0,
	      "loadServers": null
	    },
	    "latitude": 56.185097,
	    "links": [
	      {
	        "href": "/config-gtm/v1/domains/example.akadns.net/datacenters/3133",
	        "rel": "self"
	      }
	    ],
	    "longitude": -4.050264,
	    "nickname": "Winterfell",
	    "stateOrProvince": "Perthshire",
	    "virtual": true
	  },
	  "status": {
	    "changeId": "f0c51967-d119-4665-9403-364a57ea5530",
	    "links": [
	      {
	        "href": "/config-gtm/v1/domains/example.akadns.net/status/current",
	        "rel": "self"
	      }
	    ],
	    "message": "Change Pending",
	    "passingValidation": true,
	    "propagationStatus": "PENDING",
	    "propagationStatusDate": "2014-04-15T11:30:27.000+0000"
	  }
	}`

	dataCentersJson = `{
		"items":[{
			"nickname":"dcOne",
			"city":"Some City",
			"stateOrProvince":"someState",
			"country":"someCountry",
			"latitude":0,
			"longitude":0,
			"cloneOf":null,
			"virtual":false,
			"defaultLoadObject":{
				"loadObject":null,
				"loadObjectPort":0,
				"loadServers":null
			},
			"continent":null,
			"cloudServerTargeting":false,
			"links":[{
				"rel":"self",
				"href":"https://url.com"
			}],
			"datacenterId":123
		}, {
			"nickname":"dc_two",
			"city":"Another City",
			"stateOrProvince":"anotherState",
			"country":"anotherCountry",
			"latitude":0,
			"longitude":0,
			"cloneOf":null,
			"virtual":false,
			"defaultLoadObject":{
				"loadObject":null,
				"loadObjectPort":0,
				"loadServers":null
			},
			"continent":null,
			"cloudServerTargeting":false,
			"links":[{
				"rel":"self",
				"href":"https://another-url.com"
			}],
			"datacenterId":1234
		}]
	}`

	createdPropertyJson = `{
	  "resource": {
	    "backupCName": null,
	    "backupIp": null,
	    "balanceByDownloadScore": false,
	    "cname": null,
	    "comments": null,
	    "dynamicTTL": 300,
	    "failbackDelay": 0,
	    "failoverDelay": 0,
	    "handoutMode": "normal",
	    "healthMax": null,
	    "healthMultiplier": null,
	    "healthThreshold": null,
	    "ipv6": false,
	    "lastModified": null,
	    "links": [
	      {
	        "href": "/config-gtm/v1/domains/example.akadns.net/properties/origin",
	        "rel": "self"
	      }
	    ],
	    "livenessTests": [
	      {
	        "disableNonstandardPortWarning": false,
	        "hostHeader": "foo.example.com",
	        "httpError3xx": true,
	        "httpError4xx": true,
	        "httpError5xx": true,
	        "name": "health-check",
	        "requestString": null,
	        "responseString": null,
	        "sslClientCertificate": null,
	        "sslClientPrivateKey": null,
	        "testInterval": 60,
	        "testObject": "/status",
	        "testObjectPassword": null,
	        "testObjectPort": 80,
	        "testObjectProtocol": "HTTP",
	        "testObjectUsername": null,
	        "testTimeout": 25
	      }
	    ],
	    "loadImbalancePercentage": null,
	    "mapName": null,
	    "maxUnreachablePenalty": null,
	    "mxRecords": [],
	    "name": "origin",
	    "scoreAggregationType": "mean",
	    "staticTTL": 600,
	    "stickinessBonusConstant": 0,
	    "stickinessBonusPercentage": 0,
	    "trafficTargets": [
	      {
	        "datacenterId": 3134,
	        "enabled": true,
	        "handoutCName": null,
	        "name": null,
	        "servers": [
	          "1.2.3.5"
	        ],
	        "weight": 50
	      },
	      {
	        "datacenterId": 3133,
	        "enabled": true,
	        "handoutCName": null,
	        "name": null,
	        "servers": [
	          "1.2.3.4"
	        ],
	        "weight": 50
	      }
	    ],
	    "type": "weighted-round-robin",
	    "unreachableThreshold": null,
	    "useComputedTargets": false
	  },
	  "status": {
	    "changeId": "eee0c3b4-0e45-4f4b-822c-7dbc60764d18",
	    "links": [
	      {
	        "href": "/config-gtm/v1/domains/example.akadns.net/status/current",
	        "rel": "self"
	      }
	    ],
	    "message": "Change Pending",
	    "passingValidation": true,
	    "propagationStatus": "PENDING",
	    "propagationStatusDate": "2014-04-15T11:30:27.000+0000"
	  }
	}`

	propJson = `{
		"backupCName":null,
		"backupIp":null,
		"balanceByDownloadScore":false,
		"cname":null,
		"comments":null,
		"dynamicTTL":300,
		"failoverDelay":0,
		"failbackDelay":0,
		"handoutMode":"normal",
		"healthMax":null,
		"healthMultiplier":null,
		"healthThreshold":null,
		"lastModified":null,
		"livenessTests":[{
			"disableNonstandardPortWarning":false,
			"hostHeader":null,
			"httpError3xx":true,
			"httpError4xx":true,
			"httpError5xx":true,
			"name":"livenessTestOne",
			"requestString":null,
			"responseString":null,
			"testInterval":60,
			"testObject":"/",
			"testObjectPort":443,
			"testObjectProtocol":"HTTPS",
			"testObjectUsername":null,
			"testObjectPassword":null,
			"testTimeout":25.0,
			"sslClientCertificate":null,
			"sslClientPrivateKey":null
		}],
		"loadImbalancePercentage":null,
		"mapName":null,
		"maxUnreachablePenalty":null,
		"mxRecords":[],
		"scoreAggregationType":"mean",
		"stickinessBonusConstant":null,
		"stickinessBonusPercentage":null,
		"staticTTL":null,
		"trafficTargets":[{
			"datacenterId":3131,
			"enabled":true,
			"weight":50.0,
			"handoutCName":null,
			"name":null,
			"servers":["1.2.3.5"]
		},{
			"datacenterId":3132,
			"enabled":true,
			"weight":50.0,
			"handoutCName":null,
			"name":null,
			"servers":["1.2.3.4"]
		}],
		"type":"weighted-round-robin",
		"unreachableThreshold":null,
		"useComputedTargets":false,
		"ipv6":false,
		"links":[{
			"rel":"self",
			"href":"https://url.com"
		}],
		"name":"someName"
	}`

	errorResponseJson = `{
    "type": "https://problems.luna.akamaiapis.net/config-gtm/v1/serverError",
    "title": "Server Error",
    "status": 500,
    "detail": "Could not read JSON: parse exception"
  }`

	propertiesResponseJson = `{
		"items":[
			{
				"backupCName":null,
				"backupIp":null,
				"balanceByDownloadScore":false,
				"cname":null,
				"comments":null,
				"dynamicTTL":300,
				"failoverDelay":0,
				"failbackDelay":0,
				"handoutMode":"normal",
				"healthMax":null,
				"healthMultiplier":null,
				"healthThreshold":null,
				"lastModified":null,
				"livenessTests":[

				],
				"loadImbalancePercentage":null,
				"mapName":null,
				"maxUnreachablePenalty":null,
				"mxRecords":[

				],
				"scoreAggregationType":"mean",
				"stickinessBonusConstant":0,
				"stickinessBonusPercentage":0,
				"staticTTL":600,
				"trafficTargets":[
					{
						"datacenterId":3131,
						"enabled":true,
						"weight":70.0,
						"handoutCName":null,
						"name":null,
						"servers":[
							"1.2.3.5"
						]
					},
					{
						"datacenterId":3132,
						"enabled":true,
						"weight":30.0,
						"handoutCName":null,
						"name":null,
						"servers":[
							"1.2.3.4"
						]
					}
				],
				"type":"weighted-round-robin",
				"unreachableThreshold":null,
				"useComputedTargets":false,
				"ipv6":false,
				"links":[
					{
						"rel":"self",
						"href":"https://akab-7zefyjeiytziublw-o4tcc67vhvt7ewwd.luna.akamaiapis.net/config-gtm/v1/domains/example.akadns.net/properties/newprop"
					}
				],
				"name":"newprop"
			},
			{
				"backupCName":null,
				"backupIp":null,
				"balanceByDownloadScore":false,
				"cname":null,
				"comments":null,
				"dynamicTTL":300,
				"failoverDelay":0,
				"failbackDelay":0,
				"handoutMode":"normal",
				"healthMax":null,
				"healthMultiplier":null,
				"healthThreshold":null,
				"lastModified":null,
				"livenessTests":[

				],
				"loadImbalancePercentage":null,
				"mapName":null,
				"maxUnreachablePenalty":null,
				"mxRecords":[

				],
				"scoreAggregationType":"mean",
				"stickinessBonusConstant":0,
				"stickinessBonusPercentage":0,
				"staticTTL":600,
				"trafficTargets":[
					{
						"datacenterId":3131,
						"enabled":true,
						"weight":70.0,
						"handoutCName":null,
						"name":null,
						"servers":[
							"1.2.3.5"
						]
					},
					{
						"datacenterId":3132,
						"enabled":true,
						"weight":30.0,
						"handoutCName":null,
						"name":null,
						"servers":[
							"1.2.3.4"
						]
					}
				],
				"type":"weighted-round-robin",
				"unreachableThreshold":null,
				"useComputedTargets":false,
				"ipv6":false,
				"links":[
					{
						"rel":"self",
						"href":"https://akab-7zefyjeiytziublw-o4tcc67vhvt7ewwd.luna.akamaiapis.net/config-gtm/v1/domains/example.akadns.net/properties/someprop"
					}
				],
				"name":"someprop"
			},
			{
				"backupCName":null,
				"backupIp":null,
				"balanceByDownloadScore":false,
				"cname":null,
				"comments":null,
				"dynamicTTL":300,
				"failoverDelay":0,
				"failbackDelay":0,
				"handoutMode":"normal",
				"healthMax":0.0,
				"healthMultiplier":0.0,
				"healthThreshold":0.0,
				"lastModified":null,
				"livenessTests":[

				],
				"loadImbalancePercentage":null,
				"mapName":null,
				"maxUnreachablePenalty":null,
				"mxRecords":[

				],
				"scoreAggregationType":"mean",
				"stickinessBonusConstant":0,
				"stickinessBonusPercentage":0,
				"staticTTL":600,
				"trafficTargets":[
					{
						"datacenterId":3131,
						"enabled":true,
						"weight":50.0,
						"handoutCName":null,
						"name":null,
						"servers":[
							"1.2.3.5"
						]
					},
					{
						"datacenterId":3132,
						"enabled":true,
						"weight":50.0,
						"handoutCName":null,
						"name":null,
						"servers":[
							"1.2.3.4"
						]
					}
				],
				"type":"weighted-round-robin",
				"unreachableThreshold":null,
				"useComputedTargets":false,
				"ipv6":false,
				"links":[
					{
						"rel":"self",
						"href":"https://akab-7zefyjeiytziublw-o4tcc67vhvt7ewwd.luna.akamaiapis.net/config-gtm/v1/domains/example.akadns.net/properties/prop"
					}
				],
				"name":"prop"
			},
			{
				"backupCName":null,
				"backupIp":null,
				"balanceByDownloadScore":false,
				"cname":null,
				"comments":null,
				"dynamicTTL":300,
				"failoverDelay":0,
				"failbackDelay":0,
				"handoutMode":"normal",
				"healthMax":null,
				"healthMultiplier":null,
				"healthThreshold":null,
				"lastModified":null,
				"livenessTests":[

				],
				"loadImbalancePercentage":null,
				"mapName":null,
				"maxUnreachablePenalty":null,
				"mxRecords":[

				],
				"scoreAggregationType":"mean",
				"stickinessBonusConstant":0,
				"stickinessBonusPercentage":0,
				"staticTTL":600,
				"trafficTargets":[
					{
						"datacenterId":3131,
						"enabled":true,
						"weight":70.0,
						"handoutCName":null,
						"name":null,
						"servers":[
							"1.2.3.5"
						]
					},
					{
						"datacenterId":3132,
						"enabled":true,
						"weight":30.0,
						"handoutCName":null,
						"name":null,
						"servers":[
							"1.2.3.4"
						]
					}
				],
				"type":"weighted-round-robin",
				"unreachableThreshold":null,
				"useComputedTargets":false,
				"ipv6":false,
				"links":[
					{
						"rel":"self",
						"href":"https://akab-7zefyjeiytziublw-o4tcc67vhvt7ewwd.luna.akamaiapis.net/config-gtm/v1/domains/example.akadns.net/properties/anotherprop"
					}
				],
				"name":"anotherprop"
			},
			{
				"backupCName":null,
				"backupIp":null,
				"balanceByDownloadScore":false,
				"cname":null,
				"comments":null,
				"dynamicTTL":300,
				"failoverDelay":0,
				"failbackDelay":0,
				"handoutMode":"normal",
				"healthMax":0.0,
				"healthMultiplier":0.0,
				"healthThreshold":0.0,
				"lastModified":null,
				"livenessTests":[

				],
				"loadImbalancePercentage":null,
				"mapName":null,
				"maxUnreachablePenalty":null,
				"mxRecords":[

				],
				"scoreAggregationType":"mean",
				"stickinessBonusConstant":0,
				"stickinessBonusPercentage":0,
				"staticTTL":600,
				"trafficTargets":[
					{
						"datacenterId":3131,
						"enabled":true,
						"weight":50.0,
						"handoutCName":null,
						"name":null,
						"servers":[
							"1.2.3.5"
						]
					},
					{
						"datacenterId":3132,
						"enabled":true,
						"weight":50.0,
						"handoutCName":null,
						"name":null,
						"servers":[
							"1.2.3.4"
						]
					}
				],
				"type":"weighted-round-robin",
				"unreachableThreshold":null,
				"useComputedTargets":false,
				"ipv6":false,
				"links":[
					{
						"rel":"self",
						"href":"https://akab-7zefyjeiytziublw-o4tcc67vhvt7ewwd.luna.akamaiapis.net/config-gtm/v1/domains/example.akadns.net/properties/prop"
					}
				],
				"name":"aprop"
			}
		]
	}`
)
