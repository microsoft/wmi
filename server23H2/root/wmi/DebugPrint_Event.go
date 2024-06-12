// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DebugPrint_Event struct
type DebugPrint_Event struct {
	*Debugger

	//
	Component uint32

	//
	Level uint32

	//
	Message string
}

func NewDebugPrint_EventEx1(instance *cim.WmiInstance) (newInstance *DebugPrint_Event, err error) {
	tmp, err := NewDebuggerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DebugPrint_Event{
		Debugger: tmp,
	}
	return
}

func NewDebugPrint_EventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DebugPrint_Event, err error) {
	tmp, err := NewDebuggerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DebugPrint_Event{
		Debugger: tmp,
	}
	return
}

// SetComponent sets the value of Component for the instance
func (instance *DebugPrint_Event) SetPropertyComponent(value uint32) (err error) {
	return instance.SetProperty("Component", (value))
}

// GetComponent gets the value of Component for the instance
func (instance *DebugPrint_Event) GetPropertyComponent() (value uint32, err error) {
	retValue, err := instance.GetProperty("Component")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetLevel sets the value of Level for the instance
func (instance *DebugPrint_Event) SetPropertyLevel(value uint32) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *DebugPrint_Event) GetPropertyLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("Level")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetMessage sets the value of Message for the instance
func (instance *DebugPrint_Event) SetPropertyMessage(value string) (err error) {
	return instance.SetProperty("Message", (value))
}

// GetMessage gets the value of Message for the instance
func (instance *DebugPrint_Event) GetPropertyMessage() (value string, err error) {
	retValue, err := instance.GetProperty("Message")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
