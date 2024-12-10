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
	vm, err := GetVirtualMachineByVMName(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer vm.Close()

	t.Logf("VirtualMachine successfully found")
}

func TestVirtualMachineState(t *testing.T) {
	vm, err := GetVirtualMachineByVMName(whost, "test")
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
	vm, err := GetVirtualMachineByVMName(whost, "test")
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
	vm, err := GetVirtualMachineByVMName(whost, "test")
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
	vm, err := GetVirtualMachineByVMName(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vm.Close()

	err = vm.Stop(false)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
}

func TestVirtualMachinePause(t *testing.T) {
	vm, err := GetVirtualMachineByVMName(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vm.Close()

	err = vm.Pause()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
}

func TestVirtualMachineSave(t *testing.T) {
	vm, err := GetVirtualMachineByVMName(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vm.Close()

	err = vm.Save()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
}

func TestVirtualMachineResume(t *testing.T) {
	vm, err := GetVirtualMachineByVMName(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vm.Close()

	err = vm.Pause()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	err = vm.Resume()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
}

func TestVirtualMachineRestore(t *testing.T) {
	vm, err := GetVirtualMachineByVMName(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vm.Close()

	err = vm.Save()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	err = vm.Restore()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
}

func TestGetVirtualMachineSetting(t *testing.T) {
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
}

func TestGetVirtualMachineOSConfiguration(t *testing.T) {
	vm, err := GetVirtualMachineByVMName(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vm.Close()

	cName, isWindows, err := vm.GetOSConfiguration()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	t.Logf("Retrieved OS configuration. Computer Name %s IsWindows %t", cName, isWindows)
}
