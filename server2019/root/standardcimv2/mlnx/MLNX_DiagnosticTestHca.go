// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_DiagnosticTestHca struct
type MLNX_DiagnosticTestHca struct {
	*MLNX_DiagnosticTest
}

func NewMLNX_DiagnosticTestHcaEx1(instance *cim.WmiInstance) (newInstance *MLNX_DiagnosticTestHca, err error) {
	tmp, err := NewMLNX_DiagnosticTestEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticTestHca{
		MLNX_DiagnosticTest: tmp,
	}
	return
}

func NewMLNX_DiagnosticTestHcaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DiagnosticTestHca, err error) {
	tmp, err := NewMLNX_DiagnosticTestEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticTestHca{
		MLNX_DiagnosticTest: tmp,
	}
	return
}

//

// <param name="DiagnosticSettings" type="CIM_DiagnosticSettingData "></param>
// <param name="JobSettings" type="CIM_JobSettingData "></param>
// <param name="ManagedElement" type="MLNX_NetAdapter "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_DiagnosticTestHca) MlnxDiagnosticService( /* IN */ ManagedElement MLNX_NetAdapter,
	/* IN */ DiagnosticSettings CIM_DiagnosticSettingData,
	/* IN */ JobSettings CIM_JobSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("MlnxDiagnosticService", Action, PercentComplete, Timeout, ManagedElement, DiagnosticSettings, JobSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
