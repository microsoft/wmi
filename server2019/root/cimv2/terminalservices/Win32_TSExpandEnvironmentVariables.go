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

// Win32_TSExpandEnvironmentVariables struct
type Win32_TSExpandEnvironmentVariables struct {
	*CIM_LogicalElement
}

func NewWin32_TSExpandEnvironmentVariablesEx1(instance *cim.WmiInstance) (newInstance *Win32_TSExpandEnvironmentVariables, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TSExpandEnvironmentVariables{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewWin32_TSExpandEnvironmentVariablesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSExpandEnvironmentVariables, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSExpandEnvironmentVariables{
		CIM_LogicalElement: tmp,
	}
	return
}

// Expands System Defined Environment Variables

// <param name="OriginalString" type="string ">String that contains the environment variables to expand</param>

// <param name="ExpandedString" type="string ">String with the environment variables expanded</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSExpandEnvironmentVariables) EnvironmentVariables( /* IN */ OriginalString string,
	/* OUT */ ExpandedString string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EnvironmentVariables", OriginalString)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
