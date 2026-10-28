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

// MsftDiscFormat2StreamZero struct
type MsftDiscFormat2StreamZero struct {
	*EventTrace

	// Enable Flags
	Flags MsftDiscFormat2StreamZero_Flags

	// Levels
	Level MsftDiscFormat2StreamZero_Level
}

func NewMsftDiscFormat2StreamZeroEx1(instance *cim.WmiInstance) (newInstance *MsftDiscFormat2StreamZero, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MsftDiscFormat2StreamZero{
		EventTrace: tmp,
	}
	return
}

func NewMsftDiscFormat2StreamZeroEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsftDiscFormat2StreamZero, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsftDiscFormat2StreamZero{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MsftDiscFormat2StreamZero) SetPropertyFlags(value MsftDiscFormat2StreamZero_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MsftDiscFormat2StreamZero) GetPropertyFlags() (value MsftDiscFormat2StreamZero_Flags, err error) {
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

	value = MsftDiscFormat2StreamZero_Flags(valuetmp)

	return
}

// SetLevel sets the value of Level for the instance
func (instance *MsftDiscFormat2StreamZero) SetPropertyLevel(value MsftDiscFormat2StreamZero_Level) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *MsftDiscFormat2StreamZero) GetPropertyLevel() (value MsftDiscFormat2StreamZero_Level, err error) {
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

	value = MsftDiscFormat2StreamZero_Level(valuetmp)

	return
}
