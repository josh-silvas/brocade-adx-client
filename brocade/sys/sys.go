package sys

import (
	"fmt"

	"encoding/xml"

	"bytes"

	"errors"
	"github.com/josh5276/brocade-adx-client/brocade"
	"github.com/josh5276/brocade-adx-client/brocade/soap"
)

func New(a adc.ADCSoapClient) SYS {
	return SYS{
		adc: a,
	}
}

// Sys method will display and configure basic system management functions on the
// ServerIron ADX device. Initialize with an ADX struct, then call with a sys call.
func (s SYS) Sys(method string) (*SYS, int, error) {
	payload, err := xmlGetMarshal(method)
	if err != nil {
		return nil, 0, err
	}
	resp, code, err := soap.XMLADXBasicAuthPost(s.adc.URL+"/WS/SYS", method, payload, s.adc.User, s.adc.Passwd)
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
func (s SYS) SysRunCli(commands []string) (*SYS, int, error) {
	sys := &RunCliRequest{
		Soap: "http://schemas.xmlsoap.org/soap/envelope/",
		RunCLI: RunCliCommands{
			Tns:            "urn:webservicesapi",
			StringSequence: commands,
		},
	}
	payload := soap.XMLMarshalHead(sys)
	if payload == "" {
		return nil, 0, errors.New("Unable to successfully marshal a XML request")
	}
	resp, code, err := soap.XMLADXBasicAuthPost(s.adc.URL+"/WS/SYS", "runCLI", payload, s.adc.User, s.adc.Passwd)
	if err != nil {
		return nil, code, fmt.Errorf("XMLADXBasicAuthPost Error: %s", err.Error())
	}
	e, err := xmlUnmarshal(resp)
	if err != nil {
		return nil, code, err
	}
	return e, code, nil
}

// TestAuth method is designed to connect to an ServerIron ADX and return the current
// running version.  Call will also return the fault ID if login was unsuccessful.
func (s SYS) TestAuth() (string, error) {
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

// xmlUnmarshal function will take a byte array and formulate a SOAP Envelope
// structure.  Return values will be the Envelope pointer value and error.
func xmlUnmarshal(resp []byte) (*SYS, error) {
	var e SYS
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
