package slb

import (
	"fmt"

	"encoding/xml"

	"bytes"

	"github.com/josh5276/brocade-adx-client/brocade"
	"github.com/josh5276/brocade-adx-client/brocade/soap"
)

func New(a adc.ADCSoapClient) SLB {
	return SLB{
		adc: a,
	}
}

// Slb method will display and configure basic system management functions on the
// ServerIron ADX device. Initialize with an Sys struct, then call with a sys call.
func (s SLB) Slb(method string) (*SLB, int, error) {
	payload, err := xmlGetMarshal(method)
	if err != nil {
		return nil, 0, err
	}
	resp, code, err := soap.XMLADXBasicAuthPost(s.adc.URL+"/WS/SLB", method, payload, s.adc.User, s.adc.Passwd)
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
func xmlUnmarshal(resp []byte) (*SLB, error) {
	var e SLB
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
