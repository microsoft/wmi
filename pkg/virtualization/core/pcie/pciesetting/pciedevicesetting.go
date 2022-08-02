// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package pciesetting

import (
	"github.com/microsoft/wmi/pkg/base/host"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type PcieDeviceSetting struct {
	*PciExpressSettingData
}

func NewPcieDeviceSettingEx1(instance *wmi.WmiInstance) (*PcieDeviceSetting, error) {
	wmivm, err := NewPciExpressSettingDataEx1(instance)
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

func (pci *PcieDeviceSetting) GetPciExpressSettingData() (*PciExpressSettingData, error) {
	tmp, err := pci.GetRelated("Msvm_PciExpressSettingData")
	if err != nil {
		return nil, err
	}
	return NewPciExpressSettingDataEx1(tmp)
}

func (pcie *PcieDeviceSetting) CloneEx1() (*PcieDeviceSetting, error) {
	tmp, err := pcie.Clone()
	if err != nil {
		return nil, err
	}
	return NewPcieDeviceSettingEx1(tmp)
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
