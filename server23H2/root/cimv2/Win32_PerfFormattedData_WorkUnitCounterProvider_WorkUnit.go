// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_WorkUnitCounterProvider_WorkUnit struct
type Win32_PerfFormattedData_WorkUnitCounterProvider_WorkUnit struct {
	*Win32_PerfFormattedData

	//
	AppOwnerProcessID uint32

	//
	HostProcessID uint32
}

func NewWin32_PerfFormattedData_WorkUnitCounterProvider_WorkUnitEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_WorkUnitCounterProvider_WorkUnit, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_WorkUnitCounterProvider_WorkUnit{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_WorkUnitCounterProvider_WorkUnitEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_WorkUnitCounterProvider_WorkUnit, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_WorkUnitCounterProvider_WorkUnit{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetAppOwnerProcessID sets the value of AppOwnerProcessID for the instance
func (instance *Win32_PerfFormattedData_WorkUnitCounterProvider_WorkUnit) SetPropertyAppOwnerProcessID(value uint32) (err error) {
	return instance.SetProperty("AppOwnerProcessID", (value))
}

// GetAppOwnerProcessID gets the value of AppOwnerProcessID for the instance
func (instance *Win32_PerfFormattedData_WorkUnitCounterProvider_WorkUnit) GetPropertyAppOwnerProcessID() (value uint32, err error) {
	retValue, err := instance.GetProperty("AppOwnerProcessID")
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

// SetHostProcessID sets the value of HostProcessID for the instance
func (instance *Win32_PerfFormattedData_WorkUnitCounterProvider_WorkUnit) SetPropertyHostProcessID(value uint32) (err error) {
	return instance.SetProperty("HostProcessID", (value))
}

// GetHostProcessID gets the value of HostProcessID for the instance
func (instance *Win32_PerfFormattedData_WorkUnitCounterProvider_WorkUnit) GetPropertyHostProcessID() (value uint32, err error) {
	retValue, err := instance.GetProperty("HostProcessID")
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
