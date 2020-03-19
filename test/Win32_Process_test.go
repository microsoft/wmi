package main

import (
	"fmt"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/base/session"
	"github.com/microsoft/wmi/server2019/root/cimv2"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	session.StartWMI()
	defer session.StopWMI()

	p, err := cimv2.NewWin32_ProcessEx6("localhost", "ROOT\\CimV2", "", "", "", query.NewWmiQuery("Win32_Process", "Name", "System"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = p.GetPropertyCommandLine()
	if err != nil {
		t.Fatal(err)
	}
	retValue, err := p.GetProperty("CommandLine")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("CmdLine:" + retValue.(string))
	// t.Fatal("not implemented")
}
