package adc

import (
	"fmt"

)

// A Client manages communication with the ADX API.
type ADCSoapClient struct {
	URL string
	User    string
	Passwd  string
}


// NewSOAPClient function will initialize the structure needed to connect to the ServerIron
// ADX platform over its SOAP API. Function takes the system IP address, and the credentials
// to pass to the appliance
func NewSOAPClient(ip, username, password string) ADCSoapClient {
	return ADCSoapClient{
		URL: fmt.Sprintf("https://%v", ip),
		User:    username,
		Passwd:  password,
	}
}



