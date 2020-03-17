// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapterPowerManagement_WakePattern struct
type MSFT_NetAdapterPowerManagement_WakePattern struct {
	cim.WmiInstance

	//
	FriendlyName string

	//
	ID uint32

	//
	Priority uint32

	//
	WakePacketType uint32
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", value)
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetID sets the value of ID for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern) SetPropertyID(value uint32) (err error) {
	return instance.SetProperty("ID", value)
}

// GetID gets the value of ID for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern) GetPropertyID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPriority sets the value of Priority for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern) SetPropertyPriority(value uint32) (err error) {
	return instance.SetProperty("Priority", value)
}

// GetPriority gets the value of Priority for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern) GetPropertyPriority() (value uint32, err error) {
	retValue, err := instance.GetProperty("Priority")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWakePacketType sets the value of WakePacketType for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern) SetPropertyWakePacketType(value uint32) (err error) {
	return instance.SetProperty("WakePacketType", value)
}

// GetWakePacketType gets the value of WakePacketType for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern) GetPropertyWakePacketType() (value uint32, err error) {
	retValue, err := instance.GetProperty("WakePacketType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
