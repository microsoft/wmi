// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_NicSwitchCapabilities struct
type MSNdis_NicSwitchCapabilities struct {
	*MSNdis

	//
	Flags uint32

	//
	Header MSNdis_ObjectHeader

	//
	NdisReserved1 uint32

	//
	NdisReserved2 uint32

	//
	NdisReserved3 uint32

	//
	NumMacAddressesPerPort uint32

	//
	NumTotalMacAddresses uint32

	//
	NumVlansPerPort uint32
}

func NewMSNdis_NicSwitchCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSNdis_NicSwitchCapabilities, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_NicSwitchCapabilities{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_NicSwitchCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_NicSwitchCapabilities, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_NicSwitchCapabilities{
		MSNdis: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MSNdis_NicSwitchCapabilities) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSNdis_NicSwitchCapabilities) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
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

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_NicSwitchCapabilities) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_NicSwitchCapabilities) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
	retValue, err := instance.GetProperty("Header")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_ObjectHeader)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_ObjectHeader is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_ObjectHeader(valuetmp)

	return
}

// SetNdisReserved1 sets the value of NdisReserved1 for the instance
func (instance *MSNdis_NicSwitchCapabilities) SetPropertyNdisReserved1(value uint32) (err error) {
	return instance.SetProperty("NdisReserved1", (value))
}

// GetNdisReserved1 gets the value of NdisReserved1 for the instance
func (instance *MSNdis_NicSwitchCapabilities) GetPropertyNdisReserved1() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdisReserved1")
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

// SetNdisReserved2 sets the value of NdisReserved2 for the instance
func (instance *MSNdis_NicSwitchCapabilities) SetPropertyNdisReserved2(value uint32) (err error) {
	return instance.SetProperty("NdisReserved2", (value))
}

// GetNdisReserved2 gets the value of NdisReserved2 for the instance
func (instance *MSNdis_NicSwitchCapabilities) GetPropertyNdisReserved2() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdisReserved2")
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

// SetNdisReserved3 sets the value of NdisReserved3 for the instance
func (instance *MSNdis_NicSwitchCapabilities) SetPropertyNdisReserved3(value uint32) (err error) {
	return instance.SetProperty("NdisReserved3", (value))
}

// GetNdisReserved3 gets the value of NdisReserved3 for the instance
func (instance *MSNdis_NicSwitchCapabilities) GetPropertyNdisReserved3() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdisReserved3")
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

// SetNumMacAddressesPerPort sets the value of NumMacAddressesPerPort for the instance
func (instance *MSNdis_NicSwitchCapabilities) SetPropertyNumMacAddressesPerPort(value uint32) (err error) {
	return instance.SetProperty("NumMacAddressesPerPort", (value))
}

// GetNumMacAddressesPerPort gets the value of NumMacAddressesPerPort for the instance
func (instance *MSNdis_NicSwitchCapabilities) GetPropertyNumMacAddressesPerPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumMacAddressesPerPort")
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

// SetNumTotalMacAddresses sets the value of NumTotalMacAddresses for the instance
func (instance *MSNdis_NicSwitchCapabilities) SetPropertyNumTotalMacAddresses(value uint32) (err error) {
	return instance.SetProperty("NumTotalMacAddresses", (value))
}

// GetNumTotalMacAddresses gets the value of NumTotalMacAddresses for the instance
func (instance *MSNdis_NicSwitchCapabilities) GetPropertyNumTotalMacAddresses() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumTotalMacAddresses")
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

// SetNumVlansPerPort sets the value of NumVlansPerPort for the instance
func (instance *MSNdis_NicSwitchCapabilities) SetPropertyNumVlansPerPort(value uint32) (err error) {
	return instance.SetProperty("NumVlansPerPort", (value))
}

// GetNumVlansPerPort gets the value of NumVlansPerPort for the instance
func (instance *MSNdis_NicSwitchCapabilities) GetPropertyNumVlansPerPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumVlansPerPort")
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
