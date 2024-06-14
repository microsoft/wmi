// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// MSFT_NetAdapterPowerManagement_WakePattern struct
type MSFT_NetAdapterPowerManagement_WakePattern struct {
	*cim.WmiInstance

	//
	FriendlyName string

	//
	ID uint32

	//
	Priority uint32

	//
	WakePacketType uint32
}

func NewMSFT_NetAdapterPowerManagement_WakePatternEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterPowerManagement_WakePattern, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterPowerManagement_WakePattern{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapterPowerManagement_WakePatternEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterPowerManagement_WakePattern, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterPowerManagement_WakePattern{
		WmiInstance: tmp,
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", (value))
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern) GetPropertyFriendlyName() (value string, err error) {
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
func (instance *MSFT_NetAdapterPowerManagement_WakePattern) SetPropertyID(value uint32) (err error) {
	return instance.SetProperty("ID", (value))
}

// GetID gets the value of ID for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern) GetPropertyID() (value uint32, err error) {
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

// SetPriority sets the value of Priority for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern) SetPropertyPriority(value uint32) (err error) {
	return instance.SetProperty("Priority", (value))
}

// GetPriority gets the value of Priority for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern) GetPropertyPriority() (value uint32, err error) {
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

// SetWakePacketType sets the value of WakePacketType for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern) SetPropertyWakePacketType(value uint32) (err error) {
	return instance.SetProperty("WakePacketType", (value))
}

// GetWakePacketType gets the value of WakePacketType for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern) GetPropertyWakePacketType() (value uint32, err error) {
	retValue, err := instance.GetProperty("WakePacketType")
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
