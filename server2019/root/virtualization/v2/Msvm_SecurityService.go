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

// Msvm_SecurityService struct
type Msvm_SecurityService struct {
	CIM_Service
}

//

// <param name="SecuritySettingData" type="string ">An embedded instance of class Msvm_SecuritySettingData that describes modifications to the current security settings of an existing virtual machine.</param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_SecurityService) ModifySecuritySettings( /* IN */ SecuritySettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifySecuritySettings", Action, PercentComplete, Timeout, SecuritySettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="SecurityPolicy" type="uint8 []"></param>
// <param name="SecuritySettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_SecurityService) SetSecurityPolicy( /* IN */ SecuritySettingData string,
	/* IN */ SecurityPolicy []uint8,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("SetSecurityPolicy", Action, PercentComplete, Timeout, SecuritySettingData, SecurityPolicy)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="KeyProtector" type="uint8 []"></param>
// <param name="SecuritySettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_SecurityService) SetKeyProtector( /* IN */ SecuritySettingData string,
	/* IN */ KeyProtector []uint8,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("SetKeyProtector", Action, PercentComplete, Timeout, SecuritySettingData, KeyProtector)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="SecuritySettingData" type="string "></param>

// <param name="KeyProtector" type="uint8 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_SecurityService) GetKeyProtector( /* IN */ SecuritySettingData string,
	/* OUT */ KeyProtector []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetKeyProtector", SecuritySettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="SecuritySettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_SecurityService) RestoreLastKnownGoodKeyProtector( /* IN */ SecuritySettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RestoreLastKnownGoodKeyProtector", Action, PercentComplete, Timeout, SecuritySettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

func (instance *Msvm_SecurityService) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}
