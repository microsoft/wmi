// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualswitch

import (
	"github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"testing"
)

var (
	whost *host.WmiHost
)

func init() {
	whost = host.NewWmiLocalHost()
}

func TestGetVirtualSwitch(t *testing.T) {
	vs, err := GetVirtualSwitch(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer vs.Close()
}

func TestGetVirtualNetworkAdapters(t *testing.T) {
	vs, err := GetVirtualSwitch(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer vs.Close()

	vnas, err := vs.GetVirtualMachineAdapters()
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}

	defer vnas.Close()
}

func TestConnectDisconnectVirtualNetworkAdapters(t *testing.T) {
	vs, err := GetVirtualSwitch(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer vs.Close()

	vna, err := vs.GetVirtualMachineAdapterByName("Network Adapter")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}

	defer vna.Close()

	err = vs.ConnectVirtualNetworkAdapter(vna)
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}

	err = vs.DisconnectVirtualNetworkAdapter(vna)
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
}
