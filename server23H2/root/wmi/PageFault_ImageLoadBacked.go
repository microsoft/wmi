// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// PageFault_ImageLoadBacked struct
type PageFault_ImageLoadBacked struct {
	*PageFault_V2

	//
	DeviceChar uint32

	//
	FileChar uint16

	//
	FileObject uint32

	//
	LoadFlags uint16
}

func NewPageFault_ImageLoadBackedEx1(instance *cim.WmiInstance) (newInstance *PageFault_ImageLoadBacked, err error) {
	tmp, err := NewPageFault_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &PageFault_ImageLoadBacked{
		PageFault_V2: tmp,
	}
	return
}

func NewPageFault_ImageLoadBackedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PageFault_ImageLoadBacked, err error) {
	tmp, err := NewPageFault_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PageFault_ImageLoadBacked{
		PageFault_V2: tmp,
	}
	return
}

// SetDeviceChar sets the value of DeviceChar for the instance
func (instance *PageFault_ImageLoadBacked) SetPropertyDeviceChar(value uint32) (err error) {
	return instance.SetProperty("DeviceChar", (value))
}

// GetDeviceChar gets the value of DeviceChar for the instance
func (instance *PageFault_ImageLoadBacked) GetPropertyDeviceChar() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceChar")
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

// SetFileChar sets the value of FileChar for the instance
func (instance *PageFault_ImageLoadBacked) SetPropertyFileChar(value uint16) (err error) {
	return instance.SetProperty("FileChar", (value))
}

// GetFileChar gets the value of FileChar for the instance
func (instance *PageFault_ImageLoadBacked) GetPropertyFileChar() (value uint16, err error) {
	retValue, err := instance.GetProperty("FileChar")
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

// SetFileObject sets the value of FileObject for the instance
func (instance *PageFault_ImageLoadBacked) SetPropertyFileObject(value uint32) (err error) {
	return instance.SetProperty("FileObject", (value))
}

// GetFileObject gets the value of FileObject for the instance
func (instance *PageFault_ImageLoadBacked) GetPropertyFileObject() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileObject")
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

// SetLoadFlags sets the value of LoadFlags for the instance
func (instance *PageFault_ImageLoadBacked) SetPropertyLoadFlags(value uint16) (err error) {
	return instance.SetProperty("LoadFlags", (value))
}

// GetLoadFlags gets the value of LoadFlags for the instance
func (instance *PageFault_ImageLoadBacked) GetPropertyLoadFlags() (value uint16, err error) {
	retValue, err := instance.GetProperty("LoadFlags")
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
