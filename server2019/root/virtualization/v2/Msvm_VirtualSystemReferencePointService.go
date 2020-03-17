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

// Msvm_VirtualSystemReferencePointService struct
type Msvm_VirtualSystemReferencePointService struct {
	CIM_Service
}

// Creates a reference point of a virtual system.

// <param name="AffectedSystem" type="Msvm_ComputerSystem ">Reference to the affected virtual system.</param>
// <param name="ReferencePointSettings" type="string ">Parameter settings.</param>
// <param name="ReferencePointType" type="VirtualSystemReferencePointService_ReferencePointType ">Requested reference point type:
///Log based: Based on Hyper-V replica log tracking.
///RCT based: Based on Resilient Change Tracking of virtual disks.</param>
// <param name="ResultingReferencePoint" type="Msvm_VirtualSystemReferencePoint ">Resulting virtual system reference point</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned. In this case, the instance of the Msvm_VirtualSystemReferencePoint class representing the new virtual system reference point is presented via the CIM_AffectedJobElement association with the value of the AffectedElement property referring to the new instance of the Msvm_VirtualSystemReferencePoint class representing the virtual system reference point and the value of the ElementEffects set to 5 (Create).</param>
// <param name="ResultingReferencePoint" type="Msvm_VirtualSystemReferencePoint ">Resulting virtual system reference point</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemReferencePointService) CreateReferencePoint( /* IN */ AffectedSystem Msvm_ComputerSystem,
	/* IN */ ReferencePointSettings string,
	/* IN */ ReferencePointType VirtualSystemReferencePointService_ReferencePointType,
	/* IN/OUT */ ResultingReferencePoint Msvm_VirtualSystemReferencePoint,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("CreateReferencePoint", Action, PercentComplete, Timeout, AffectedSystem, ReferencePointSettings, ReferencePointType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ExportDirectory" type="string "></param>
// <param name="ExportSettingData" type="string "></param>
// <param name="ReferencePoint" type="Msvm_VirtualSystemReferencePoint "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemReferencePointService) ExportReferencePoint( /* IN */ ReferencePoint Msvm_VirtualSystemReferencePoint,
	/* IN */ ExportDirectory string,
	/* IN */ ExportSettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ExportReferencePoint", Action, PercentComplete, Timeout, ReferencePoint, ExportDirectory, ExportSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AffectedReferencePoint" type="Msvm_VirtualSystemReferencePoint "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemReferencePointService) DestroyReferencePoint( /* IN */ AffectedReferencePoint Msvm_VirtualSystemReferencePoint,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("DestroyReferencePoint", Action, PercentComplete, Timeout, AffectedReferencePoint)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AffectedReferencePoint" type="Msvm_VirtualSystemReferencePoint "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemReferencePointService) RemoveAssociatedData( /* IN */ AffectedReferencePoint Msvm_VirtualSystemReferencePoint,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RemoveAssociatedData", Action, PercentComplete, Timeout, AffectedReferencePoint)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AffectedSystem" type="Msvm_ComputerSystem ">Reference to the affected virtual system.</param>
// <param name="ConfigFilePath" type="string "></param>
// <param name="RuntimeStateFilePath" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemReferencePointService) ImportReferencePointMetadata( /* IN */ AffectedSystem Msvm_ComputerSystem,
	/* IN */ ConfigFilePath string,
	/* IN */ RuntimeStateFilePath string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ImportReferencePointMetadata", Action, PercentComplete, Timeout, AffectedSystem, ConfigFilePath, RuntimeStateFilePath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

func (instance *Msvm_VirtualSystemReferencePointService) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}
