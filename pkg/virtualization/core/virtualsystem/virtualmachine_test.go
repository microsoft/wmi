// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

import (
	"github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"testing"

	"github.com/microsoft/wmi/pkg/virtualization/core/service"
	"github.com/microsoft/wmi/pkg/virtualization/network/virtualswitch"
)

var (
	whost *host.WmiHost
)

func init() {
	whost = host.NewWmiLocalHost()
}

func TestVirtualMachineCreateGen1(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	setting, err := virtualsystem.GetVirtualSystemSettingData(whost, "test")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer setting.Close()
	t.Logf("Create VMSettings")

	err = settings.SetPropertyVirtualSystemSubType(VirtualSystemSettingData_VirtualSystemSubType_Microsoft_Hyper_V_SubType_1)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	err = settings.SetPropertySecureBootEnabled(false)
	if err != nil {
		return
	}

	memorySettings, err := memory.GetDefaultMemorySettingData(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	memorySettings.SetSizeMB(2048)

	processorSettings, err := processor.GetDefaultProcessorSettingData(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	processorSettings.SetCPUCount(2)

	vm, err := vmms.CreateVirtualMachine(setting, memorySettings, processorSettings)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Created VM [%s]", "test")
	defer func() {
		if vm != nil {
			vm.Close()
		}
	}()
	return

}

func TestAddRemoveIsoDisk(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	vm, err := vmms.GetVirtualMachineByName("test")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer vm.Close()
	t.Logf("Found [%s] VMs", "test")

	// make sure there is a controller
	err = vmms.AddSCSIController(vm)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	path := "c:\\test\\tmp.iso"

	// create an iso file
	err = generateIso(path)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	isodisk, dvddrive, err := vmms.AddISODisk(vm, path)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer isodisk.Close()
	defer dvddrive.Close()

	controllerlocation, err := dvddrive.GetPropertyAddressOnParent()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	controllerNumber, err := dvddrive.GetControllerNumber()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("ControllerNumber [%s], ControllerLocation [%s]", controllerNumber, controllerlocation)

	// remove the iso disk
	err = vmms.RemoveISODisk(isodisk)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Removed ISO disk [%s] from [%s]", path, "test")

	err = os.Remove(path)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
}

func TestAddRemoveTPM(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	vm, err := vmms.GetVirtualMachineByName("test")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer vm.Close()
	t.Logf("Found [%s] VMs", "test")

	tpm, err := vmms.AddTPM(vm)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	fmt.Scanln()
	t.Logf("Added TPM to [%s] VMs", "test")

	err = vmms.RemoveTPM(tpm)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Removed TPM from [%s] VMs", "test")
	fmt.Scanln()
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
