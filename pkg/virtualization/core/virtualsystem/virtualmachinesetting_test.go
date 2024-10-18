// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

import (
	"testing"

	_ "github.com/microsoft/wmi/pkg/base/session"
)

func TestGetVirtualNetworkAdapters(t *testing.T) {
	vm, err := GetVirtualMachineByVMName(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vm.Close()

	setting, err := vm.GetVirtualSystemSettingData()
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
	vm, err := GetVirtualMachineByVMName(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vm.Close()

	setting, err := vm.GetVirtualSystemSettingData()
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

func TestGetVirtualDvdDrives(t *testing.T) {
	vm, err := GetVirtualMachineByVMName(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vm.Close()

	setting, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer setting.Close()

	dvds, err := setting.GetVirtualDvdDrives()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer dvds.Close()
	t.Logf("DVD drives Found [%d]", len(dvds))
}
