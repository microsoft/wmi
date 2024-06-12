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

// FltIoFailure struct
type FltIoFailure struct {
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

	//
	Status uint32
}

func NewFltIoFailureEx1(instance *cim.WmiInstance) (newInstance *FltIoFailure, err error) {
	tmp, err := NewFileIoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FltIoFailure{
		FileIo: tmp,
	}
	return
}

func NewFltIoFailureEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FltIoFailure, err error) {
	tmp, err := NewFileIoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FltIoFailure{
		FileIo: tmp,
	}
	return
}

// SetCallbackDataPtr sets the value of CallbackDataPtr for the instance
func (instance *FltIoFailure) SetPropertyCallbackDataPtr(value uint32) (err error) {
	return instance.SetProperty("CallbackDataPtr", (value))
}

// GetCallbackDataPtr gets the value of CallbackDataPtr for the instance
func (instance *FltIoFailure) GetPropertyCallbackDataPtr() (value uint32, err error) {
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
func (instance *FltIoFailure) SetPropertyFileContext(value uint32) (err error) {
	return instance.SetProperty("FileContext", (value))
}

// GetFileContext gets the value of FileContext for the instance
func (instance *FltIoFailure) GetPropertyFileContext() (value uint32, err error) {
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
func (instance *FltIoFailure) SetPropertyFileObject(value uint32) (err error) {
	return instance.SetProperty("FileObject", (value))
}

// GetFileObject gets the value of FileObject for the instance
func (instance *FltIoFailure) GetPropertyFileObject() (value uint32, err error) {
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
func (instance *FltIoFailure) SetPropertyIrpPtr(value uint32) (err error) {
	return instance.SetProperty("IrpPtr", (value))
}

// GetIrpPtr gets the value of IrpPtr for the instance
func (instance *FltIoFailure) GetPropertyIrpPtr() (value uint32, err error) {
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
func (instance *FltIoFailure) SetPropertyMajorFunction(value uint32) (err error) {
	return instance.SetProperty("MajorFunction", (value))
}

// GetMajorFunction gets the value of MajorFunction for the instance
func (instance *FltIoFailure) GetPropertyMajorFunction() (value uint32, err error) {
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
func (instance *FltIoFailure) SetPropertyRoutineAddr(value uint32) (err error) {
	return instance.SetProperty("RoutineAddr", (value))
}

// GetRoutineAddr gets the value of RoutineAddr for the instance
func (instance *FltIoFailure) GetPropertyRoutineAddr() (value uint32, err error) {
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

// SetStatus sets the value of Status for the instance
func (instance *FltIoFailure) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *FltIoFailure) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
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
