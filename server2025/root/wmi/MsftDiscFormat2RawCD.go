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

// MsftDiscFormat2RawCD struct
type MsftDiscFormat2RawCD struct {
	*EventTrace

	// Enable Flags
	Flags MsftDiscFormat2RawCD_Flags

	// Levels
	Level MsftDiscFormat2RawCD_Level
}

func NewMsftDiscFormat2RawCDEx1(instance *cim.WmiInstance) (newInstance *MsftDiscFormat2RawCD, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MsftDiscFormat2RawCD{
		EventTrace: tmp,
	}
	return
}

func NewMsftDiscFormat2RawCDEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsftDiscFormat2RawCD, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsftDiscFormat2RawCD{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MsftDiscFormat2RawCD) SetPropertyFlags(value MsftDiscFormat2RawCD_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MsftDiscFormat2RawCD) GetPropertyFlags() (value MsftDiscFormat2RawCD_Flags, err error) {
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

	value = MsftDiscFormat2RawCD_Flags(valuetmp)

	return
}

// SetLevel sets the value of Level for the instance
func (instance *MsftDiscFormat2RawCD) SetPropertyLevel(value MsftDiscFormat2RawCD_Level) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *MsftDiscFormat2RawCD) GetPropertyLevel() (value MsftDiscFormat2RawCD_Level, err error) {
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

	value = MsftDiscFormat2RawCD_Level(valuetmp)

	return
}
