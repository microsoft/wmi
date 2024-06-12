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

// MSNdis_WmiOutputInfo struct
type MSNdis_WmiOutputInfo struct {
	*MSNdis

	//
	DataOffset uint32

	//
	Flags uint32

	//
	Header MSNdis_ObjectHeader

	//
	Padding1 uint8

	//
	Padding2 uint16

	//
	SupportedRevision uint8
}

func NewMSNdis_WmiOutputInfoEx1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiOutputInfo, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiOutputInfo{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_WmiOutputInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_WmiOutputInfo, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiOutputInfo{
		MSNdis: tmp,
	}
	return
}

// SetDataOffset sets the value of DataOffset for the instance
func (instance *MSNdis_WmiOutputInfo) SetPropertyDataOffset(value uint32) (err error) {
	return instance.SetProperty("DataOffset", (value))
}

// GetDataOffset gets the value of DataOffset for the instance
func (instance *MSNdis_WmiOutputInfo) GetPropertyDataOffset() (value uint32, err error) {
	retValue, err := instance.GetProperty("DataOffset")
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

// SetFlags sets the value of Flags for the instance
func (instance *MSNdis_WmiOutputInfo) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSNdis_WmiOutputInfo) GetPropertyFlags() (value uint32, err error) {
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
func (instance *MSNdis_WmiOutputInfo) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_WmiOutputInfo) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
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

// SetPadding1 sets the value of Padding1 for the instance
func (instance *MSNdis_WmiOutputInfo) SetPropertyPadding1(value uint8) (err error) {
	return instance.SetProperty("Padding1", (value))
}

// GetPadding1 gets the value of Padding1 for the instance
func (instance *MSNdis_WmiOutputInfo) GetPropertyPadding1() (value uint8, err error) {
	retValue, err := instance.GetProperty("Padding1")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetPadding2 sets the value of Padding2 for the instance
func (instance *MSNdis_WmiOutputInfo) SetPropertyPadding2(value uint16) (err error) {
	return instance.SetProperty("Padding2", (value))
}

// GetPadding2 gets the value of Padding2 for the instance
func (instance *MSNdis_WmiOutputInfo) GetPropertyPadding2() (value uint16, err error) {
	retValue, err := instance.GetProperty("Padding2")
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

// SetSupportedRevision sets the value of SupportedRevision for the instance
func (instance *MSNdis_WmiOutputInfo) SetPropertySupportedRevision(value uint8) (err error) {
	return instance.SetProperty("SupportedRevision", (value))
}

// GetSupportedRevision gets the value of SupportedRevision for the instance
func (instance *MSNdis_WmiOutputInfo) GetPropertySupportedRevision() (value uint8, err error) {
	retValue, err := instance.GetProperty("SupportedRevision")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}
