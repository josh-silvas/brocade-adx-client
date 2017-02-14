package sys

import (
	"encoding/xml"
	"github.com/josh5276/brocade-adx-client/brocade"
)


type Sys struct {
	xmlName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    SoapBody
	adx     brocade.ADXSoapClient
}
type SoapBody struct {
	xmlName        struct{} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Version        *GetVersion
	Chassis        *GetChassis
	ConfigSync     *GetConfigSync
	CPU            *GetCPU
	RunningConfig  *GetRunningConfig
	Flash          *GetFlash
	VirtualSummary *GetVirtualSummary
	Cli            *RunCLI
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

// /WS/SYS Related structures.
type GetVersion struct {
	XMLName xml.Name `xml:"getVersionResponse"`
	Tns     string   `xml:"tns,attr"`
	Version string   `xml:"version"`
}

type GetChassis struct {
	XMLName         xml.Name `xml:"getChassisResponse"`
	PartNum         []string `xml:"chassisData>power>powerSupplySequence>partNum"`
	FirmwareVersion []string `xml:"chassisData>power>powerSupplySequence>firmwareVersion"`
	FanRpm          []string `xml:"chassisData>fans>fanSequence>fanRpm"`
	TempSequence    struct {
		Index string `xml:"index,attr"`
	}
	UnitNum             []string `xml:"chassisData>power>powerSupplySequence>unitNum"`
	SerialNum           string   `xml:"chassisData>identification>serialNum"`
	FanSpeed            []string `xml:"chassisData>fans>fanSequence>fanSpeed"`
	Model               string   `xml:"chassisData>identification>model"`
	Wattage             string   `xml:"chassisData>wattage"`
	FanId               []string `xml:"chassisData>fans>fanSequence>fanId"`
	WarnDegC            string   `xml:"chassisData>temperature>tempSequence>warnDegC"`
	ShutDegC            string   `xml:"chassisData>temperature>tempSequence>shutDegC"`
	Power               []string `xml:"chassisData>power>powerSupplySequence>power"`
	BootPromMac         string   `xml:"chassisData>identification>bootPromMac"`
	PowerSupplySequence []struct {
		Index string `xml:"index,attr"`
	}
	Status      []string `xml:"chassisData>power>powerSupplySequence>status"`
	PSSerialNum []string `xml:"chassisData>power>powerSupplySequence>serialNum"`
	FanSequence []struct {
		Index string `xml:"index,attr"`
	}
	FanStatus   []string `xml:"chassisData>fans>fanSequence>fanStatus"`
	Module      string   `xml:"chassisData>temperature>tempSequence>module"`
	Tns         string   `xml:"tns,attr"`
	CurrentDegC string   `xml:"chassisData>temperature>tempSequence>currentDegC"`
}

type GetConfigSync struct {
	XMLName                   xml.Name `xml:"getConfigSyncConfigResponse"`
	PortString                string   `xml:"config>config>interfaceID>portString"`
	Slot                      string   `xml:"config>config>interfaceID>slot"`
	TimeSinceLastConfigSync   string   `xml:"config>syncStatus>timeElapsedSinceLastConfigSync"`
	Tns                       string   `xml:"tns,attr"`
	Port                      string   `xml:"config>config>interfaceID>port"`
	VlanID                    string   `xml:"config>config>vlanID"`
	SyncStatus                string   `xml:"config>syncStatus>syncStatus"`
	LastConfigSyncOpStatus    string   `xml:"config>syncStatus>lastConfigSyncOperationStatus"`
	TimeSinceLastConfigChange string   `xml:"config>syncStatus>timeElapsedSinceLastConfigChange"`
	Mode                      string   `xml:"config>mode"`
	PeerMacAddress            string   `xml:"config>config>peerMacAddress"`
	InterfaceType             string   `xml:"config>config>interfaceID>interfaceType"`
}

type GetCPU struct {
	XMLName xml.Name `xml:"getCPUResponse"`
	CPU     []struct {
		Tns                       string `xml:"tns,attr"`
		Index                     string `xml:"index,attr"`
		SecondsSince              string `xml:"last1MinAvg>secondsSince"`
		ModName                   string `xml:"modName"`
		SecondsSinceAverageCpu    string `xml:"average>secondsSince"`
		PercentLoadLastPeakCpu    string `xml:"lastPeak>percentLoad"`
		PeakCpu                   string `xml:"lastPeak>secondsSince"`
		PercentLoadLast5SecAvgCpu string `xml:"last5SecAvg>percentLoad"`
		PercentLoad               string `xml:"last1MinAvg>percentLoad"`
		FiveMinAvgCpu             string `xml:"last5MinAvg>secondsSince"`
		PercentLoadAverageCpu     string `xml:"average>percentLoad"`
		FiveSecAvgCpu             string `xml:"last5SecAvg>secondsSince"`
		PercentLoadLast5MinAvgCpu string `xml:"last5MinAvg>percentLoad"`
		PercentLoadLast1SecAvgCpu string `xml:"last1SecAvg>percentLoad"`
		OneSecAvgCpu              string `xml:"last1SecAvg>secondsSince"`
	} `xml:"modCpuUtil>cpuUtilSequence"`
}

type GetFlash struct {
	XMLName          xml.Name `xml:"getFlashResponse"`
	PrimaryVersion   string   `xml:"flashData>primaryImage>version"`
	Tns              string   `xml:"tns,attr"`
	SecondaryLabel   string   `xml:"flashData>secondaryImage>label"`
	FreeKBytes       string   `xml:"flashData>codeFlash>freeKBytes"`
	FreeBytes        string   `xml:"flashData>configFlash>freeBytes"`
	TotalKBytes      string   `xml:"flashData>codeFlash>totalKBytes"`
	PrimarySize      string   `xml:"flashData>primaryImage>size"`
	SecondarySize    string   `xml:"flashData>secondaryImage>size"`
	UsedBytes        string   `xml:"flashData>configFlash>usedBytes"`
	PrimaryBuiltOn   string   `xml:"flashData>primaryImage>builtOn"`
	NameCodeFlash    string   `xml:"flashData>codeFlash>name"`
	UsedKBytes       string   `xml:"flashData>codeFlash>usedKBytes"`
	Name             string   `xml:"flashData>configFlash>name"`
	TotalBytes       string   `xml:"flashData>configFlash>totalBytes"`
	ModuleName       string   `xml:"flashData>moduleName"`
	PrimaryLabel     string   `xml:"flashData>primaryImage>label"`
	SecondaryVersion string   `xml:"flashData>secondaryImage>version"`
	SecondaryBuiltOn string   `xml:"flashData>secondaryImage>builtOn"`
}

type GetRunningConfig struct {
	XMLName   xml.Name `xml:"getRunningConfigResponse"`
	Tns       string   `xml:"tns,attr"`
	RunConfig string   `xml:"runConfig"`
}

type RunCLI struct {
	XMLName        xml.Name `xml:"runCLIResponse"`
	ResultSequence []struct {
		Index   string `xml:"index,attr"`
		Command string `xml:"command"`
		Output  string `xml:"output"`
	} `xml:"result>ResultSequence"`
	Tns string `xml:"tns,attr"`
}

// /WS/SYS Related structures.
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
