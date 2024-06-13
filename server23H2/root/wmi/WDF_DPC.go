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

// WDF_DPC struct
type WDF_DPC struct {
	*PerfInfo_V2

	//
	Routine uint32
}

func NewWDF_DPCEx1(instance *cim.WmiInstance) (newInstance *WDF_DPC, err error) {
	tmp, err := NewPerfInfo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &WDF_DPC{
		PerfInfo_V2: tmp,
	}
	return
}

func NewWDF_DPCEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WDF_DPC, err error) {
	tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WDF_DPC{
		PerfInfo_V2: tmp,
	}
	return
}

// SetRoutine sets the value of Routine for the instance
func (instance *WDF_DPC) SetPropertyRoutine(value uint32) (err error) {
	return instance.SetProperty("Routine", (value))
}

// GetRoutine gets the value of Routine for the instance
func (instance *WDF_DPC) GetPropertyRoutine() (value uint32, err error) {
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
