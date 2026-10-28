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

// MsftDiscFormat2MsfAddress struct
type MsftDiscFormat2MsfAddress struct {
	*EventTrace

	// Enable Flags
	Flags MsftDiscFormat2MsfAddress_Flags

	// Levels
	Level MsftDiscFormat2MsfAddress_Level
}

func NewMsftDiscFormat2MsfAddressEx1(instance *cim.WmiInstance) (newInstance *MsftDiscFormat2MsfAddress, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MsftDiscFormat2MsfAddress{
		EventTrace: tmp,
	}
	return
}

func NewMsftDiscFormat2MsfAddressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsftDiscFormat2MsfAddress, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsftDiscFormat2MsfAddress{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MsftDiscFormat2MsfAddress) SetPropertyFlags(value MsftDiscFormat2MsfAddress_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MsftDiscFormat2MsfAddress) GetPropertyFlags() (value MsftDiscFormat2MsfAddress_Flags, err error) {
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

	value = MsftDiscFormat2MsfAddress_Flags(valuetmp)

	return
}

// SetLevel sets the value of Level for the instance
func (instance *MsftDiscFormat2MsfAddress) SetPropertyLevel(value MsftDiscFormat2MsfAddress_Level) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *MsftDiscFormat2MsfAddress) GetPropertyLevel() (value MsftDiscFormat2MsfAddress_Level, err error) {
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

	value = MsftDiscFormat2MsfAddress_Level(valuetmp)

	return
}
