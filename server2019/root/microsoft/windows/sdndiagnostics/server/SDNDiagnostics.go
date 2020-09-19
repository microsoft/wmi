// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.SDNDiagnostics.Server
//////////////////////////////////////////////
package server

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SDNDiagnostics struct
type SDNDiagnostics struct {
	*cim.WmiInstance
}

func NewSDNDiagnosticsEx1(instance *cim.WmiInstance) (newInstance *SDNDiagnostics, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &SDNDiagnostics{
		WmiInstance: tmp,
	}
	return
}

func NewSDNDiagnosticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SDNDiagnostics, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SDNDiagnostics{
		WmiInstance: tmp,
	}
	return
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
