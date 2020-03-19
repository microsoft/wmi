// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_TerminalTerminalSetting struct
type Win32_TerminalTerminalSetting struct {
	*CIM_ElementSetting

	//
	Element Win32_Terminal

	//
	Setting Win32_TerminalSetting
}

func NewWin32_TerminalTerminalSettingEx1(instance *cim.WmiInstance) (newInstance *Win32_TerminalTerminalSetting, err error) {
	tmp, err := NewCIM_ElementSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TerminalTerminalSetting{
		CIM_ElementSetting: tmp,
	}
	return
}

func NewWin32_TerminalTerminalSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TerminalTerminalSetting, err error) {
	tmp, err := NewCIM_ElementSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TerminalTerminalSetting{
		CIM_ElementSetting: tmp,
	}
	return
}

// SetElement sets the value of Element for the instance
func (instance *Win32_TerminalTerminalSetting) SetPropertyElement(value Win32_Terminal) (err error) {
	return instance.SetProperty("Element", value)
}

// GetElement gets the value of Element for the instance
func (instance *Win32_TerminalTerminalSetting) GetPropertyElement() (value Win32_Terminal, err error) {
	retValue, err := instance.GetProperty("Element")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_Terminal)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSetting sets the value of Setting for the instance
func (instance *Win32_TerminalTerminalSetting) SetPropertySetting(value Win32_TerminalSetting) (err error) {
	return instance.SetProperty("Setting", value)
}

// GetSetting gets the value of Setting for the instance
func (instance *Win32_TerminalTerminalSetting) GetPropertySetting() (value Win32_TerminalSetting, err error) {
	retValue, err := instance.GetProperty("Setting")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_TerminalSetting)
	if !ok {
		// TODO: Set an error
	}
	return
}
