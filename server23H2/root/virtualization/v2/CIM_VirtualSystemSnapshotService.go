// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_VirtualSystemSnapshotService struct
type CIM_VirtualSystemSnapshotService struct {
	*CIM_Service
}

func NewCIM_VirtualSystemSnapshotServiceEx1(instance *cim.WmiInstance) (newInstance *CIM_VirtualSystemSnapshotService, err error) {
	tmp, err := NewCIM_ServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_VirtualSystemSnapshotService{
		CIM_Service: tmp,
	}
	return
}

func NewCIM_VirtualSystemSnapshotServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_VirtualSystemSnapshotService, err error) {
	tmp, err := NewCIM_ServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_VirtualSystemSnapshotService{
		CIM_Service: tmp,
	}
	return
}

// Creates a snapshot of a virtual system.

// <param name="AffectedSystem" type="CIM_ComputerSystem ">Reference to the affected virtual system.</param>
// <param name="ResultingSnapshot" type="CIM_VirtualSystemSettingData ">Resulting virtual system snapshot</param>
// <param name="SnapshotSettings" type="string ">Parameter settings.</param>
// <param name="SnapshotType" type="VirtualSystemSnapshotService_SnapshotType ">Requested snapshot type:
///Full Snapshot: Complete snapshot of the virtual system.
///Disk Snapshot: Snapshot of virtual system disks.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned. In this case, the instance of the CIM_VirtualSystemSettingData class representing the new virtual system snapshot is presented via the CIM_AffectedJobElement association with the value of the AffectedElement property referring to the new instance of the CIM_VirtualSystemSettingData class representing the virtual system snapshot and and the value of the ElementEffects set to 5 (Create).</param>
// <param name="ResultingSnapshot" type="CIM_VirtualSystemSettingData ">Resulting virtual system snapshot</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_VirtualSystemSnapshotService) CreateSnapshot( /* IN */ AffectedSystem CIM_ComputerSystem,
	/* IN */ SnapshotSettings string,
	/* IN */ SnapshotType VirtualSystemSnapshotService_SnapshotType,
	/* IN/OUT */ ResultingSnapshot CIM_VirtualSystemSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("CreateSnapshot", Action, PercentComplete, Timeout, AffectedSystem, SnapshotSettings, SnapshotType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Destroy an existing virtual system snapshot.This method may as a side effect destroy other snapshots that are dependent on the affected snapshot.

// <param name="AffectedSnapshot" type="CIM_VirtualSystemSettingData ">Reference to the affected virtual system snapshot.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_VirtualSystemSnapshotService) DestroySnapshot( /* IN */ AffectedSnapshot CIM_VirtualSystemSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("DestroySnapshot", Action, PercentComplete, Timeout, AffectedSnapshot)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Apply a virtual system snapshot to the virtual system that it was created from.

// <param name="Snapshot" type="CIM_VirtualSystemSettingData ">Reference to the virtual system snapshot.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_VirtualSystemSnapshotService) ApplySnapshot( /* IN */ Snapshot CIM_VirtualSystemSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ApplySnapshot", Action, PercentComplete, Timeout, Snapshot)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
