// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
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

// MSFT_NetAdapterPowerManagement_Offload struct
type MSFT_NetAdapterPowerManagement_Offload struct {
	*cim.WmiInstance

	//
	FriendlyName string

	//
	ID uint32

	//
	OffloadType uint32

	//
	Priority uint32
}

func NewMSFT_NetAdapterPowerManagement_OffloadEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterPowerManagement_Offload, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterPowerManagement_Offload{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapterPowerManagement_OffloadEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterPowerManagement_Offload, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterPowerManagement_Offload{
		WmiInstance: tmp,
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", (value))
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
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

// SetID sets the value of ID for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload) SetPropertyID(value uint32) (err error) {
	return instance.SetProperty("ID", (value))
}

// GetID gets the value of ID for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload) GetPropertyID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ID")
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

// SetOffloadType sets the value of OffloadType for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload) SetPropertyOffloadType(value uint32) (err error) {
	return instance.SetProperty("OffloadType", (value))
}

// GetOffloadType gets the value of OffloadType for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload) GetPropertyOffloadType() (value uint32, err error) {
	retValue, err := instance.GetProperty("OffloadType")
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

// SetPriority sets the value of Priority for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload) SetPropertyPriority(value uint32) (err error) {
	return instance.SetProperty("Priority", (value))
}

// GetPriority gets the value of Priority for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload) GetPropertyPriority() (value uint32, err error) {
	retValue, err := instance.GetProperty("Priority")
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
