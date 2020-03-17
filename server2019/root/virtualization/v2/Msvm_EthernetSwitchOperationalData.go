// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_EthernetSwitchOperationalData struct
type Msvm_EthernetSwitchOperationalData struct {
	Msvm_EthernetSwitchData

	//
	CurrentSwitchingMode uint32

	//
	SupportedSwitchingModes []uint32
}

// SetCurrentSwitchingMode sets the value of CurrentSwitchingMode for the instance
func (instance *Msvm_EthernetSwitchOperationalData) SetPropertyCurrentSwitchingMode(value uint32) (err error) {
	return instance.SetProperty("CurrentSwitchingMode", value)
}

// GetCurrentSwitchingMode gets the value of CurrentSwitchingMode for the instance
func (instance *Msvm_EthernetSwitchOperationalData) GetPropertyCurrentSwitchingMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentSwitchingMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedSwitchingModes sets the value of SupportedSwitchingModes for the instance
func (instance *Msvm_EthernetSwitchOperationalData) SetPropertySupportedSwitchingModes(value []uint32) (err error) {
	return instance.SetProperty("SupportedSwitchingModes", value)
}

// GetSupportedSwitchingModes gets the value of SupportedSwitchingModes for the instance
func (instance *Msvm_EthernetSwitchOperationalData) GetPropertySupportedSwitchingModes() (value []uint32, err error) {
	retValue, err := instance.GetProperty("SupportedSwitchingModes")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_EthernetSwitchOperationalData) GetRelatedVirtualEthernetSwitch() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitch")
}
