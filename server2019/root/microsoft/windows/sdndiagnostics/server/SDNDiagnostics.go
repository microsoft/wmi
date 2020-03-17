// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.SDNDiagnostics.Server
//////////////////////////////////////////////
package server

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// SDNDiagnostics struct
type SDNDiagnostics struct {
	cim.WmiInstance
}

//

// <param name="DiagnosticInformation" type="DiagnosticInfo "></param>

// <param name="ChangeInErrorCode" type="bool "></param>
// <param name="ErrorCode" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDNDiagnostics) Enable( /* IN */ DiagnosticInformation DiagnosticInfo,
	/* OUT */ ErrorCode uint32,
	/* OUT */ ChangeInErrorCode bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable", DiagnosticInformation)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDNDiagnostics) Disable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDNDiagnostics) UpdateTraceProvidersList() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("UpdateTraceProvidersList")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
