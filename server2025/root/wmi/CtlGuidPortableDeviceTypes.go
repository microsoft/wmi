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

// CtlGuidPortableDeviceTypes struct
type CtlGuidPortableDeviceTypes struct {
	*EventTrace

	// Enable Flags
	Flags CtlGuidPortableDeviceTypes_Flags
}

func NewCtlGuidPortableDeviceTypesEx1(instance *cim.WmiInstance) (newInstance *CtlGuidPortableDeviceTypes, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CtlGuidPortableDeviceTypes{
		EventTrace: tmp,
	}
	return
}

func NewCtlGuidPortableDeviceTypesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CtlGuidPortableDeviceTypes, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CtlGuidPortableDeviceTypes{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *CtlGuidPortableDeviceTypes) SetPropertyFlags(value CtlGuidPortableDeviceTypes_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *CtlGuidPortableDeviceTypes) GetPropertyFlags() (value CtlGuidPortableDeviceTypes_Flags, err error) {
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

	value = CtlGuidPortableDeviceTypes_Flags(valuetmp)

	return
}
