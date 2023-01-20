// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"github.com/microsoft/wmi/pkg/virtualization/core/memory"
	"github.com/microsoft/wmi/pkg/virtualization/core/processor"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/disk"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/service"
	"github.com/microsoft/wmi/pkg/virtualization/core/virtualsystem"
	vswitchservice "github.com/microsoft/wmi/pkg/virtualization/network/service"
	"github.com/microsoft/wmi/pkg/virtualization/network/virtualswitch"
	//v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
	"github.com/nwoodmsft/iso9660"
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

	err = setting.SetHyperVGeneration(virtualsystem.HyperVGeneration_V2)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
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

func TestCreateVirtualMachinesGen1(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	setting, err := virtualsystem.GetVirtualSystemSettingData(whost, "testGen1")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer setting.Close()
	t.Logf("Create VMSettings")

	err = setting.SetHyperVGeneration(virtualsystem.HyperVGeneration_V1)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	err = setting.SetPropertySecureBootEnabled(false)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
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
	t.Logf("Created VM [%s]", "testGen1")
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

func TestAddRemoveIsoDiskGen1(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	vm, err := vmms.GetVirtualMachineByName("testGen1")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer vm.Close()
	t.Logf("Found [%s] VMs", "testGen1")

	// make sure there is a controller
	err = vmms.AddSCSIController(vm)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	path := "c:\\test\\tmpgen1.iso"

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

func generateIso(path string) error {
	writer, err := iso9660.NewWriter()
	if err != nil {
		return err
	}
	defer writer.Cleanup()

	isofile, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	isofile.Close()
	return nil
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

func TestModifyVirtualMachineGen1(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	vm, err := vmms.GetVirtualMachineByName("testGen1")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Found [%s] VMs", "testGen1")
	defer vm.Close()

	t.Logf("Setting Memory [%d] [%s] VMs", 2048, "testGen1")
	err = vmms.SetMemoryMB(vm, 2048)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Setting Processor [%d] [%s] VMs", 4, "testGen1")
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

func TestGetVirtualMachinesGen1(t *testing.T) {
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
	vm, err := vmms.GetVirtualMachineByName("testGen1")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Found [%s] VMs", "testGen1")
	defer vm.Close()
}

func TestVirtualMachineAdapterScenario(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	vsms, err := vswitchservice.GetVirtualEthernetSwitchManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	setting, err := virtualswitch.GetVirtualEthernetSwitchSettingData(whost, "test")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer setting.Close()
	vswitch, err := vsms.CreatePrivateVirtualSwitch(setting)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer vswitch.Close()

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

	err = vsms.DeleteVirtualSwitch(vswitch)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

}

func TestVirtualMachineAdapterScenarioGen1(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	vsms, err := vswitchservice.GetVirtualEthernetSwitchManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	setting, err := virtualswitch.GetVirtualEthernetSwitchSettingData(whost, "testGen1")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer setting.Close()
	vswitch, err := vsms.CreatePrivateVirtualSwitch(setting)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer vswitch.Close()

	vs, err := virtualswitch.GetVirtualSwitch(whost, "testGen1")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer vs.Close()
	t.Logf("Got VirtualSwitch[%s]", "testGen1")

	vm, err := vmms.GetVirtualMachineByName("testGen1")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Found [%s] VMs", "testGen1")
	defer vm.Close()

	for i := 1; i <= 4; i++ {
		adapterName := fmt.Sprintf("testadapter-%d", i)
		t.Logf("Adding VmNic to [%s]", "testGen1")
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
		t.Logf("Connect VM[%s] to VirtualSwitch[%s]", "testGen1", "testGen1")
		err = vmms.SetVirtualNetworkAdapterAccessVLAN(testna, uint16(i))
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		t.Logf("Set Adapter VLAN [%d]", i)
		err = vmms.SetVirtualNetworkAdapterPortProfile(testna, "testGen1",
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
		t.Logf("Removing VmNic [%s] from VM[%s] ", adapterName, "testGen1")
		err = vmms.RemoveVirtualNetworkAdapter(testna)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)

		}
	}

	err = vsms.DeleteVirtualSwitch(vswitch)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
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

func TestVirtualMachineAdapterCreationWithMacGen1(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	vs, err := virtualswitch.GetVirtualSwitch(whost, "testGen1")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer vs.Close()
	t.Logf("Got VirtualSwitch[%s]", "testGen1")

	vm, err := vmms.GetVirtualMachineByName("testGen1")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Found [%s] VMs", "testGen1")
	defer vm.Close()

	adapterName := "testadapter"
	macAddress := "001122334450"
	t.Logf("Adding VmNic with MAC to [%s]", "testGen1")
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

	t.Logf("Removing VmNic [%s] from VM[%s] ", adapterName, "testGen1")
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
		setting, err := disk.GetVirtualHardDiskSettingData(whost, path, 512, 512, 0, 1024*1024*10, true, disk.VirtualHardDiskFormat_2)
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

func TestAddRemoveVirtualHardDiskGen1(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	vm, err := vmms.GetVirtualMachineByName("testGen1")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer vm.Close()
	t.Logf("Found [%s] VMs", "testGen1")

	err = vmms.AddSCSIController(vm)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	ims, err := service.GetImageManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Got ImageManagementService ")

	path := fmt.Sprintf("c:\\test\\tmp-%d.vhd", 1)
	setting, err := disk.GetVirtualHardDiskSettingData(whost, path, 512, 512, 0, 1024*1024*10, true, disk.VirtualHardDiskFormat_1)
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
	t.Logf("Attached vhd [%s] to [%s]", path, "testGen1")
	controllerlocation, err := vhddrive.GetControllerLocation()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	controllerNumber, err := vhddrive.GetControllerNumber()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("ControllerNumber [%s], ControllerLocation [%s]", controllerNumber, controllerlocation)

	path = fmt.Sprintf("c:\\test\\tmp-%d.vhd", 1)
	vhd, err = vm.GetVirtualHardDiskByLocation(0, 0)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer vhd.Close()
	err = vmms.DetachVirtualHardDisk(vhd)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Detached vhd [%s] from [%s]", path, "testGen1")

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

func TestAddRemoveTPMGen1(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	vm, err := vmms.GetVirtualMachineByName("testGen1")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer vm.Close()
	t.Logf("Found [%s] VMs", "testGen1")

	tpm, err := vmms.AddTPM(vm)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	fmt.Scanln()
	t.Logf("Added TPM to [%s] VMs", "testGen1")

	err = vmms.RemoveTPM(tpm)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Removed TPM from [%s] VMs", "testGen1")
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

func TestVirtualMachineDeleteGen1(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	vm, err := vmms.GetVirtualMachineByName("testGen1")
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

func TestCreateDynamicMemoryVirtualMachine(t *testing.T) {
	// create
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	setting, err := virtualsystem.GetVirtualSystemSettingData(whost, "dynamic-memory-test")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer setting.Close()

	err = setting.SetHyperVGeneration(virtualsystem.HyperVGeneration_V2)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	err = setting.SetPropertyVirtualNumaEnabled(false)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	memorySettings, err := memory.GetDefaultMemorySettingData(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	memorySettings.SetSizeMB(2048)

	config := memory.DynamicMemoryConfiguration{
		DynamicMemoryEnabled: true,
		MaximumMemoryMB:      4096,
		MinimumMemoryMB:      1024,
		TargetMemoryBuffer:   20,
	}
	err = memorySettings.ConfigureDynamicMemory(&config)
	if err != nil {
		t.Fatalf("Failed to configure dynamic memory [%v]", err)
	}

	processorSettings, err := processor.GetDefaultProcessorSettingData(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	processorSettings.SetCPUCount(2)

	vm, err := vmms.CreateVirtualMachine(setting, memorySettings, processorSettings)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer func() {
		if vm != nil {
			err = vmms.DeleteVirtualMachine(vm)
			if err != nil {
				t.Logf("VM deletion failed [%+v]", err)
			}
			vm.Close()
		}
	}()

	// validate
	vm, err = vmms.GetVirtualMachineByName("dynamic-memory-test")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	mem, err := vm.GetMemory()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	min, err := mem.GetMinimumMemoryMB()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	max, err := mem.GetMaximumMemoryMB()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	buf, err := mem.GetTargetMemoryBuffer()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	enabled, err := mem.GetPropertyDynamicMemoryEnabled()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	if config.DynamicMemoryEnabled != enabled {
		t.Fatal("Failed DynamicMemoryEnabled doesn't match config")
	} else if config.MaximumMemoryMB != max {
		t.Fatal("Failed MaximumMemoryMB doesn't match config")
	} else if config.MinimumMemoryMB != min {
		t.Fatal("Failed MinimumMemoryMB doesn't match config")
	} else if config.TargetMemoryBuffer != buf {
		t.Fatal("Failed TargetMemoryBuffer doesn't match config")
	}

	err = vm.Start()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	err = vm.Stop(true)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
}

func TestCreateDynamicMemoryVirtualMachineGen1(t *testing.T) {
	// create
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	setting, err := virtualsystem.GetVirtualSystemSettingData(whost, "dynamic-memory-test-gen1")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer setting.Close()

	err = setting.SetHyperVGeneration(virtualsystem.HyperVGeneration_V1)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	err = setting.SetPropertySecureBootEnabled(false)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	err = setting.SetPropertyVirtualNumaEnabled(false)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	memorySettings, err := memory.GetDefaultMemorySettingData(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	memorySettings.SetSizeMB(2048)

	config := memory.DynamicMemoryConfiguration{
		DynamicMemoryEnabled: true,
		MaximumMemoryMB:      4096,
		MinimumMemoryMB:      1024,
		TargetMemoryBuffer:   20,
	}
	err = memorySettings.ConfigureDynamicMemory(&config)
	if err != nil {
		t.Fatalf("Failed to configure dynamic memory [%v]", err)
	}

	processorSettings, err := processor.GetDefaultProcessorSettingData(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	processorSettings.SetCPUCount(2)

	vm, err := vmms.CreateVirtualMachine(setting, memorySettings, processorSettings)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer func() {
		if vm != nil {
			err = vmms.DeleteVirtualMachine(vm)
			if err != nil {
				t.Logf("VM deletion failed [%+v]", err)
			}
			vm.Close()
		}
	}()

	// validate
	vm, err = vmms.GetVirtualMachineByName("dynamic-memory-test-gen1")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	mem, err := vm.GetMemory()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	min, err := mem.GetMinimumMemoryMB()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	max, err := mem.GetMaximumMemoryMB()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	buf, err := mem.GetTargetMemoryBuffer()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	enabled, err := mem.GetPropertyDynamicMemoryEnabled()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	if config.DynamicMemoryEnabled != enabled {
		t.Fatal("Failed DynamicMemoryEnabled doesn't match config")
	} else if config.MaximumMemoryMB != max {
		t.Fatal("Failed MaximumMemoryMB doesn't match config")
	} else if config.MinimumMemoryMB != min {
		t.Fatal("Failed MinimumMemoryMB doesn't match config")
	} else if config.TargetMemoryBuffer != buf {
		t.Fatal("Failed TargetMemoryBuffer doesn't match config")
	}

	err = vm.Start()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	err = vm.Stop(true)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
}

func TestBindCpuGroupVirtualMachine(t *testing.T) {
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

	err = setting.SetHyperVGeneration(virtualsystem.HyperVGeneration_V2)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
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
	g, err := uuid.NewUUID()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	cpuGroupId := g.String()
	err = vmms.SetCPUGroupID(vm, cpuGroupId)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Binded VM [%s] to Cpugroup [%s]", "test", cpuGroupId)
}

func TestBindCpuGroupVirtualMachineGen1(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	setting, err := virtualsystem.GetVirtualSystemSettingData(whost, "testGen1")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer setting.Close()
	t.Logf("Create VMSettings")

	err = setting.SetHyperVGeneration(virtualsystem.HyperVGeneration_V1)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	err = setting.SetPropertySecureBootEnabled(false)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
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
	g, err := uuid.NewUUID()
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	cpuGroupId := g.String()
	err = vmms.SetCPUGroupID(vm, cpuGroupId)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Binded VM [%s] to Cpugroup [%s]", "test", cpuGroupId)
}

func TestAddRemoveGpuDda(t *testing.T) {
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

	testHostResource := "customGpuHostResource"
	err = vmms.AttachGPU(vm, testHostResource)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Attached GPU-DDA to [%s] VMs", "test")

	err = vmms.DetachGPU(vm, testHostResource)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Detached GPU-DDA from [%s] VMs", "test")
}

func TestAddRemoveGpuDdaGen1(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	vm, err := vmms.GetVirtualMachineByName("testGen1")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer vm.Close()
	t.Logf("Found [%s] VMs", "testGen1")

	testHostResource := "customGpuHostResource"
	err = vmms.AttachGPU(vm, testHostResource)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Attached GPU-DDA to [%s] VMs", "testGen1")

	err = vmms.DetachGPU(vm, testHostResource)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	t.Logf("Detached GPU-DDA from [%s] VMs", "testGen1")
}
