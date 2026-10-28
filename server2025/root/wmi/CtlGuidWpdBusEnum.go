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

// CtlGuidWpdBusEnum struct
type CtlGuidWpdBusEnum struct {
	*EventTrace

	// Enable Flags
	Flags CtlGuidWpdBusEnum_Flags
}

func NewCtlGuidWpdBusEnumEx1(instance *cim.WmiInstance) (newInstance *CtlGuidWpdBusEnum, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CtlGuidWpdBusEnum{
		EventTrace: tmp,
	}
	return
}

func NewCtlGuidWpdBusEnumEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CtlGuidWpdBusEnum, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CtlGuidWpdBusEnum{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *CtlGuidWpdBusEnum) SetPropertyFlags(value CtlGuidWpdBusEnum_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *CtlGuidWpdBusEnum) GetPropertyFlags() (value CtlGuidWpdBusEnum_Flags, err error) {
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

	value = CtlGuidWpdBusEnum_Flags(valuetmp)

	return
}
