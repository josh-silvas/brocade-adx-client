package adc

import (
	"fmt"
	"errors"

	"github.com/josh5276/brocade-adx-client/brocade/sys"
)

// A Client manages communication with the ADX API.
type ADCSoapClient struct {
	URL string
	User    string
	Passwd  string
}

type ADC struct {
	c ADCSoapClient
	Sys  sys.Sys
}

// NewSOAPClient function will initialize the structure needed to connect to the ServerIron
// ADX platform over its SOAP API. Function takes the system IP address, and the credentials
// to pass to the appliance
func NewSOAPClient(ip, username, password string) *ADCSoapClient {
	return &ADCSoapClient{
		URL: fmt.Sprintf("https://%v", ip),
		User:    username,
		Passwd:  password,
	}
}

// NewADC function will initialize the NewSOAPClient and return an ADC structure with
// the sub-elements initialized.
func NewADC(ip, username, password string) (*ADCSoapClient, error) {
	sc := NewSOAPClient(ip, username, password)
	return &ADC{
		c: sc,
		Sys: sys.New(sc),
	}, nil
}

// TestAuth method is designed to connect to an ServerIron ADX and return the current
// running version.  Call will also return the fault ID if login was unsuccessful.
func TestAuth(adc *ADCSoapClient) (string, error) {
	s := sys.New(adc)
	r, code, err := s.Sys("getVersion")
	if err != nil {
		return "", err
	}
	if code == 403 {
		return "", fmt.Errorf("Invalid username or password: Unauthorized %v", code)
	}
	if code != 200 {
		return "", fmt.Errorf("Non 200 Code received from the ServerIron ADX: %v", code)
	}
	if r.Body.Msg != nil {
		return "", errors.New(r.Body.Msg.FaultId)
	}
	if r.Body.Version == nil {
		return "", errors.New("Unable to determine version")
	}
	return r.Body.Version.Version, nil
}

