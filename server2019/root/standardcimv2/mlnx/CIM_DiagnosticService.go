// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_DiagnosticService struct
type CIM_DiagnosticService struct {
	*CIM_Service
}

func NewCIM_DiagnosticServiceEx1(instance *cim.WmiInstance) (newInstance *CIM_DiagnosticService, err error) {
	tmp, err := NewCIM_ServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_DiagnosticService{
		CIM_Service: tmp,
	}
	return
}

func NewCIM_DiagnosticServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_DiagnosticService, err error) {
	tmp, err := NewCIM_ServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_DiagnosticService{
		CIM_Service: tmp,
	}
	return
}

//

// <param name="DiagSetting" type="CIM_DiagnosticSetting "></param>
// <param name="JobSetting" type="CIM_JobSettingData "></param>
// <param name="ManagedElement" type="CIM_ManagedElement "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_DiagnosticService) RunDiagnostic( /* IN */ ManagedElement CIM_ManagedElement,
	/* IN */ DiagSetting CIM_DiagnosticSetting,
	/* IN */ JobSetting CIM_JobSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RunDiagnostic", Action, PercentComplete, Timeout, ManagedElement, DiagSetting, JobSetting)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DiagnosticSettings" type="CIM_DiagnosticSettingData "></param>
// <param name="JobSettings" type="CIM_JobSettingData "></param>
// <param name="ManagedElement" type="CIM_ManagedElement "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_DiagnosticService) RunDiagnosticService( /* IN */ ManagedElement CIM_ManagedElement,
	/* IN */ DiagnosticSettings CIM_DiagnosticSettingData,
	/* IN */ JobSettings CIM_JobSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RunDiagnosticService", Action, PercentComplete, Timeout, ManagedElement, DiagnosticSettings, JobSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
