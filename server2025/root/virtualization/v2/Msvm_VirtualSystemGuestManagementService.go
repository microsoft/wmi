// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VirtualSystemGuestManagementService struct
type Msvm_VirtualSystemGuestManagementService struct {
	*CIM_Service
}

func NewMsvm_VirtualSystemGuestManagementServiceEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemGuestManagementService, err error) {
	tmp, err := NewCIM_ServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemGuestManagementService{
		CIM_Service: tmp,
	}
	return
}

func NewMsvm_VirtualSystemGuestManagementServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemGuestManagementService, err error) {
	tmp, err := NewCIM_ServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemGuestManagementService{
		CIM_Service: tmp,
	}
	return
}

//

// <param name="VmId" type="string "></param>

// <param name="CurrentUpdateId" type="uint64 ">The current updated ID.</param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="Settings" type="uint8 []">String represents VTL2 settings in JSON format.</param>
func (instance *Msvm_VirtualSystemGuestManagementService) GetVtl2Settings( /* IN */ VmId string,
	/* OUT */ Settings []uint8,
	/* OUT */ CurrentUpdateId uint64) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetVtl2Settings", VmId)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CurrentUpdateId" type="uint64 "></param>
// <param name="Settings" type="uint8 []"></param>
// <param name="VmId" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob ">May contain a reference to the ConcreteJob created to track the state transition initiated by the method invocation.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemGuestManagementService) ModifyVtl2Settings( /* IN */ VmId string,
	/* IN */ Settings []uint8,
	/* IN */ CurrentUpdateId uint64,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifyVtl2Settings", Action, PercentComplete, Timeout, VmId, Settings, CurrentUpdateId)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="VmId" type="string "></param>

// <param name="CurrentUpdateId" type="uint64 ">The current updated ID.</param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="Settings" type="string ">String represents VTL2 settings in JSON format.</param>
func (instance *Msvm_VirtualSystemGuestManagementService) QueryVtl2Settings( /* IN */ VmId string,
	/* OUT */ Settings string,
	/* OUT */ CurrentUpdateId uint64) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("QueryVtl2Settings", VmId)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CurrentUpdateId" type="uint64 "></param>
// <param name="Settings" type="string "></param>
// <param name="VmId" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob ">May contain a reference to the ConcreteJob created to track the state transition initiated by the method invocation.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemGuestManagementService) UpdateVtl2Settings( /* IN */ VmId string,
	/* IN */ Settings string,
	/* IN */ CurrentUpdateId uint64,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("UpdateVtl2Settings", Action, PercentComplete, Timeout, VmId, Settings, CurrentUpdateId)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Namespace" type="string "></param>
// <param name="VmId" type="string "></param>

// <param name="CurrentUpdateId" type="uint64 ">The current updated ID.</param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="Settings" type="uint8 []">Bytes representing the settings.</param>
func (instance *Msvm_VirtualSystemGuestManagementService) GetManagementVtlSettings( /* IN */ VmId string,
	/* IN */ Namespace string,
	/* OUT */ Settings []uint8,
	/* OUT */ CurrentUpdateId uint64) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetManagementVtlSettings", VmId, Namespace)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CurrentUpdateId" type="uint64 "></param>
// <param name="Namespace" type="string "></param>
// <param name="Settings" type="uint8 []"></param>
// <param name="VmId" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob ">May contain a reference to the ConcreteJob created to track the state transition initiated by the method invocation.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemGuestManagementService) SetManagementVtlSettings( /* IN */ VmId string,
	/* IN */ Namespace string,
	/* IN */ Settings []uint8,
	/* IN */ CurrentUpdateId uint64,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("SetManagementVtlSettings", Action, PercentComplete, Timeout, VmId, Namespace, Settings, CurrentUpdateId)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Namespace" type="string "></param>
// <param name="VmId" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob ">May contain a reference to the ConcreteJob created to track the state transition initiated by the method invocation.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemGuestManagementService) RemoveManagementVtlSettings( /* IN */ VmId string,
	/* IN */ Namespace string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RemoveManagementVtlSettings", Action, PercentComplete, Timeout, VmId, Namespace)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Options" type="uint32 "></param>
// <param name="TimeoutHintSecs" type="uint16 "></param>
// <param name="VmId" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemGuestManagementService) ReloadManagementVtl( /* IN */ VmId string,
	/* IN */ Options uint32,
	/* IN */ TimeoutHintSecs uint16,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ReloadManagementVtl", Action, PercentComplete, Timeout, VmId, Options, TimeoutHintSecs)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
