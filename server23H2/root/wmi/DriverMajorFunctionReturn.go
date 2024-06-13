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

// DriverMajorFunctionReturn struct
type DriverMajorFunctionReturn struct {
	*DiskIo_V2

	//
	Irp uint32

	//
	UniqMatchId uint32
}

func NewDriverMajorFunctionReturnEx1(instance *cim.WmiInstance) (newInstance *DriverMajorFunctionReturn, err error) {
	tmp, err := NewDiskIo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &DriverMajorFunctionReturn{
		DiskIo_V2: tmp,
	}
	return
}

func NewDriverMajorFunctionReturnEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DriverMajorFunctionReturn, err error) {
	tmp, err := NewDiskIo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DriverMajorFunctionReturn{
		DiskIo_V2: tmp,
	}
	return
}

// SetIrp sets the value of Irp for the instance
func (instance *DriverMajorFunctionReturn) SetPropertyIrp(value uint32) (err error) {
	return instance.SetProperty("Irp", (value))
}

// GetIrp gets the value of Irp for the instance
func (instance *DriverMajorFunctionReturn) GetPropertyIrp() (value uint32, err error) {
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

// SetUniqMatchId sets the value of UniqMatchId for the instance
func (instance *DriverMajorFunctionReturn) SetPropertyUniqMatchId(value uint32) (err error) {
	return instance.SetProperty("UniqMatchId", (value))
}

// GetUniqMatchId gets the value of UniqMatchId for the instance
func (instance *DriverMajorFunctionReturn) GetPropertyUniqMatchId() (value uint32, err error) {
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
