// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystemmanagement

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
	setting, err := virtualsystem.GetVirtualSystemSettingData(whost, "temp")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer setting.Close()
	t.Logf("Create VMSettings")

	_, err = vmms.CreateVirtualMachine(setting)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	t.Logf("Created VM [%s]", "temp")
	//defer vm.Close()
}

func TestVirtualMachineDelete(t *testing.T) {
	vmms, err := GetVirtualSystemManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	vm, err := vmms.GetVirtualMachineByName("temp")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	defer vm.Close()
	err = vmms.DeleteVirtualMachine(vm)
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
