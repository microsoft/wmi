// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_ShortcutSAP struct
type Win32_ShortcutSAP struct {
	*cim.WmiInstance

	//
	Action Win32_ShortcutAction

	//
	Element Win32_CommandLineAccess
}

func NewWin32_ShortcutSAPEx1(instance *cim.WmiInstance) (newInstance *Win32_ShortcutSAP, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Win32_ShortcutSAP{
		WmiInstance: tmp,
	}
	return
}

func NewWin32_ShortcutSAPEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_ShortcutSAP, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_ShortcutSAP{
		WmiInstance: tmp,
	}
	return
}

// SetAction sets the value of Action for the instance
func (instance *Win32_ShortcutSAP) SetPropertyAction(value Win32_ShortcutAction) (err error) {
	return instance.SetProperty("Action", value)
}

// GetAction gets the value of Action for the instance
func (instance *Win32_ShortcutSAP) GetPropertyAction() (value Win32_ShortcutAction, err error) {
	retValue, err := instance.GetProperty("Action")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_ShortcutAction)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetElement sets the value of Element for the instance
func (instance *Win32_ShortcutSAP) SetPropertyElement(value Win32_CommandLineAccess) (err error) {
	return instance.SetProperty("Element", value)
}

// GetElement gets the value of Element for the instance
func (instance *Win32_ShortcutSAP) GetPropertyElement() (value Win32_CommandLineAccess, err error) {
	retValue, err := instance.GetProperty("Element")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_CommandLineAccess)
	if !ok {
		// TODO: Set an error
	}
	return
}
