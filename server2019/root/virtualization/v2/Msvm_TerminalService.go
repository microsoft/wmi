// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_TerminalService struct
type Msvm_TerminalService struct {
	CIM_Service
}

//

// <param name="ServiceSettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_TerminalService) ModifyServiceSettings( /* IN */ ServiceSettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifyServiceSettings", Action, PercentComplete, Timeout, ServiceSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="Trustees" type="string []"></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_TerminalService) GrantInteractiveSessionAccess( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ Trustees []string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("GrantInteractiveSessionAccess", Action, PercentComplete, Timeout, ComputerSystem, Trustees)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="Trustees" type="string []"></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_TerminalService) RevokeInteractiveSessionAccess( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ Trustees []string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RevokeInteractiveSessionAccess", Action, PercentComplete, Timeout, ComputerSystem, Trustees)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>

// <param name="AccessControlList" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_TerminalService) GetInteractiveSessionACL( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* OUT */ AccessControlList []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetInteractiveSessionACL", ComputerSystem)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

func (instance *Msvm_TerminalService) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_TerminalService) GetRelatedTerminalServiceSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_TerminalServiceSettingData")
}
