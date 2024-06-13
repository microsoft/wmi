// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DriverMajorFunctionCall struct
type DriverMajorFunctionCall struct {
	*DiskIo_V2

	//
	FileObject uint32

	//
	Irp uint32

	//
	MajorFunction uint32

	//
	MinorFunction uint32

	//
	RoutineAddr uint32

	//
	UniqMatchId uint32
}

func NewDriverMajorFunctionCallEx1(instance *cim.WmiInstance) (newInstance *DriverMajorFunctionCall, err error) {
	tmp, err := NewDiskIo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &DriverMajorFunctionCall{
		DiskIo_V2: tmp,
	}
	return
}

func NewDriverMajorFunctionCallEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DriverMajorFunctionCall, err error) {
	tmp, err := NewDiskIo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DriverMajorFunctionCall{
		DiskIo_V2: tmp,
	}
	return
}

// SetFileObject sets the value of FileObject for the instance
func (instance *DriverMajorFunctionCall) SetPropertyFileObject(value uint32) (err error) {
	return instance.SetProperty("FileObject", (value))
}

// GetFileObject gets the value of FileObject for the instance
func (instance *DriverMajorFunctionCall) GetPropertyFileObject() (value uint32, err error) {
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

// SetIrp sets the value of Irp for the instance
func (instance *DriverMajorFunctionCall) SetPropertyIrp(value uint32) (err error) {
	return instance.SetProperty("Irp", (value))
}

// GetIrp gets the value of Irp for the instance
func (instance *DriverMajorFunctionCall) GetPropertyIrp() (value uint32, err error) {
	retValue, err := instance.GetProperty("Irp")
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
func (instance *DriverMajorFunctionCall) SetPropertyMajorFunction(value uint32) (err error) {
	return instance.SetProperty("MajorFunction", (value))
}

// GetMajorFunction gets the value of MajorFunction for the instance
func (instance *DriverMajorFunctionCall) GetPropertyMajorFunction() (value uint32, err error) {
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

// SetMinorFunction sets the value of MinorFunction for the instance
func (instance *DriverMajorFunctionCall) SetPropertyMinorFunction(value uint32) (err error) {
	return instance.SetProperty("MinorFunction", (value))
}

// GetMinorFunction gets the value of MinorFunction for the instance
func (instance *DriverMajorFunctionCall) GetPropertyMinorFunction() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinorFunction")
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
func (instance *DriverMajorFunctionCall) SetPropertyRoutineAddr(value uint32) (err error) {
	return instance.SetProperty("RoutineAddr", (value))
}

// GetRoutineAddr gets the value of RoutineAddr for the instance
func (instance *DriverMajorFunctionCall) GetPropertyRoutineAddr() (value uint32, err error) {
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

// SetUniqMatchId sets the value of UniqMatchId for the instance
func (instance *DriverMajorFunctionCall) SetPropertyUniqMatchId(value uint32) (err error) {
	return instance.SetProperty("UniqMatchId", (value))
}

// GetUniqMatchId gets the value of UniqMatchId for the instance
func (instance *DriverMajorFunctionCall) GetPropertyUniqMatchId() (value uint32, err error) {
	retValue, err := instance.GetProperty("UniqMatchId")
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
