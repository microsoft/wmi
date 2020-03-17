// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

// Win32_TSExpandEnvironmentVariables struct
type Win32_TSExpandEnvironmentVariables struct {
	CIM_LogicalElement
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
