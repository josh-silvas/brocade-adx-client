package net

import (
	"encoding/xml"

	"github.com/josh5276/brocade-adx-client/brocade"
)

type NET struct {
	xmlName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    Body
	adc     adc.ADCSoapClient
}
type Body struct {
	xmlName         struct{} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	InterfaceConfig *GetInterfaceConfigResponse
	Msg             *Fault
}

type Fault struct {
	XMLName     xml.Name `xml:"Fault"`
	FaultCode   string   `xml:"faultcode"`
	FaultString string   `xml:"faultstring"`
	Detail      struct {
		Tns string `xml:"tns,attr"`
	}
	FaultId string `xml:"detail>RuntimeFault>faultId"`
	S       string `xml:"s,attr"`
}

///WS/NET Related structures.
type GetInterfaceConfig struct {
	XMLName         xml.Name        `xml:"soap:Envelope"`
	InterfaceConfig interfaceConfig `xml:"soap:Body>tns:getInterfaceConfig"`
	Soap            string          `xml:"xmlns:soap,attr"`
}

type interfaceConfig struct {
	Tns string      `xml:"xmlns:tns,attr"`
	Id  interfaceID `xml:"idList>InterfaceIDSequence"`
}

type interfaceID struct {
	PortType      string `xml:"portType"`
	PortString    string `xml:"portString"`
	InterfaceType string `xml:"interfaceType"`
	Slot          string `xml:"slot"`
	Port          string `xml:"port"`
}

type GetInterfaceConfigResponse struct {
	XMLName      xml.Name `xml:"getInterfaceConfigResponse"`
	Tns          string   `xml:"tns,attr"`
	IntConfigSeq struct {
		PortString       string `xml:"id>portString"`
		Port             string `xml:"id>port"`
		AdminState       string `xml:"adminState"`
		EnableIpv6       string `xml:"enableIpv6"`
		Index            string `xml:"index,attr"`
		Ipv6Mtu          string `xml:"ipv6Mtu"`
		InterfaceType    string `xml:"id>interfaceType"`
		Mtu              string `xml:"mtu"`
		MacAddress       string `xml:"macAddress"`
		EnableSynProxy   string `xml:"enableSynProxy"`
		EnableSynDefense string `xml:"enableSynDefense"`
	} `xml:"configList>InterfaceConfigSequence"`
}

type IntConfigSeqIndex struct {
	Index string `xml:"index,attr"`
}

type XmlGet struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Soap    string   `xml:"xmlns:soap,attr"`
	Call    GetCall  `xml:"soap:Body>xmlns:tns"`
}

type GetCall struct {
	XMLName xml.Name
	Tns     string `xml:"xmlns:tns,attr"`
}
