// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package pciesetting

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type PcieDeviceSetting struct {
	*PciExpressSettingData
}

func NewPcieDeviceSetting(instance *wmi.WmiInstance) (*PcieDeviceSetting, error) {
	wmivm, err := NewPciExpressSettingData(instance)
	if err != nil {
		return nil, err
	}
	return &PcieDeviceSetting{wmivm}, nil
}

func NewPcieDeviceSettingEx6(whost *host.WmiHost) (*PcieDeviceSetting, error) {
	wmivm, err := NewPciExpressSettingDataEx6(whost)
	if err != nil {
		return nil, err
	}
	return &PcieDeviceSetting{wmivm}, nil
}

func GetDefaultPcieDeviceSetting(whost *host.WmiHost) (*PcieDeviceSetting, error) {
	inst, err := instance.CreateWmiInstance(whost, string(constant.Virtualization), "Msvm_PciExpressSettingData")
	if err != nil {
		return nil, err
	}
	return NewPcieDeviceSetting(inst)
}

func (pcie *PcieDeviceSetting) CloneEx1() (*PcieDeviceSetting, error) {
	tmp, err := pcie.Clone()
	if err != nil {
		return nil, err
	}
	return NewPcieDeviceSetting(tmp)
}

// GetVirtualMachine gets the VM that the pci express setting data is attached to
func (pcie *PcieDeviceSetting) GetVirtualMachine() (instance *wmi.WmiInstance, err error) {
	vmsetting, err := pcie.GetRelated("Msvm_VirtualSystemSettingData")
	if err != nil {
		return
	}
	defer vmsetting.Close()
	return vmsetting.GetRelated("Msvm_ComputerSystem")
}
