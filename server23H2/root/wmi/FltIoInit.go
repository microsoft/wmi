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

// FltIoInit struct
type FltIoInit struct {
	*FileIo

	//
	CallbackDataPtr uint32

	//
	FileContext uint32

	//
	FileObject uint32

	//
	IrpPtr uint32

	//
	MajorFunction uint32

	//
	RoutineAddr uint32
}

func NewFltIoInitEx1(instance *cim.WmiInstance) (newInstance *FltIoInit, err error) {
	tmp, err := NewFileIoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FltIoInit{
		FileIo: tmp,
	}
	return
}

func NewFltIoInitEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FltIoInit, err error) {
	tmp, err := NewFileIoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FltIoInit{
		FileIo: tmp,
	}
	return
}

// SetCallbackDataPtr sets the value of CallbackDataPtr for the instance
func (instance *FltIoInit) SetPropertyCallbackDataPtr(value uint32) (err error) {
	return instance.SetProperty("CallbackDataPtr", (value))
}

// GetCallbackDataPtr gets the value of CallbackDataPtr for the instance
func (instance *FltIoInit) GetPropertyCallbackDataPtr() (value uint32, err error) {
	retValue, err := instance.GetProperty("CallbackDataPtr")
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

// SetFileContext sets the value of FileContext for the instance
func (instance *FltIoInit) SetPropertyFileContext(value uint32) (err error) {
	return instance.SetProperty("FileContext", (value))
}

// GetFileContext gets the value of FileContext for the instance
func (instance *FltIoInit) GetPropertyFileContext() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileContext")
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

// SetFileObject sets the value of FileObject for the instance
func (instance *FltIoInit) SetPropertyFileObject(value uint32) (err error) {
	return instance.SetProperty("FileObject", (value))
}

// GetFileObject gets the value of FileObject for the instance
func (instance *FltIoInit) GetPropertyFileObject() (value uint32, err error) {
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

// SetIrpPtr sets the value of IrpPtr for the instance
func (instance *FltIoInit) SetPropertyIrpPtr(value uint32) (err error) {
	return instance.SetProperty("IrpPtr", (value))
}

// GetIrpPtr gets the value of IrpPtr for the instance
func (instance *FltIoInit) GetPropertyIrpPtr() (value uint32, err error) {
	retValue, err := instance.GetProperty("IrpPtr")
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

// SetMajorFunction sets the value of MajorFunction for the instance
func (instance *FltIoInit) SetPropertyMajorFunction(value uint32) (err error) {
	return instance.SetProperty("MajorFunction", (value))
}

// GetMajorFunction gets the value of MajorFunction for the instance
func (instance *FltIoInit) GetPropertyMajorFunction() (value uint32, err error) {
	retValue, err := instance.GetProperty("MajorFunction")
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

// SetRoutineAddr sets the value of RoutineAddr for the instance
func (instance *FltIoInit) SetPropertyRoutineAddr(value uint32) (err error) {
	return instance.SetProperty("RoutineAddr", (value))
}

// GetRoutineAddr gets the value of RoutineAddr for the instance
func (instance *FltIoInit) GetPropertyRoutineAddr() (value uint32, err error) {
	retValue, err := instance.GetProperty("RoutineAddr")
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
