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

// WMPNetwkGuid struct
type WMPNetwkGuid struct {
	*EventTrace

	// Enable Flags
	Flags WMPNetwkGuid_Flags

	// Levels
	Level WMPNetwkGuid_Level
}

func NewWMPNetwkGuidEx1(instance *cim.WmiInstance) (newInstance *WMPNetwkGuid, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WMPNetwkGuid{
		EventTrace: tmp,
	}
	return
}

func NewWMPNetwkGuidEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WMPNetwkGuid, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WMPNetwkGuid{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *WMPNetwkGuid) SetPropertyFlags(value WMPNetwkGuid_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *WMPNetwkGuid) GetPropertyFlags() (value WMPNetwkGuid_Flags, err error) {
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

	value = WMPNetwkGuid_Flags(valuetmp)

	return
}

// SetLevel sets the value of Level for the instance
func (instance *WMPNetwkGuid) SetPropertyLevel(value WMPNetwkGuid_Level) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *WMPNetwkGuid) GetPropertyLevel() (value WMPNetwkGuid_Level, err error) {
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

	value = WMPNetwkGuid_Level(valuetmp)

	return
}
