// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

import (
	"testing"

	"github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	memory "github.com/microsoft/wmi/pkg/virtualization/core/memory"
)

var (
	memsettings *VirtualSystemSettingData
)

func init() {
	whost = host.NewWmiLocalHost()
	vm, err := GetVirtualMachineByVMName(whost, "test")
	if err != nil {
		return
	}

	memsettings, err = vm.GetVirtualSystemSettingData()
	if err != nil {
		return
	}
}

func TestGetMemorySettingData(t *testing.T) {
	defer memsettings.Close()

	msd, err := memsettings.GetRelated("Msvm_MemorySettingData")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	memorySettingData, err := memory.NewMemorySettingData(msd)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	memorySettingData.Close()
}

func TestGetDefaultMemorySettingData(t *testing.T) {
	defer memsettings.Close()

	memorySettingData, err := memory.GetDefaultMemorySettingData(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	memorySettingData.Close()
}

func TestGetSizeMb(t *testing.T) {
	defer memsettings.Close()

	msd, err := memsettings.GetRelated("Msvm_MemorySettingData")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	memorySettingData, err := memory.NewMemorySettingData(msd)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	size, err := memorySettingData.GetSizeMB()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	} else if size == 0 {
		t.Fatal("Failed, retrieved size 0.")
	}

	memorySettingData.Close()
}

func TestGetMaximumMemoryMB(t *testing.T) {
	defer memsettings.Close()

	msd, err := memsettings.GetRelated("Msvm_MemorySettingData")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	memorySettingData, err := memory.NewMemorySettingData(msd)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	size, err := memorySettingData.GetMaximumMemoryMB()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	} else if size == 0 {
		t.Fatal("Failed, retrieved size 0.")
	}

	memorySettingData.Close()
}

func TestGetMinimumMemoryMB(t *testing.T) {
	defer memsettings.Close()

	msd, err := memsettings.GetRelated("Msvm_MemorySettingData")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	memorySettingData, err := memory.NewMemorySettingData(msd)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	size, err := memorySettingData.GetMinimumMemoryMB()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	} else if size == 0 {
		t.Fatal("Failed, retrieved size 0.")
	}

	memorySettingData.Close()
}

func TestGetTargetMemoryBuffer(t *testing.T) {
	defer memsettings.Close()

	msd, err := memsettings.GetRelated("Msvm_MemorySettingData")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	memorySettingData, err := memory.NewMemorySettingData(msd)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	size, err := memorySettingData.GetTargetMemoryBuffer()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	} else if size == 0 {
		t.Fatal("Failed, retrieved size 0.")
	}

	memorySettingData.Close()
}
