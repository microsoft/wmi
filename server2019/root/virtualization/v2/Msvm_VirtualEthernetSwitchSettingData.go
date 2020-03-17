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

// Msvm_VirtualEthernetSwitchSettingData struct
type Msvm_VirtualEthernetSwitchSettingData struct {
	CIM_VirtualEthernetSwitchSettingData

	//
	BandwidthReservationMode uint32

	//
	ExtensionOrder []string

	//
	IOVPreferred bool

	//
	PacketDirectEnabled bool

	//
	TeamingEnabled bool
}

// SetBandwidthReservationMode sets the value of BandwidthReservationMode for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) SetPropertyBandwidthReservationMode(value uint32) (err error) {
	return instance.SetProperty("BandwidthReservationMode", value)
}

// GetBandwidthReservationMode gets the value of BandwidthReservationMode for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) GetPropertyBandwidthReservationMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("BandwidthReservationMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExtensionOrder sets the value of ExtensionOrder for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) SetPropertyExtensionOrder(value []string) (err error) {
	return instance.SetProperty("ExtensionOrder", value)
}

// GetExtensionOrder gets the value of ExtensionOrder for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) GetPropertyExtensionOrder() (value []string, err error) {
	retValue, err := instance.GetProperty("ExtensionOrder")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOVPreferred sets the value of IOVPreferred for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) SetPropertyIOVPreferred(value bool) (err error) {
	return instance.SetProperty("IOVPreferred", value)
}

// GetIOVPreferred gets the value of IOVPreferred for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) GetPropertyIOVPreferred() (value bool, err error) {
	retValue, err := instance.GetProperty("IOVPreferred")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketDirectEnabled sets the value of PacketDirectEnabled for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) SetPropertyPacketDirectEnabled(value bool) (err error) {
	return instance.SetProperty("PacketDirectEnabled", value)
}

// GetPacketDirectEnabled gets the value of PacketDirectEnabled for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) GetPropertyPacketDirectEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("PacketDirectEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTeamingEnabled sets the value of TeamingEnabled for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) SetPropertyTeamingEnabled(value bool) (err error) {
	return instance.SetProperty("TeamingEnabled", value)
}

// GetTeamingEnabled gets the value of TeamingEnabled for the instance
func (instance *Msvm_VirtualEthernetSwitchSettingData) GetPropertyTeamingEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("TeamingEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_VirtualEthernetSwitchSettingData) GetRelatedEthernetPortAllocationSettingData() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_EthernetPortAllocationSettingData")
}

func (instance *Msvm_VirtualEthernetSwitchSettingData) GetRelatedVirtualEthernetSwitch() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitch")
}
