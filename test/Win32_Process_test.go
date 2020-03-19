package main

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/base/session"
	"github.com/microsoft/wmi/server2019/root/cimv2"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	session.StartWMI()
	defer session.StopWMI()
	p, err := cimv2.NewWin32_Process("localhost", "root/cimv2", "", "", "", query.NewWmiQuery("Win32_Process", "Name", "System"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p)
	// t.Fatal("not implemented")
}
