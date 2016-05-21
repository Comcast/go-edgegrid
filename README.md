# edgegrid

[![Build Status](https://travis-ci.org/Comcast/go-edgegrid.svg?branch=master)](https://travis-ci.org/Comcast/go-edgegrid)

A Golang [Akamai](https://developer.akamai.com/api/) API client.

[Akamai {OPEN} EdgeGrid Authentication](https://developer.akamai.com/introduction/Client_Auth.html) for Golang.

`edgegrid` implements the [Akamai {OPEN} EdgeGrid Authentication](https://developer.akamai.com/introduction/Client_Auth.html) scheme for the Golang `net/http` package, and also offers a convenience API for interacting with the Akamai [GTM](https://developer.akamai.com/api/luna/config-gtm/overview.html) and [PAPI](https://developer.akamai.com/api/luna/papi/overview.html) APIs.

## Installation

Assuming Golang is properly installed:

```bash
go get github.com/comcast/go-edgegrid/edgegrid
```

## Basic Usage

Basic usage assumes you've established up the following environment variables.

```bash
AKAMAI_EDGEGRID_HOST=
AKAMAI_EDGEGRID_CLIENT_TOKEN=
AKAMAI_EDGEGRID_ACCESS_TOKEN=
AKAMAI_EDGEGRID_CLIENT_SECRET=
```

To log request/response details:

```bash
AK_LOG=true
```

### GTMClient

```go
import "github.com/comcast/go-edgegrid/edgegrid"

client := edgegrid.NewGTMClient()

// get all domains
client.Domains()

// get domain
client.Domain("example.akadns.net")

// create domain
domain := &Domain{
  // your data here; see resources.go
}
client.DomainCreate(domain)

// update domain
domain := &Domain{
  // your data here; see gtm_resources.go
}
client.DomainUpdate(domain)

// delete domain
// NOTE: this may 403, as it appears Akamai professional services is needed to perform DELETEs
client.DomainDelete("domain")

// get datacenters info for domain
client.DataCenters("domain")

// get datacenter info by dc id
client.DataCenter("domain", "123")

// get datacenter info by dc nickname
client.DataCenterByName("domain", "some_nickname")

// create datacenter
dc := &DataCenter{
  // your data here; see gtm_resources.go
}
client.DataCenterCreate("domain", dc)

// update datacenter
dc := &DataCenter{
  // your data here; see gtm_resources.go
}
client.DataCenterUpdate("domain", dc)

// delete datacenter by id
client.DataCenterDelete("domain", "123")

// get property info
client.Property("domain", "property")

// create property
prop := &Property{
  // your data here; see gtm_resources.go
}
client.PropertyCreate("domain", prop)

// adjust aspects of a property
prop := &Property{
  // your data here; see gtm_resources.go
}
client.PropertyUpdate("domain", prop)

// delete a property
client.PropertyDelete("domain", "property")
```

### PAPIClient

```go
import "github.com/comcast/go-edgegrid/edgegrid"

client := edgegrid.NewPAPIClient()

// get all PAPI groups
client.Groups()

// get all PAPI products
client.Groups("someContractId")

// get all PAPI CP codes for a contract and group
client.CpCodes("someContractId", "someGroupID")

// get a PAPI CP code
client.CpCode("someCPCodeID", "someContractId", "someGroupID")

// get all PAPI hostnames for a contract & group
client.Hostnames("someContractId", "someGroupID")

// get a PAPI hostname for a contract & group
client.Hostnames("someHostID", "someContractId", "someGroupID")

// get all PAPI properties for a contract & group
client.Properties("someContractId", "someGroupID")

// get a PAPI property for a contract & group
client.Property("somePropertyID", "someContractId", "someGroupID")

// get all PAPI property versions for a property, contract, & group
client.Versions("somePropertyID", "someContractId", "someGroupID")

// get a PAPI property version for a property, contract, & group
client.Version("somePropertyVersion", "somePropertyID", "someContractId", "someGroupID")

// get a PAPI property version XML for a property, contract, & group
client.PropertyVersionXML("somePropertyVersion", "somePropertyID", "someContractId", "someGroupID")

// get the latest PAPI property version for a property, contract, & group
client.PropertyLatestVersion("somePropertyID", "someContractId", "someGroupID")

// get the PAPI property rules for a property, contract, & group
client.PropertyRules("somePropertyID", "someContractId", "someGroupID")

// get the PAPI property activations for a property, contract, & group
client.PropertyActivations("somePropertyID", "someContractId", "someGroupID")
```

### Alternative Usage

In addition to the `edgegrid.GTMClient` and `edgegrid.PAPIClient` convenience clients, `edgegrid.Auth` offers a standalone implementation of the [Akamai {OPEN} EdgeGrid Authentication](https://developer.akamai.com/introduction/Client_Auth.html) scheme for the Golang net/http library and can be used independent of the `GTM` and `PAPI` convenience clients.

`edgegrid.Auth` offers a utility for establishing a valid `Authorization` header for use with authenticated requests to Akamai using your own HTTP client.

Example:

```go
package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"

  "github.com/comcast/go-edgegrid/edgegrid"
)

func main() {
  client := &http.Client{}
  req, _ := http.NewRequest("GET", "<SOME_AKAMAI_URL>", nil)
  accessToken := os.Getenv("AKAMAI_EDGEGRID_ACCESS_TOKEN")
  clientToken := os.Getenv("AKAMAI_EDGEGRID_CLIENT_TOKEN")
  clientSecret := os.Getenv("AKAMAI_EDGEGRID_CLIENT_SECRET")
  params := edgegrid.NewAuthParams(req, accessToken, clientToken, clientSecret)
  auth := edgegrid.Auth(params)

  req.Header.Add("Authorization", auth)

  resp, _ := client.Do(req)
  contents, _ := ioutil.ReadAll(resp.Body)

  fmt.Printf("%s\n", string(contents))
}
```

## Run tests

Tests use the built-in `testing` package:

```bash
make test
```
