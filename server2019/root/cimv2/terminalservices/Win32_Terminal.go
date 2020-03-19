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

// Win32_Terminal struct
type Win32_Terminal struct {
	*CIM_LogicalElement

	//
	fEnableTerminal uint32

	//
	LoggedOnUsers uint32

	//
	TerminalName string
}

func NewWin32_TerminalEx1(instance *cim.WmiInstance) (newInstance *Win32_Terminal, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_Terminal{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewWin32_TerminalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_Terminal, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_Terminal{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetfEnableTerminal sets the value of fEnableTerminal for the instance
func (instance *Win32_Terminal) SetPropertyfEnableTerminal(value uint32) (err error) {
	return instance.SetProperty("fEnableTerminal", value)
}

// GetfEnableTerminal gets the value of fEnableTerminal for the instance
func (instance *Win32_Terminal) GetPropertyfEnableTerminal() (value uint32, err error) {
	retValue, err := instance.GetProperty("fEnableTerminal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLoggedOnUsers sets the value of LoggedOnUsers for the instance
func (instance *Win32_Terminal) SetPropertyLoggedOnUsers(value uint32) (err error) {
	return instance.SetProperty("LoggedOnUsers", value)
}

// GetLoggedOnUsers gets the value of LoggedOnUsers for the instance
func (instance *Win32_Terminal) GetPropertyLoggedOnUsers() (value uint32, err error) {
	retValue, err := instance.GetProperty("LoggedOnUsers")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTerminalName sets the value of TerminalName for the instance
func (instance *Win32_Terminal) SetPropertyTerminalName(value string) (err error) {
	return instance.SetProperty("TerminalName", value)
}

// GetTerminalName gets the value of TerminalName for the instance
func (instance *Win32_Terminal) GetPropertyTerminalName() (value string, err error) {
	retValue, err := instance.GetProperty("TerminalName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="fEnableTerminal" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Terminal) Enable( /* IN */ fEnableTerminal uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable", fEnableTerminal)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NewTerminalName" type="string "></param>
// <param name="TerminalProtocol" type="string "></param>
// <param name="Transport" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Terminal) Create( /* IN */ NewTerminalName string,
	/* IN */ Transport string,
	/* IN */ TerminalProtocol string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Create", NewTerminalName, Transport, TerminalProtocol)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NewTerminalName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Terminal) Rename( /* IN */ NewTerminalName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Rename", NewTerminalName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NewTerminalName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Terminal) Delete( /* IN */ NewTerminalName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Delete", NewTerminalName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
