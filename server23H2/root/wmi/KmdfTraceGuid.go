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

// KmdfTraceGuid struct
type KmdfTraceGuid struct {
	*EventTrace

	// Enable Flags
	Flags KmdfTraceGuid_Flags

	// Levels
	Level KmdfTraceGuid_Level
}

func NewKmdfTraceGuidEx1(instance *cim.WmiInstance) (newInstance *KmdfTraceGuid, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KmdfTraceGuid{
		EventTrace: tmp,
	}
	return
}

func NewKmdfTraceGuidEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KmdfTraceGuid, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KmdfTraceGuid{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *KmdfTraceGuid) SetPropertyFlags(value KmdfTraceGuid_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *KmdfTraceGuid) GetPropertyFlags() (value KmdfTraceGuid_Flags, err error) {
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

	value = KmdfTraceGuid_Flags(valuetmp)

	return
}

// SetLevel sets the value of Level for the instance
func (instance *KmdfTraceGuid) SetPropertyLevel(value KmdfTraceGuid_Level) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *KmdfTraceGuid) GetPropertyLevel() (value KmdfTraceGuid_Level, err error) {
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

	value = KmdfTraceGuid_Level(valuetmp)

	return
}
