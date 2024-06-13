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

// DriverCompletionRoutine struct
type DriverCompletionRoutine struct {
	*DiskIo_V2

	//
	IrpPtr uint32

	//
	Routine uint32

	//
	UniqMatchId uint32
}

func NewDriverCompletionRoutineEx1(instance *cim.WmiInstance) (newInstance *DriverCompletionRoutine, err error) {
	tmp, err := NewDiskIo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &DriverCompletionRoutine{
		DiskIo_V2: tmp,
	}
	return
}

func NewDriverCompletionRoutineEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DriverCompletionRoutine, err error) {
	tmp, err := NewDiskIo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DriverCompletionRoutine{
		DiskIo_V2: tmp,
	}
	return
}

// SetIrpPtr sets the value of IrpPtr for the instance
func (instance *DriverCompletionRoutine) SetPropertyIrpPtr(value uint32) (err error) {
	return instance.SetProperty("IrpPtr", (value))
}

// GetIrpPtr gets the value of IrpPtr for the instance
func (instance *DriverCompletionRoutine) GetPropertyIrpPtr() (value uint32, err error) {
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

// SetRoutine sets the value of Routine for the instance
func (instance *DriverCompletionRoutine) SetPropertyRoutine(value uint32) (err error) {
	return instance.SetProperty("Routine", (value))
}

// GetRoutine gets the value of Routine for the instance
func (instance *DriverCompletionRoutine) GetPropertyRoutine() (value uint32, err error) {
	retValue, err := instance.GetProperty("Routine")
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
func (instance *DriverCompletionRoutine) SetPropertyUniqMatchId(value uint32) (err error) {
	return instance.SetProperty("UniqMatchId", (value))
}

// GetUniqMatchId gets the value of UniqMatchId for the instance
func (instance *DriverCompletionRoutine) GetPropertyUniqMatchId() (value uint32, err error) {
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
