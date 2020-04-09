package main

import (
	"fmt"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/base/session"
	"github.com/microsoft/wmi/server2019/root/cimv2"
	"testing"
)

func TestWin32_Process(t *testing.T) {
	defer session.StopWMI()

	p, err := cimv2.NewWin32_ProcessEx6("localhost", "root/cimv2", "", "", "", query.NewWmiQuery("Win32_Process", "Name", "powershell.exe"))
	if err != nil {
		t.Fatal(err)
	}
	retValue, err := p.GetPropertyCommandLine()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("CmdLine: [%s]", retValue)
}
