// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package switchport

import (
	"strings"

	"github.com/microsoft/wmi/pkg/errors"
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

func (epas *EthernetPortAllocationSettingData) GetEthernetSwitchPortProfileSettingDataForVendor(vendorGuid string) (*EthernetSwitchPortProfileSettingData, error) {

	col, err := epas.GetAssociated("Msvm_EthernetPortSettingDataComponent", "Msvm_EthernetSwitchPortProfileSettingData", "", "")

	if err != nil {
		return nil, err
	}
	defer col.Close()
	wmiVendorId := "{" + vendorGuid + "}"
	for _, item := range col {
		v, err := item.GetProperty("VendorId")
		if err != nil {
			continue
		}
		currentVendor := v.(string)
		if strings.EqualFold(currentVendor, wmiVendorId) {
			profileSettings, err := item.Clone()
			if err != nil {
				continue
			}
			return NewEthernetSwitchPortProfileSettingData(profileSettings)
		}
	}
	return nil, errors.Wrapf(errors.NotFound, "No item found for vendor id")
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

func (epas *EthernetPortAllocationSettingData) GetEthernetSwitchPortVLANSettingData() (*EthernetSwitchPortVLANSettingData, error) {

	tmp, err := epas.GetFirstRelatedEx("",
		"Msvm_EthernetSwitchPortVLANSettingData",
		"",
		"")
	if err != nil {
		return nil, err
	}

	return NewEthernetSwitchPortVLANSettingData(tmp)
}

func (epas *EthernetPortAllocationSettingData) GetEthernetSwitchPortIsolationSettingData() (*EthernetSwitchPortIsolationSettingData, error) {

	tmp, err := epas.GetFirstRelatedEx("",
		"Msvm_EthernetSwitchPortIsolationSettingData",
		"",
		"")
	if err != nil {
		return nil, err
	}

	return NewEthernetSwitchPortIsolationSettingData(tmp)
}
