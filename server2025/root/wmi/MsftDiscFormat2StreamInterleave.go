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

// MsftDiscFormat2StreamInterleave struct
type MsftDiscFormat2StreamInterleave struct {
	*EventTrace

	// Enable Flags
	Flags MsftDiscFormat2StreamInterleave_Flags

	// Levels
	Level MsftDiscFormat2StreamInterleave_Level
}

func NewMsftDiscFormat2StreamInterleaveEx1(instance *cim.WmiInstance) (newInstance *MsftDiscFormat2StreamInterleave, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MsftDiscFormat2StreamInterleave{
		EventTrace: tmp,
	}
	return
}

func NewMsftDiscFormat2StreamInterleaveEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsftDiscFormat2StreamInterleave, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsftDiscFormat2StreamInterleave{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MsftDiscFormat2StreamInterleave) SetPropertyFlags(value MsftDiscFormat2StreamInterleave_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MsftDiscFormat2StreamInterleave) GetPropertyFlags() (value MsftDiscFormat2StreamInterleave_Flags, err error) {
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

	value = MsftDiscFormat2StreamInterleave_Flags(valuetmp)

	return
}

// SetLevel sets the value of Level for the instance
func (instance *MsftDiscFormat2StreamInterleave) SetPropertyLevel(value MsftDiscFormat2StreamInterleave_Level) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *MsftDiscFormat2StreamInterleave) GetPropertyLevel() (value MsftDiscFormat2StreamInterleave_Level, err error) {
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

	value = MsftDiscFormat2StreamInterleave_Level(valuetmp)

	return
}
