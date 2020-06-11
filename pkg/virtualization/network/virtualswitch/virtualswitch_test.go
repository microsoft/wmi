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

func TestGetSwitchExtensions(t *testing.T) {
	vs, err := GetVirtualSwitch(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer vs.Close()

	vnas, err := vs.GetEthernetSwitchExtensions()
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}

	defer vnas.Close()
}

func TestCheckSwitchExtensions(t *testing.T) {
	vs, err := GetVirtualSwitch(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer vs.Close()

	t.Logf("Get Switch Extension by Name")
	se, err := vs.GetEthernetSwitchExtensionByName("Microsoft Azure VFP Switch Extension")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer se.Close()
	t.Logf("Check Enabled state ")
	enabled, err := se.IsEnabled()
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	t.Logf("Enabled state [%t]\n", enabled)
	return
}
