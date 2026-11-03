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

// WMPNsSciGuid struct
type WMPNsSciGuid struct {
	*EventTrace

	// Enable Flags
	Flags WMPNsSciGuid_Flags

	// Levels
	Level WMPNsSciGuid_Level
}

func NewWMPNsSciGuidEx1(instance *cim.WmiInstance) (newInstance *WMPNsSciGuid, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WMPNsSciGuid{
		EventTrace: tmp,
	}
	return
}

func NewWMPNsSciGuidEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WMPNsSciGuid, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WMPNsSciGuid{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *WMPNsSciGuid) SetPropertyFlags(value WMPNsSciGuid_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *WMPNsSciGuid) GetPropertyFlags() (value WMPNsSciGuid_Flags, err error) {
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

	value = WMPNsSciGuid_Flags(valuetmp)

	return
}

// SetLevel sets the value of Level for the instance
func (instance *WMPNsSciGuid) SetPropertyLevel(value WMPNsSciGuid_Level) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *WMPNsSciGuid) GetPropertyLevel() (value WMPNsSciGuid_Level, err error) {
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

	value = WMPNsSciGuid_Level(valuetmp)

	return
}
