package brocade

import (
	"fmt"

	"encoding/xml"

	"bytes"

	"errors"

	"github.com/josh5276/brocade-adx-client/webutil"
	"github.com/josh5276/brocade-adx-client/adx/sys"
)

// A Client manages communication with the ADX API.
type ADXSoapClient struct {
	baseURL string
	user    string
	passwd  string
}

// NewSOAPClient function will initialize the structure needed to connect to the ServerIron
// ADX platform over its SOAP API. Function takes the system IP address, and the credentials
// to pass to the appliance
func NewSOAPClient(ip, username, password string) (*ADXSoapClient, error) {
	return &ADXSoapClient{
		baseURL: fmt.Sprintf("https://%v", ip),
		user:    username,
		passwd:  password,
	}, nil
}

// TestAuth method is designed to connect to an ServerIron ADX and return the current
// running version.  Call will also return the fault ID if login was unsuccessful.
func (adx *ADXSoapClient) TestAuth() (string, error) {
	r, code, err := sys.Sys("getVersion")
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

// Sys method will display and configure basic system management functions on the
// ServerIron ADX device. Initialize with an ADX struct, then call with a sys call.
func (adx *ADXSoapClient) Sys(method string) (*Envelope, int, error) {
	payload, err := xmlGetMarshal(method)
	if err != nil {
		return nil, 0, err
	}
	resp, code, err := webutil.XMLADXBasicAuthPost(adx.baseURL+"/WS/SYS", method, payload, adx.user, adx.passwd)
	if err != nil {
		return nil, code, fmt.Errorf("XMLADXBasicAuthPost Error: %s", err.Error())
	}
	e, err := xmlUnmarshal(resp)
	if err != nil {
		return nil, code, err
	}
	return e, code, nil
}

// Sys method will display and configure basic system management functions on the
// ServerIron ADX device. Initialize with an ADX struct, then call with a sys call.
func (adx *ADXSoapClient) SysRunCli(commands []string) (*Envelope, int, error) {
	s := &RunCliRequest{
		Soap: "http://schemas.xmlsoap.org/soap/envelope/",
		RunCLI: RunCliCommands{
			Tns:            "urn:webservicesapi",
			StringSequence: commands,
		},
	}
	payload := webutil.XMLMarshalHead(s)
	if payload == "" {
		return nil, 0, errors.New("Unable to successfully marshal a XML request")
	}
	resp, code, err := webutil.XMLADXBasicAuthPost(adx.baseURL+"/WS/SYS", "runCLI", payload, adx.user, adx.passwd)
	if err != nil {
		return nil, code, fmt.Errorf("XMLADXBasicAuthPost Error: %s", err.Error())
	}
	e, err := xmlUnmarshal(resp)
	if err != nil {
		return nil, code, err
	}
	return e, code, nil
}

// Slb method will display and configure basic system management functions on the
// ServerIron ADX device. Initialize with an ADX struct, then call with a sys call.
func (adx *ADXSoapClient) Slb(method string) (*Envelope, int, error) {
	payload, err := xmlGetMarshal(method)
	if err != nil {
		return nil, 0, err
	}
	resp, code, err := webutil.XMLADXBasicAuthPost(adx.baseURL+"/WS/SLB", method, payload, adx.user, adx.passwd)
	if err != nil {
		return nil, code, fmt.Errorf("XMLADXBasicAuthPost Error: %s", err.Error())
	}
	e, err := xmlUnmarshal(resp)
	if err != nil {
		return nil, code, err
	}
	return e, code, nil
}

// xmlUnmarshal function will take a byte array and formulate a SOAP Envelope
// structure.  Return values will be the Envelope pointer value and error.
func xmlUnmarshal(resp []byte) (*Envelope, error) {
	var e Envelope
	parser := xml.NewDecoder(bytes.NewBuffer(resp))
	err := parser.DecodeElement(&e, nil)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

// xmlMarshal is a simple method to format a SOAP call for ServerIron ADX GET
// request.
func xmlGetMarshal(call string) (string, error) {
	retVal := []byte(xml.Header)
	envByte, err := xml.Marshal(
		&XmlGet{
			Soap: "http://schemas.xmlsoap.org/soap/envelope/",
			Call: GetCall{
				XMLName: xml.Name{
					Local: fmt.Sprintf("tns:%s", call),
				},
				Tns: "urn:webservicesapi",
			},
		},
	)
	if err != nil {
		return "", err
	}
	retVal = append(retVal, envByte...)
	return string(retVal), nil
}
