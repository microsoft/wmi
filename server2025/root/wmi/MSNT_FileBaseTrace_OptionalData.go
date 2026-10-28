// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNT_FileBaseTrace_OptionalData struct
type MSNT_FileBaseTrace_OptionalData struct {
	*EventTrace

	//
	Flags OptionalData_Flags
}

func NewMSNT_FileBaseTrace_OptionalDataEx1(instance *cim.WmiInstance) (newInstance *MSNT_FileBaseTrace_OptionalData, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNT_FileBaseTrace_OptionalData{
		EventTrace: tmp,
	}
	return
}

func NewMSNT_FileBaseTrace_OptionalDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNT_FileBaseTrace_OptionalData, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNT_FileBaseTrace_OptionalData{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MSNT_FileBaseTrace_OptionalData) SetPropertyFlags(value OptionalData_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSNT_FileBaseTrace_OptionalData) GetPropertyFlags() (value OptionalData_Flags, err error) {
	retValue, err := instance.GetProperty("Flags")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = OptionalData_Flags(valuetmp)

	return
}
