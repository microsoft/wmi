// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystemmanagement

import (
	"github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"testing"

	"github.com/microsoft/wmi/pkg/virtualization/network/virtualswitch"
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
	t.Logf("Found [%d] VMs", len(vms))
	defer vms.Close()
	vm, err := vmms.GetVirtualMachineByName("test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	t.Logf("Found [%s] VMs", "test")
	defer vm.Close()
}
func TestVirtualMachineAdapterScenario(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	vm, err := vmms.GetVirtualMachineByName("test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	t.Logf("Found [%s] VMs", "test")
	defer vm.Close()

	err = vmms.RenameVirtualNetworkAdapter(vm, "Network Adapter", "testadapter")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	t.Logf("Renamed Adapter [%s]", "testadapter")

	testna, err := vm.GetVirtualNetworkAdapterByName("testadapter")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer testna.Close()
	t.Logf("Found Renamed Adapter [%s]", "testadapter")

	err = vmms.SetVirtualNetworkAdapterMACAddress(vm, "testadapter", "001122334455")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	t.Logf("Set Adapter Mac [%s]", "001122334455")

	vs, err := virtualswitch.GetVirtualSwitch(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vs.Close()
	t.Logf("Got VirtualSwitch[%s]", "test")

	err = vmms.ConnectAdapterToVirtualSwitch(vm, "testadapter", vs)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	t.Logf("Connect VM[%s] to VirtualSwitch[%s]", "test", "test")
	err = vmms.DisconnectAdapterFromVirtualSwitch(vm, "testadapter")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	t.Logf("Disconnect VM[%s] from VirtualSwitch[%s]", "test", "test")
}

func TestVirtualMachineDelete(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	vm, err := vmms.GetVirtualMachineByName("test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	defer vm.Close()
	err = vmms.DeleteVirtualMachine(vm)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
}
