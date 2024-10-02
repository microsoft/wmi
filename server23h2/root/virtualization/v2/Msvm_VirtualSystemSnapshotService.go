// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VirtualSystemSnapshotService struct
type Msvm_VirtualSystemSnapshotService struct {
	*CIM_VirtualSystemSnapshotService
}

func NewMsvm_VirtualSystemSnapshotServiceEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemSnapshotService, err error) {
	tmp, err := NewCIM_VirtualSystemSnapshotServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemSnapshotService{
		CIM_VirtualSystemSnapshotService: tmp,
	}
	return
}

func NewMsvm_VirtualSystemSnapshotServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemSnapshotService, err error) {
	tmp, err := NewCIM_VirtualSystemSnapshotServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemSnapshotService{
		CIM_VirtualSystemSnapshotService: tmp,
	}
	return
}

//

// <param name="SnapshotSettingData" type="CIM_VirtualSystemSettingData "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemSnapshotService) DestroySnapshotTree( /* IN */ SnapshotSettingData CIM_VirtualSystemSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("DestroySnapshotTree", Action, PercentComplete, Timeout, SnapshotSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="SnapshotSettingData" type="CIM_VirtualSystemSettingData "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemSnapshotService) ClearSnapshotState( /* IN */ SnapshotSettingData CIM_VirtualSystemSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ClearSnapshotState", Action, PercentComplete, Timeout, SnapshotSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Convert an existing virtual system snapshot to a reference point. The snapshot gets deleted as a side effect. Only recovery snapshots can be converted to reference points.

// <param name="AffectedSnapshot" type="CIM_VirtualSystemSettingData ">Reference to the affected virtual system snapshot.</param>
// <param name="ReferencePointSettings" type="string ">Parameter settings.</param>
// <param name="ResultingReferencePoint" type="Msvm_VirtualSystemReferencePoint ">Resulting virtual system reference point</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned.</param>
// <param name="ResultingReferencePoint" type="Msvm_VirtualSystemReferencePoint ">Resulting virtual system reference point</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemSnapshotService) ConvertToReferencePoint( /* IN */ AffectedSnapshot CIM_VirtualSystemSettingData,
	/* IN */ ReferencePointSettings string,
	/* IN/OUT */ ResultingReferencePoint Msvm_VirtualSystemReferencePoint,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ConvertToReferencePoint", Action, PercentComplete, Timeout, AffectedSnapshot, ReferencePointSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

func (instance *Msvm_VirtualSystemSnapshotService) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}
