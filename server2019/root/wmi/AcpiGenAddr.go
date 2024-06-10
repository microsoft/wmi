// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// AcpiGenAddr struct
type AcpiGenAddr struct {
	*cim.WmiInstance

	//
	Address uint64

	//
	AddressSpaceID uint32

	//
	BitOffset uint32

	//
	BitWidth uint32

	//
	Reserved uint32
}

func NewAcpiGenAddrEx1(instance *cim.WmiInstance) (newInstance *AcpiGenAddr, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &AcpiGenAddr{
		WmiInstance: tmp,
	}
	return
}

func NewAcpiGenAddrEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AcpiGenAddr, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AcpiGenAddr{
		WmiInstance: tmp,
	}
	return
}

// SetAddress sets the value of Address for the instance
func (instance *AcpiGenAddr) SetPropertyAddress(value uint64) (err error) {
	return instance.SetProperty("Address", (value))
}

// GetAddress gets the value of Address for the instance
func (instance *AcpiGenAddr) GetPropertyAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("Address")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAddressSpaceID sets the value of AddressSpaceID for the instance
func (instance *AcpiGenAddr) SetPropertyAddressSpaceID(value uint32) (err error) {
	return instance.SetProperty("AddressSpaceID", (value))
}

// GetAddressSpaceID gets the value of AddressSpaceID for the instance
func (instance *AcpiGenAddr) GetPropertyAddressSpaceID() (value uint32, err error) {
	retValue, err := instance.GetProperty("AddressSpaceID")
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

// SetBitOffset sets the value of BitOffset for the instance
func (instance *AcpiGenAddr) SetPropertyBitOffset(value uint32) (err error) {
	return instance.SetProperty("BitOffset", (value))
}

// GetBitOffset gets the value of BitOffset for the instance
func (instance *AcpiGenAddr) GetPropertyBitOffset() (value uint32, err error) {
	retValue, err := instance.GetProperty("BitOffset")
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

// SetBitWidth sets the value of BitWidth for the instance
func (instance *AcpiGenAddr) SetPropertyBitWidth(value uint32) (err error) {
	return instance.SetProperty("BitWidth", (value))
}

// GetBitWidth gets the value of BitWidth for the instance
func (instance *AcpiGenAddr) GetPropertyBitWidth() (value uint32, err error) {
	retValue, err := instance.GetProperty("BitWidth")
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

// SetReserved sets the value of Reserved for the instance
func (instance *AcpiGenAddr) SetPropertyReserved(value uint32) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *AcpiGenAddr) GetPropertyReserved() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reserved")
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
