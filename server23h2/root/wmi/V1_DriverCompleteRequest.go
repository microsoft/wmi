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

// V1_DriverCompleteRequest struct
type V1_DriverCompleteRequest struct {
	*DiskIo_V1

	//
	Irp uint32

	//
	RoutineAddr uint32

	//
	UniqMatchId uint32
}

func NewV1_DriverCompleteRequestEx1(instance *cim.WmiInstance) (newInstance *V1_DriverCompleteRequest, err error) {
	tmp, err := NewDiskIo_V1Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &V1_DriverCompleteRequest{
		DiskIo_V1: tmp,
	}
	return
}

func NewV1_DriverCompleteRequestEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *V1_DriverCompleteRequest, err error) {
	tmp, err := NewDiskIo_V1Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &V1_DriverCompleteRequest{
		DiskIo_V1: tmp,
	}
	return
}

// SetIrp sets the value of Irp for the instance
func (instance *V1_DriverCompleteRequest) SetPropertyIrp(value uint32) (err error) {
	return instance.SetProperty("Irp", (value))
}

// GetIrp gets the value of Irp for the instance
func (instance *V1_DriverCompleteRequest) GetPropertyIrp() (value uint32, err error) {
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

// SetRoutineAddr sets the value of RoutineAddr for the instance
func (instance *V1_DriverCompleteRequest) SetPropertyRoutineAddr(value uint32) (err error) {
	return instance.SetProperty("RoutineAddr", (value))
}

// GetRoutineAddr gets the value of RoutineAddr for the instance
func (instance *V1_DriverCompleteRequest) GetPropertyRoutineAddr() (value uint32, err error) {
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
func (instance *V1_DriverCompleteRequest) SetPropertyUniqMatchId(value uint32) (err error) {
	return instance.SetProperty("UniqMatchId", (value))
}

// GetUniqMatchId gets the value of UniqMatchId for the instance
func (instance *V1_DriverCompleteRequest) GetPropertyUniqMatchId() (value uint32, err error) {
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
