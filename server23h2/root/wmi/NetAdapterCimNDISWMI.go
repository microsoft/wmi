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

// NetAdapterCimNDISWMI struct
type NetAdapterCimNDISWMI struct {
	*EventTrace

	//
	Flags NetAdapterCimNDISWMI_Flags

	//
	Level NetAdapterCimNDISWMI_Level
}

func NewNetAdapterCimNDISWMIEx1(instance *cim.WmiInstance) (newInstance *NetAdapterCimNDISWMI, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NetAdapterCimNDISWMI{
		EventTrace: tmp,
	}
	return
}

func NewNetAdapterCimNDISWMIEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NetAdapterCimNDISWMI, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NetAdapterCimNDISWMI{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *NetAdapterCimNDISWMI) SetPropertyFlags(value NetAdapterCimNDISWMI_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *NetAdapterCimNDISWMI) GetPropertyFlags() (value NetAdapterCimNDISWMI_Flags, err error) {
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

	value = NetAdapterCimNDISWMI_Flags(valuetmp)

	return
}

// SetLevel sets the value of Level for the instance
func (instance *NetAdapterCimNDISWMI) SetPropertyLevel(value NetAdapterCimNDISWMI_Level) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *NetAdapterCimNDISWMI) GetPropertyLevel() (value NetAdapterCimNDISWMI_Level, err error) {
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

	value = NetAdapterCimNDISWMI_Level(valuetmp)

	return
}
