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
	cimv2 "github.com/microsoft/wmi/server2019/root/cimv2"
)

// Win32_TSVirtualIPSetting struct
type Win32_TSVirtualIPSetting struct {
	*CIM_ElementSetting

	//
	Element cimv2.Win32_TerminalService

	//
	Setting Win32_TSVirtualIP
}

func NewWin32_TSVirtualIPSettingEx1(instance *cim.WmiInstance) (newInstance *Win32_TSVirtualIPSetting, err error) {
	tmp, err := NewCIM_ElementSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TSVirtualIPSetting{
		CIM_ElementSetting: tmp,
	}
	return
}

func NewWin32_TSVirtualIPSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSVirtualIPSetting, err error) {
	tmp, err := NewCIM_ElementSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSVirtualIPSetting{
		CIM_ElementSetting: tmp,
	}
	return
}

// SetElement sets the value of Element for the instance
func (instance *Win32_TSVirtualIPSetting) SetPropertyElement(value cimv2.Win32_TerminalService) (err error) {
	return instance.SetProperty("Element", value)
}

// GetElement gets the value of Element for the instance
func (instance *Win32_TSVirtualIPSetting) GetPropertyElement() (value cimv2.Win32_TerminalService, err error) {
	retValue, err := instance.GetProperty("Element")
	if err != nil {
		return
	}
	value, ok := retValue.(cimv2.Win32_TerminalService)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSetting sets the value of Setting for the instance
func (instance *Win32_TSVirtualIPSetting) SetPropertySetting(value Win32_TSVirtualIP) (err error) {
	return instance.SetProperty("Setting", value)
}

// GetSetting gets the value of Setting for the instance
func (instance *Win32_TSVirtualIPSetting) GetPropertySetting() (value Win32_TSVirtualIP, err error) {
	retValue, err := instance.GetProperty("Setting")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_TSVirtualIP)
	if !ok {
		// TODO: Set an error
	}
	return
}
