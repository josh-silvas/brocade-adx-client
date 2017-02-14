package adc

import (
	"testing"
	"github.com/josh5276/brocade-adx-client/brocade/sys"
)


var (
	// user/password need to be valid ACS credentials to
	// authenticate the ADX SOAP Client
	user   = ""
	passwd = ""
	adxip  = ""
)

func TestADXSoapClient_TestAuth(t *testing.T) {
	if user == "" || passwd == "" {
		t.Fatal("Username a password must be provided to run tests")
	}
	t.Logf("Testing connectivity with the ServerIron SOAP client using  %v", user)
	resp, err := TestAuth(NewSOAPClient(adxip, user, passwd))
	if err != nil {
		t.Error(err)
	}
	t.Logf("Successful response from the ADX. ServerIron ADX Version: %v", resp)
}

func TestADXSoapClient_Sys(t *testing.T) {
	if user == "" || passwd == "" {
		t.Fatal("Username a password must be provided to run tests")
	}
	t.Logf("Testing connectivity with the ServerIron SOAP client using  %v", user)
	s := sys.New(NewSOAPClient(adxip, user, passwd))
	resp, code, err := s.Sys("getChassis")
	if err != nil {
		t.Error(err, code)
	}
	t.Logf("Successful response from the ADX. ServerIron ADX Model: %v", resp.Body.Chassis.Model)
}

func BenchmarkADXSoapClient_Sys(b *testing.B) {
	if user == "" || passwd == "" {
		b.Fatal("Username a password must be provided to run benchmark")
	}
	b.Logf("Benchmark testing the ServerIron SOAP client using  %v", user)
	s := sys.New(NewSOAPClient(adxip, user, passwd))
	for i := 0; i < b.N; i++ {
		resp, code, err := s.Sys("getChassis")
		if err != nil {
			b.Fatal(err, code)
		}
		b.Logf("Success. ServerIron ADX Model: %v", resp.Body.Chassis.Model)
	}
	b.Log("Benchmark completed for the ADX SOAP Client")
}

func BenchmarkADXSoapClient_Sys_RunningConfig(b *testing.B) {
	if user == "" || passwd == "" {
		b.Fatal("Username a password must be provided to run benchmark")
	}
	b.Logf("Benchmark testing the ServerIron SOAP client using  %v", user)
	s := sys.New(NewSOAPClient(adxip, user, passwd))
	for i := 0; i < b.N; i++ {
		resp, code, err := s.Sys("getRunningConfig")
		if err != nil {
			b.Fatal(err, code)
		}
		b.Logf("Success. ServerIron ADX Running config tests: %v", resp.Body.RunningConfig.Tns)
	}
	b.Log("Benchmark completed for the ADX SOAP Client")
}
