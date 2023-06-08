// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

import (
	"testing"

	"github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	processor "github.com/microsoft/wmi/pkg/virtualization/core/processor"
)

var (
	procsettings *VirtualSystemSettingData
)

func init() {
	whost = host.NewWmiLocalHost()
	vm, err := GetVirtualMachineByVMName(whost, "test")
	if err != nil {
		return
	}

	procsettings, err = vm.GetVirtualSystemSettingData()
	if err != nil {
		return
	}
}

func TestGetProcessorSettingData(t *testing.T) {

	psd, err := procsettings.GetRelated("Msvm_ProcessorSettingData")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	_, err = processor.NewProcessorSettingData(psd)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
}

func TestGetDefaultProcessorSettingData(t *testing.T) {

	_, err := processor.GetDefaultProcessorSettingData(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
}

func TestGetCpuCount(t *testing.T) {

	psd, err := procsettings.GetRelated("Msvm_ProcessorSettingData")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	processorSettingData, err := processor.NewProcessorSettingData(psd)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	count, err := processorSettingData.GetCPUCount()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	} else if count == 0 {
		t.Fatal("Failed, retrieved cpu count 0.")
	}
}
