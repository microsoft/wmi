// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"fmt"
	"os"
	"testing"

	"github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
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
		t.Fatal("Failed " + err.Error())
	}
}

func TestCreateVirtualMachines(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	setting, err := virtualsystem.GetVirtualSystemSettingData(whost, "test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer setting.Close()
	t.Logf("Create VMSettings")

	_, err = vmms.CreateVirtualMachine(setting)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	t.Logf("Created VM [%s]", "test")
	//defer vm.Close()
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

	vs, err := virtualswitch.GetVirtualSwitch(whost, "test")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer vs.Close()
	t.Logf("Got VirtualSwitch[%s]", "test")

	vm, err := vmms.GetVirtualMachineByName("test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	t.Logf("Found [%s] VMs", "test")
	defer vm.Close()

	for i := 1; i <= 4; i++ {
		adapterName := fmt.Sprintf("testadapter-%d", i)
		t.Logf("Adding VmNic to [%s]", "test")
		na, err := vmms.AddVirtualNetworkAdapter(vm, "Network Adapter")
		if err != nil {
			t.Fatal("Failed " + err.Error())
		}
		defer na.Close()

		err = vmms.RenameVirtualNetworkAdapter(vm, "Network Adapter", adapterName)
		if err != nil {
			t.Fatal("Failed " + err.Error())
		}
		t.Logf("Renamed Adapter [%s]", adapterName)

		testna, err := vm.GetVirtualNetworkAdapterByName(adapterName)
		if err != nil {
			t.Fatal("Failed " + err.Error())
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
			"1fa41b39-b444-b35a-e1f7985fd548",
			"00000000-0000-0000-000000000000", 1)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		t.Logf("Set Adapter VLAN [%d]", i)

	}
	for i := 1; i <= 4; i++ {
		adapterName := fmt.Sprintf("testadapter-%d", i)
		err = vmms.DisconnectAdapterFromVirtualSwitch(vm, adapterName)
		if err != nil {
			t.Fatal("Failed " + err.Error())
		}
		t.Logf("Disconnect VM[%s] from VirtualSwitch[%s]", "test", "test")
		testna, err := vm.GetVirtualNetworkAdapterByName(adapterName)
		if err != nil {
			t.Fatal("Failed " + err.Error())
		}
		defer testna.Close()
		t.Logf("Removing VmNic [%s] from VM[%s] ", adapterName, "test")
		err = vmms.RemoveVirtualNetworkAdapter(testna)
		if err != nil {
			t.Fatal("Failed " + err.Error())

		}
	}
}

func TestAddRemoveVirtualHardDisk(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	vm, err := vmms.GetVirtualMachineByName("test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer vm.Close()
	t.Logf("Found [%s] VMs", "test")

	err = vmms.AddSCSIController(vm)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}

	ims, err := service.GetImageManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	t.Logf("Got ImageManagementService ")

	for i := 1; i <= 4; i++ {
		path := fmt.Sprintf("c:\\test\\tmp-%d.vhdx", i)
		setting, err := disk.GetVirtualHardDiskSettingData(whost, path, 512, 512, 0, 1024*1024*10, true)
		if err != nil {
			t.Fatal("Failed " + err.Error())
		}
		defer setting.Close()
		err = ims.CreateDisk(setting)
		if err != nil {
			t.Fatal("Create Failed " + err.Error())
		}
		defer os.RemoveAll(path)
		t.Logf("Created vhd [%s]", path)

		vhd, vhddrive, err := vmms.AttachVirtualHardDisk(vm, path)
		if err != nil {
			t.Fatal("Failed " + err.Error())
		}
		defer vhd.Close()
		defer vhddrive.Close()
		t.Logf("Attached vhd [%s] to [%s]", path, "test")
		controllerlocation, err := vhddrive.GetControllerLocation()
		if err != nil {
			t.Fatal("Failed " + err.Error())
		}
		controllerNumber, err := vhddrive.GetControllerNumber()
		if err != nil {
			t.Fatal("Failed " + err.Error())
		}
		t.Logf("ControllerNumber [%s], ControllerLocation [%s]", controllerNumber, controllerlocation)
	}

	for i := 1; i <= 4; i++ {
		path := fmt.Sprintf("c:\\test\\tmp-%d.vhdx", i)
		vhd, err := vm.GetVirtualHardDiskByLocation(0, i-1)
		if err != nil {
			t.Fatal("Failed " + err.Error())
		}
		defer vhd.Close()
		err = vmms.DetachVirtualHardDisk(vhd)
		if err != nil {
			t.Fatal("Failed " + err.Error())
		}
		t.Logf("Detached vhd [%s] from [%s]", path, "test")
	}
}

func TestAddRemoveTPM(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
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
