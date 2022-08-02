// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

import (
	"testing"

	"github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
)

var (
	whost *host.WmiHost
)

func init() {
	whost = host.NewWmiLocalHost()
}

func TestVirtualMachineCreate(t *testing.T) {
	//TODO
}

func TestGetVirtualMachine(t *testing.T) {
	vm, err := GetVirtualMachine(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer vm.Close()

	t.Logf("VirtualMachine successfully found")
}

func TestVirtualMachineState(t *testing.T) {
	vm, err := GetVirtualMachine(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vm.Close()

	t.Logf("VirtualMachine successfully found")
	state, err := vm.State()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	t.Logf("VirtualMachine State [%d]", state)

}

func TestVirtualMachineStart(t *testing.T) {
	vm, err := GetVirtualMachine(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vm.Close()

	err = vm.Start()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
}

func TestVirtualMachineStop(t *testing.T) {
	vm, err := GetVirtualMachine(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vm.Close()

	err = vm.Stop(true)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
}

func TestVirtualMachineShutdown(t *testing.T) {
	vm, err := GetVirtualMachine(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vm.Close()

	err = vm.Stop(false)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
}

func TestGetVirtualMachineSetting(t *testing.T) {
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
}
