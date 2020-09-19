// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetLldpMsap struct
type MSFT_NetLldpMsap struct {
	*MSFT_NetSettingData

	//
	AddressScope uint32

	//
	InterfaceAlias string

	//
	InterfaceIndex uint32

	//
	PhysicalAddress []string
}

func NewMSFT_NetLldpMsapEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetLldpMsap, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetLldpMsap{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetLldpMsapEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetLldpMsap, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetLldpMsap{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetAddressScope sets the value of AddressScope for the instance
func (instance *MSFT_NetLldpMsap) SetPropertyAddressScope(value uint32) (err error) {
	return instance.SetProperty("AddressScope", (value))
}

// GetAddressScope gets the value of AddressScope for the instance
func (instance *MSFT_NetLldpMsap) GetPropertyAddressScope() (value uint32, err error) {
	retValue, err := instance.GetProperty("AddressScope")
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

// SetInterfaceAlias sets the value of InterfaceAlias for the instance
func (instance *MSFT_NetLldpMsap) SetPropertyInterfaceAlias(value string) (err error) {
	return instance.SetProperty("InterfaceAlias", (value))
}

// GetInterfaceAlias gets the value of InterfaceAlias for the instance
func (instance *MSFT_NetLldpMsap) GetPropertyInterfaceAlias() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceAlias")
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

// SetInterfaceIndex sets the value of InterfaceIndex for the instance
func (instance *MSFT_NetLldpMsap) SetPropertyInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("InterfaceIndex", (value))
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_NetLldpMsap) GetPropertyInterfaceIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceIndex")
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

// SetPhysicalAddress sets the value of PhysicalAddress for the instance
func (instance *MSFT_NetLldpMsap) SetPropertyPhysicalAddress(value []string) (err error) {
	return instance.SetProperty("PhysicalAddress", (value))
}

// GetPhysicalAddress gets the value of PhysicalAddress for the instance
func (instance *MSFT_NetLldpMsap) GetPropertyPhysicalAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("PhysicalAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
