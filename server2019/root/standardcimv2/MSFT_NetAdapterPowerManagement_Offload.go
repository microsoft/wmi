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

// MSFT_NetAdapterPowerManagement_Offload struct
type MSFT_NetAdapterPowerManagement_Offload struct {
	cim.WmiInstance

	//
	FriendlyName string

	//
	ID uint32

	//
	OffloadType uint32

	//
	Priority uint32
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", value)
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload) GetPropertyFriendlyName() (value string, err error) {
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
func (instance *MSFT_NetAdapterPowerManagement_Offload) SetPropertyID(value uint32) (err error) {
	return instance.SetProperty("ID", value)
}

// GetID gets the value of ID for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload) GetPropertyID() (value uint32, err error) {
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

// SetOffloadType sets the value of OffloadType for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload) SetPropertyOffloadType(value uint32) (err error) {
	return instance.SetProperty("OffloadType", value)
}

// GetOffloadType gets the value of OffloadType for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload) GetPropertyOffloadType() (value uint32, err error) {
	retValue, err := instance.GetProperty("OffloadType")
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
func (instance *MSFT_NetAdapterPowerManagement_Offload) SetPropertyPriority(value uint32) (err error) {
	return instance.SetProperty("Priority", value)
}

// GetPriority gets the value of Priority for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload) GetPropertyPriority() (value uint32, err error) {
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
