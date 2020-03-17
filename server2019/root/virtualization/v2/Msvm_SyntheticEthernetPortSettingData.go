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

// Msvm_SyntheticEthernetPortSettingData struct
type Msvm_SyntheticEthernetPortSettingData struct {
	CIM_EthernetPortAllocationSettingData

	//
	AllowPacketDirect bool

	//
	ClusterMonitored bool

	//
	DeviceNamingEnabled bool

	//
	InterruptModeration bool

	//
	MediaType uint32

	//
	NumaAwarePlacement bool

	//
	StaticMacAddress bool

	//
	VirtualSystemIdentifiers []string
}

// SetAllowPacketDirect sets the value of AllowPacketDirect for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) SetPropertyAllowPacketDirect(value bool) (err error) {
	return instance.SetProperty("AllowPacketDirect", value)
}

// GetAllowPacketDirect gets the value of AllowPacketDirect for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) GetPropertyAllowPacketDirect() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowPacketDirect")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetClusterMonitored sets the value of ClusterMonitored for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) SetPropertyClusterMonitored(value bool) (err error) {
	return instance.SetProperty("ClusterMonitored", value)
}

// GetClusterMonitored gets the value of ClusterMonitored for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) GetPropertyClusterMonitored() (value bool, err error) {
	retValue, err := instance.GetProperty("ClusterMonitored")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceNamingEnabled sets the value of DeviceNamingEnabled for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) SetPropertyDeviceNamingEnabled(value bool) (err error) {
	return instance.SetProperty("DeviceNamingEnabled", value)
}

// GetDeviceNamingEnabled gets the value of DeviceNamingEnabled for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) GetPropertyDeviceNamingEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DeviceNamingEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterruptModeration sets the value of InterruptModeration for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) SetPropertyInterruptModeration(value bool) (err error) {
	return instance.SetProperty("InterruptModeration", value)
}

// GetInterruptModeration gets the value of InterruptModeration for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) GetPropertyInterruptModeration() (value bool, err error) {
	retValue, err := instance.GetProperty("InterruptModeration")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMediaType sets the value of MediaType for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) SetPropertyMediaType(value uint32) (err error) {
	return instance.SetProperty("MediaType", value)
}

// GetMediaType gets the value of MediaType for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) GetPropertyMediaType() (value uint32, err error) {
	retValue, err := instance.GetProperty("MediaType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumaAwarePlacement sets the value of NumaAwarePlacement for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) SetPropertyNumaAwarePlacement(value bool) (err error) {
	return instance.SetProperty("NumaAwarePlacement", value)
}

// GetNumaAwarePlacement gets the value of NumaAwarePlacement for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) GetPropertyNumaAwarePlacement() (value bool, err error) {
	retValue, err := instance.GetProperty("NumaAwarePlacement")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStaticMacAddress sets the value of StaticMacAddress for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) SetPropertyStaticMacAddress(value bool) (err error) {
	return instance.SetProperty("StaticMacAddress", value)
}

// GetStaticMacAddress gets the value of StaticMacAddress for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) GetPropertyStaticMacAddress() (value bool, err error) {
	retValue, err := instance.GetProperty("StaticMacAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualSystemIdentifiers sets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) SetPropertyVirtualSystemIdentifiers(value []string) (err error) {
	return instance.SetProperty("VirtualSystemIdentifiers", value)
}

// GetVirtualSystemIdentifiers gets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_SyntheticEthernetPortSettingData) GetPropertyVirtualSystemIdentifiers() (value []string, err error) {
	retValue, err := instance.GetProperty("VirtualSystemIdentifiers")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_SyntheticEthernetPortSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
