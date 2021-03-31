// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"fmt"
	"os"
	"testing"

	"github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"github.com/microsoft/wmi/pkg/virtualization/core/memory"
	"github.com/microsoft/wmi/pkg/virtualization/core/processor"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/disk"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/service"
	"github.com/microsoft/wmi/pkg/virtualization/core/virtualsystem"
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
		t.Fatalf("Failed [%+v]", err)
	}
}

func TestCreateVirtualMachines(t *testing.T) {
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

	memorySettings, err := memory.GetDefaultMemorySettingData(whost)
	if err != nil {
		return
	}
	memorySettings.SetSizeMB(2048)

	processorSettings, err := processor.GetDefaultProcessorSettingData(whost)
	if err != nil {
		return
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

func TestModifyVirtualMachine(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	vm, err := vmms.GetVirtualMachineByName("test")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Found [%s] VMs", "test")
	defer vm.Close()

	t.Logf("Setting Memory [%d] [%s] VMs", 2048, "test")
	err = vmms.SetMemoryMB(vm, 2048)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Setting Processor [%d] [%s] VMs", 4, "test")
	err = vmms.SetProcessorCount(vm, 4)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
}

func TestGetVirtualMachines(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	vms, err := vmms.GetVirtualMachines()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Found [%d] VMs", len(vms))
	defer vms.Close()
	vm, err := vmms.GetVirtualMachineByName("test")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Found [%s] VMs", "test")
	defer vm.Close()
}

func TestVirtualMachineAdapterScenario(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	vs, err := virtualswitch.GetVirtualSwitch(whost, "test")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer vs.Close()
	t.Logf("Got VirtualSwitch[%s]", "test")

	vm, err := vmms.GetVirtualMachineByName("test")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Found [%s] VMs", "test")
	defer vm.Close()

	for i := 1; i <= 4; i++ {
		adapterName := fmt.Sprintf("testadapter-%d", i)
		t.Logf("Adding VmNic to [%s]", "test")
		na, err := vmms.AddVirtualNetworkAdapter(vm, "Network Adapter")
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		defer na.Close()

		err = vmms.RenameVirtualNetworkAdapter(vm, "Network Adapter", adapterName)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		t.Logf("Renamed Adapter [%s]", adapterName)

		testna, err := vm.GetVirtualNetworkAdapterByName(adapterName)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		defer testna.Close()
		t.Logf("Found Renamed Adapter [%s]", adapterName)

		macAddress := fmt.Sprintf("00112233445%d", i)
		err = vmms.SetVirtualNetworkAdapterMACAddress(vm, adapterName, macAddress)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		t.Logf("Set Adapter Mac [%s]", macAddress)

		err = vmms.ConnectAdapterToVirtualSwitch(vm, adapterName, vs)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		t.Logf("Connect VM[%s] to VirtualSwitch[%s]", "test", "test")
		err = vmms.SetVirtualNetworkAdapterAccessVLAN(testna, uint16(i))
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		t.Logf("Set Adapter VLAN [%d]", i)
		err = vmms.SetVirtualNetworkAdapterPortProfile(testna, "test",
			"1FA41B39-B444-4E43-B35A-E1F7985FD548",
			"08cf6a3b-f4ea-48d9-a29c-60370363fb19", 1)
		if err != nil {
			t.Fatalf("Set PortProfile Failed [%+v]", err)
		}
		t.Logf("Set Adapter Port Profile [%d]", i)
	}
	for i := 1; i <= 4; i++ {
		adapterName := fmt.Sprintf("testadapter-%d", i)
		err = vmms.DisconnectAdapterFromVirtualSwitch(vm, adapterName)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		t.Logf("Disconnect VM[%s] from VirtualSwitch[%s]", "test", "test")
		testna, err := vm.GetVirtualNetworkAdapterByName(adapterName)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		defer testna.Close()
		t.Logf("Removing VmNic [%s] from VM[%s] ", adapterName, "test")
		err = vmms.RemoveVirtualNetworkAdapter(testna)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)

		}
	}
}

func TestVirtualMachineAdapterCreationWithMac(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	vs, err := virtualswitch.GetVirtualSwitch(whost, "test")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer vs.Close()
	t.Logf("Got VirtualSwitch[%s]", "test")

	vm, err := vmms.GetVirtualMachineByName("test")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Found [%s] VMs", "test")
	defer vm.Close()

	adapterName := "testadapter"
	macAddress := "001122334450"
	t.Logf("Adding VmNic with MAC to [%s]", "test")
	na, err := vmms.AddVirtualNetworkAdapterWithMac(vm, adapterName, macAddress)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer na.Close()

	testna, err := vm.GetVirtualNetworkAdapterByName(adapterName)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer testna.Close()
	t.Logf("Found Adapter [%s]", adapterName)

	t.Logf("Removing VmNic [%s] from VM[%s] ", adapterName, "test")
	err = vmms.RemoveVirtualNetworkAdapter(testna)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
}

func TestAddRemoveVirtualHardDisk(t *testing.T) {
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

	err = vmms.AddSCSIController(vm)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	ims, err := service.GetImageManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Got ImageManagementService ")

	for i := 1; i <= 4; i++ {
		path := fmt.Sprintf("c:\\test\\tmp-%d.vhdx", i)
		setting, err := disk.GetVirtualHardDiskSettingData(whost, path, 512, 512, 0, 1024*1024*10, true)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		defer setting.Close()
		err = ims.CreateDisk(setting)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		defer os.RemoveAll(path)
		t.Logf("Created vhd [%s]", path)

		vhd, vhddrive, err := vmms.AttachVirtualHardDisk(vm, path)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		defer vhd.Close()
		defer vhddrive.Close()
		t.Logf("Attached vhd [%s] to [%s]", path, "test")
		controllerlocation, err := vhddrive.GetControllerLocation()
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		controllerNumber, err := vhddrive.GetControllerNumber()
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		t.Logf("ControllerNumber [%s], ControllerLocation [%s]", controllerNumber, controllerlocation)
	}

	for i := 1; i <= 4; i++ {
		path := fmt.Sprintf("c:\\test\\tmp-%d.vhdx", i)
		vhd, err := vm.GetVirtualHardDiskByLocation(0, i-1)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		defer vhd.Close()
		err = vmms.DetachVirtualHardDisk(vhd)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		t.Logf("Detached vhd [%s] from [%s]", path, "test")
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

func TestVirtualMachineDelete(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	vm, err := vmms.GetVirtualMachineByName("test")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	defer vm.Close()
	err = vm.Start()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	err = vm.Stop(true)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	err = vmms.DeleteVirtualMachine(vm)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
}
