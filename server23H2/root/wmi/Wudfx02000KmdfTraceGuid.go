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

// Wudfx02000KmdfTraceGuid struct
type Wudfx02000KmdfTraceGuid struct {
	*EventTrace

	//
	Flags Wudfx02000KmdfTraceGuid_Flags

	//
	Level Wudfx02000KmdfTraceGuid_Level
}

func NewWudfx02000KmdfTraceGuidEx1(instance *cim.WmiInstance) (newInstance *Wudfx02000KmdfTraceGuid, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Wudfx02000KmdfTraceGuid{
		EventTrace: tmp,
	}
	return
}

func NewWudfx02000KmdfTraceGuidEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Wudfx02000KmdfTraceGuid, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Wudfx02000KmdfTraceGuid{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *Wudfx02000KmdfTraceGuid) SetPropertyFlags(value Wudfx02000KmdfTraceGuid_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *Wudfx02000KmdfTraceGuid) GetPropertyFlags() (value Wudfx02000KmdfTraceGuid_Flags, err error) {
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

	value = Wudfx02000KmdfTraceGuid_Flags(valuetmp)

	return
}

// SetLevel sets the value of Level for the instance
func (instance *Wudfx02000KmdfTraceGuid) SetPropertyLevel(value Wudfx02000KmdfTraceGuid_Level) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *Wudfx02000KmdfTraceGuid) GetPropertyLevel() (value Wudfx02000KmdfTraceGuid_Level, err error) {
	retValue, err := instance.GetProperty("Level")
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

	value = Wudfx02000KmdfTraceGuid_Level(valuetmp)

	return
}
