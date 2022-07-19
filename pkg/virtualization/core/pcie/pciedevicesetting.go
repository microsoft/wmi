// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package pcie

import (
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

func (pci *PcieDeviceSetting) GetPciExpressSettingData() (*PciExpressSettingData, error) {
	tmp, err := pci.GetRelated("Msvm_PciExpressSettingData")
	if err != nil {
		return nil, err
	}
	return NewPciExpressSettingData(tmp)
}

func (pci *PcieDeviceSetting) CloneEx1() (*PcieDeviceSetting, error) {
	tmp, err := pci.Clone()
	if err != nil {
		return nil, err
	}
	return NewPcieDeviceSetting(tmp)
}

// GetVirtualMachine gets the VM that the pci express is attached to
func (pci *PcieDeviceSetting) GetVirtualMachine() (instance *wmi.WmiInstance, err error) {
	vmsetting, err := pci.GetRelated("Msvm_VirtualSystemSettingData")
	if err != nil {
		return
	}
	defer vmsetting.Close()
	return vmsetting.GetRelated("Msvm_ComputerSystem")
}
