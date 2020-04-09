// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystemmanagement

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

func TestGetVirtualMachineManagementService(t *testing.T) {
	_, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
}

func TestGetVirtualMachines(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	vms, err := vmms.GetVirtualMachines()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	vms.Close()
}

func TestVirtualMachineDelete(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	vm, err := vmms.FindVirtualMachineByName("test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	if vm == nil {
		t.Fatal("Failed to get VM")
	}
	defer vm.Close()
	err = vmms.DeleteVirtualMachine(vm)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
}
