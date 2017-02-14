package slb

import (
	"encoding/xml"
	"github.com/josh5276/brocade-adx-client/brocade"
)


type SLB struct {
	xmlName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    Body
	adc     adc.ADCSoapClient
}
type Body struct {
	xmlName        struct{} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	VirtualSummary *GetVirtualSummary
	Msg            *Fault
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

// /WS/SLB Related structures.
type GetVirtualSummary struct {
	XMLName xml.Name `xml:"getAllVirtualServerSummaryResponse"`
	Virtual []struct {
		AdminState                string `xml:"ports>VirtualServerPortSummarySequence>adminState"`
		CurrentConnection         string `xml:"ports>VirtualServerPortSummarySequence>currentConnection"`
		PortTxPktsRate            string `xml:"ports>VirtualServerPortSummarySequence>txPktsRate"`
		Name                      string `xml:"server>Name"`
		IP                        string `xml:"server>IP"`
		RxPktsRate                string `xml:"rxPktsRate"`
		TotalConn                 string `xml:"totalConn"`
		CustomHealthCheck         string `xml:"ports>VirtualServerPortSummarySequence>customHealthCheck"`
		VirtualServerPortNameSrvr string `xml:"ports>VirtualServerPortSummarySequence>virtualServerPort>srvr>Name"`
		Predictor                 string `xml:"Predictor"`
		CurrentConnRate           string `xml:"currentConnRate"`
		TotalPorts                string `xml:"totalPorts"`
		VirtualTxPktsRate         string `xml:"txPktsRate"`
		PortSequence              struct {
			Index string `xml:"index,attr"`
		}
		RunTimeStatus           string `xml:"ports>VirtualServerPortSummarySequence>runTimeStatus"`
		RxBytes                 string `xml:"rxBytes"`
		SymmetricPriority       string `xml:"symmetricPriority"`
		EnableSticky            string `xml:"ports>VirtualServerPortSummarySequence>enableSticky"`
		RcvPkts                 string `xml:"rcvPkts"`
		CurrentConnections      string `xml:"currentConnections"`
		EnableDsr               string `xml:"ports>VirtualServerPortSummarySequence>enableDsr"`
		IsConcurrent            string `xml:"ports>VirtualServerPortSummarySequence>isConcurrent"`
		TcpOnly                 string `xml:"ports>VirtualServerPortSummarySequence>tcpOnly"`
		VirtualServerPortIPSrvr string `xml:"ports>VirtualServerPortSummarySequence>virtualServerPort>srvr>IP"`
		NameOrNumber            string `xml:"ports>VirtualServerPortSummarySequence>virtualServerPort>port>NameOrNumber"`
		TxPkts                  string `xml:"txPkts"`
		TxBytes                 string `xml:"txBytes"`
		Stateless               string `xml:"ports>VirtualServerPortSummarySequence>Stateless"`
		UdpOnly                 string `xml:"ports>VirtualServerPortSummarySequence>udpOnly"`
		VirtualSequence         struct {
			Index string `xml:"index,attr"`
		}
		Tns               string   `xml:"tns,attr"`
		RunTimeState      string   `xml:"runTimeState"`
		PortRxPktsRate    []string `xml:"ports>VirtualServerPortSummarySequence>rxPktsRate"`
		VirtualAdminState string   `xml:"adminState"`
	} `xml:"virtualServerSummary>VirtualServerSummarySequence"`
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

type RunCliRequest struct {
	XMLName xml.Name       `xml:"soap:Envelope"`
	RunCLI  RunCliCommands `xml:"soap:Body>tns:runCLI"`
	Soap    string         `xml:"xmlns:soap,attr"`
}

type RunCliCommands struct {
	Tns            string   `xml:"xmlns:tns,attr"`
	StringSequence []string `xml:"cmds>StringSequence"`
}
