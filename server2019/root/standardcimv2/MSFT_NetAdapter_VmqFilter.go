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

// MSFT_NetAdapter_VmqFilter struct
type MSFT_NetAdapter_VmqFilter struct {
	cim.WmiInstance

	//
	FilterID uint32

	//
	MacAddress string

	//
	VlanID uint16
}

// SetFilterID sets the value of FilterID for the instance
func (instance *MSFT_NetAdapter_VmqFilter) SetPropertyFilterID(value uint32) (err error) {
	return instance.SetProperty("FilterID", value)
}

// GetFilterID gets the value of FilterID for the instance
func (instance *MSFT_NetAdapter_VmqFilter) GetPropertyFilterID() (value uint32, err error) {
	retValue, err := instance.GetProperty("FilterID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMacAddress sets the value of MacAddress for the instance
func (instance *MSFT_NetAdapter_VmqFilter) SetPropertyMacAddress(value string) (err error) {
	return instance.SetProperty("MacAddress", value)
}

// GetMacAddress gets the value of MacAddress for the instance
func (instance *MSFT_NetAdapter_VmqFilter) GetPropertyMacAddress() (value string, err error) {
	retValue, err := instance.GetProperty("MacAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVlanID sets the value of VlanID for the instance
func (instance *MSFT_NetAdapter_VmqFilter) SetPropertyVlanID(value uint16) (err error) {
	return instance.SetProperty("VlanID", value)
}

// GetVlanID gets the value of VlanID for the instance
func (instance *MSFT_NetAdapter_VmqFilter) GetPropertyVlanID() (value uint16, err error) {
	retValue, err := instance.GetProperty("VlanID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
