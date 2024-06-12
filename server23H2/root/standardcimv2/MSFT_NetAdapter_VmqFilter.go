// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapter_VmqFilter struct
type MSFT_NetAdapter_VmqFilter struct {
	*cim.WmiInstance

	//
	FilterID uint32

	//
	MacAddress string

	//
	VlanID uint16
}

func NewMSFT_NetAdapter_VmqFilterEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapter_VmqFilter, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_VmqFilter{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapter_VmqFilterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapter_VmqFilter, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_VmqFilter{
		WmiInstance: tmp,
	}
	return
}

// SetFilterID sets the value of FilterID for the instance
func (instance *MSFT_NetAdapter_VmqFilter) SetPropertyFilterID(value uint32) (err error) {
	return instance.SetProperty("FilterID", (value))
}

// GetFilterID gets the value of FilterID for the instance
func (instance *MSFT_NetAdapter_VmqFilter) GetPropertyFilterID() (value uint32, err error) {
	retValue, err := instance.GetProperty("FilterID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetMacAddress sets the value of MacAddress for the instance
func (instance *MSFT_NetAdapter_VmqFilter) SetPropertyMacAddress(value string) (err error) {
	return instance.SetProperty("MacAddress", (value))
}

// GetMacAddress gets the value of MacAddress for the instance
func (instance *MSFT_NetAdapter_VmqFilter) GetPropertyMacAddress() (value string, err error) {
	retValue, err := instance.GetProperty("MacAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetVlanID sets the value of VlanID for the instance
func (instance *MSFT_NetAdapter_VmqFilter) SetPropertyVlanID(value uint16) (err error) {
	return instance.SetProperty("VlanID", (value))
}

// GetVlanID gets the value of VlanID for the instance
func (instance *MSFT_NetAdapter_VmqFilter) GetPropertyVlanID() (value uint16, err error) {
	retValue, err := instance.GetProperty("VlanID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
