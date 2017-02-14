# brocade-adx-client

[![GoDoc](https://godoc.org/github.com/josh5276/brocade-adx-client/brocade?status.png)](http://godoc.org/github.com/josh5276/brocade-adx-client/brocade)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg?style=flat)](https://github.com/josh5276/brocade-adx-client/blob/master/LICENSE)

The brocade-adx-client is a GoLang client library to interface with a Brocade ServerIron ADX via the SOAP API

## Installation

```
go get -u github.com/josh5276/brocade-adx-client/brocade
```


## Usage

```go
package main

import (
	"log"

	"github.com/josh5276/brocade-adx-client/brocade/sys"
    "github.com/josh5276/brocade-adx-client/brocade"
)

func main() {
	// Establish a new ADX device
	adx := adc.NewSOAPClient("ADX_IP", "Username", "Password")
	
	// Setup the SYS client to make SYS related calls
	s := sys.New(adx)
	
	// Test authentication 
    resp, err := s.TestAuth()
    if err != nil {
    	log.Fatal(err)
    }
    log.Printf("Successful response from the ADX. ServerIron ADX Version: %v", resp)
}
```


## License

This software is released under the MIT License. The full terms of that
license can be found in `LICENSE` file of this repository.
