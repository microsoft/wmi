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

// MsftWriteEngine2 struct
type MsftWriteEngine2 struct {
	*EventTrace

	// Enable Flags
	Flags MsftWriteEngine2_Flags

	// Levels
	Level MsftWriteEngine2_Level
}

func NewMsftWriteEngine2Ex1(instance *cim.WmiInstance) (newInstance *MsftWriteEngine2, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MsftWriteEngine2{
		EventTrace: tmp,
	}
	return
}

func NewMsftWriteEngine2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsftWriteEngine2, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsftWriteEngine2{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MsftWriteEngine2) SetPropertyFlags(value MsftWriteEngine2_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MsftWriteEngine2) GetPropertyFlags() (value MsftWriteEngine2_Flags, err error) {
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

	value = MsftWriteEngine2_Flags(valuetmp)

	return
}

// SetLevel sets the value of Level for the instance
func (instance *MsftWriteEngine2) SetPropertyLevel(value MsftWriteEngine2_Level) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *MsftWriteEngine2) GetPropertyLevel() (value MsftWriteEngine2_Level, err error) {
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

	value = MsftWriteEngine2_Level(valuetmp)

	return
}
