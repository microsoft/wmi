// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

import (
	_ "github.com/microsoft/wmi/pkg/base/session"
	"testing"
)

func TestGetVirtualNetworkAdapters(t *testing.T) {
	vm, err := GetVirtualMachine(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vm.Close()

	setting, err := vm.GetVirtualMachineSetting()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer setting.Close()

	nas, err := setting.GetVirtualNetworkAdapters()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer nas.Close()
	t.Logf("Virtual Network Adapters Found [%d]", len(nas))
}

func TestGetVirtualHardDisks(t *testing.T) {
	vm, err := GetVirtualMachine(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vm.Close()

	setting, err := vm.GetVirtualMachineSetting()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer setting.Close()

	vhds, err := setting.GetVirtualNetworkAdapters()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vhds.Close()
	t.Logf("Virtual Hard Disks Found [%d]", len(vhds))
}
