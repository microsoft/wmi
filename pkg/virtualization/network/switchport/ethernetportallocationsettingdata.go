// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package switchport

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type EthernetPortAllocationSettingData struct {
	*v2.Msvm_EthernetPortAllocationSettingData
}

// NewEthernetPortAllocationSettingData
func NewEthernetPortAllocationSettingData(instance *wmi.WmiInstance) (*EthernetPortAllocationSettingData, error) {
	wmivm, err := v2.NewMsvm_EthernetPortAllocationSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &EthernetPortAllocationSettingData{wmivm}, nil
}

func (epas *EthernetPortAllocationSettingData) CloneEx1() (epasnew *EthernetPortAllocationSettingData, err error) {
	tmp, err := epas.Clone()
	if err != nil {
		return
	}
	return NewEthernetPortAllocationSettingData(tmp)
}

func (epas *EthernetPortAllocationSettingData) GetEthernetSwitchPortProfileSettingData() (*EthernetSwitchPortProfileSettingData, error) {
	tmp, err := epas.GetFirstRelatedEx("Msvm_EthernetSwitchPortProfileSettingData",
		"Msvm_EthernetPortSettingDataComponent",
		"PartComponent",
		"GroupComponent")
	if err != nil {
		return nil, err
	}

	return NewEthernetSwitchPortProfileSettingData(tmp)
}

func (epas *EthernetPortAllocationSettingData) GetEthernetSwitchPortOffloadSettingData() (*EthernetSwitchPortOffloadSettingData, error) {
	tmp, err := epas.GetFirstRelatedEx("",
		"Msvm_EthernetSwitchPortOffloadSettingData",
		"",
		"")
	if err != nil {
		return nil, err
	}

	return NewEthernetSwitchPortOffloadSettingData(tmp)
}
