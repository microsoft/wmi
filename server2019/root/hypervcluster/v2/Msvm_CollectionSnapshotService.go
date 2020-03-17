// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_CollectionSnapshotService struct
type Msvm_CollectionSnapshotService struct {
	CIM_Service
}

//

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>
// <param name="ResultingSnapshotCollection" type="CIM_Collection "></param>
// <param name="SnapshotSettings" type="string "></param>
// <param name="SnapshotType" type="CollectionSnapshotService_SnapshotType "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ResultingSnapshotCollection" type="CIM_Collection "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionSnapshotService) CreateSnapshot( /* IN */ Collection CIM_CollectionOfMSEs,
	/* IN */ SnapshotSettings string,
	/* IN */ SnapshotType CollectionSnapshotService_SnapshotType,
	/* IN/OUT */ ResultingSnapshotCollection CIM_Collection,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("CreateSnapshot", Action, PercentComplete, Timeout, Collection, SnapshotSettings, SnapshotType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AffectedSnapshotCollection" type="CIM_Collection "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionSnapshotService) DestroySnapshot( /* IN */ AffectedSnapshotCollection CIM_Collection,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("DestroySnapshot", Action, PercentComplete, Timeout, AffectedSnapshotCollection)
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
// <param name="SnapshotCollection" type="CIM_Collection "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionSnapshotService) ExportSnapshot( /* IN */ SnapshotCollection CIM_Collection,
	/* IN */ ExportDirectory string,
	/* IN */ ExportSettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ExportSnapshot", Action, PercentComplete, Timeout, SnapshotCollection, ExportDirectory, ExportSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="SnapshotCollection" type="CIM_Collection "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionSnapshotService) ApplySnapshot( /* IN */ SnapshotCollection CIM_Collection,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ApplySnapshot", Action, PercentComplete, Timeout, SnapshotCollection)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AffectedSnapshotCollection" type="Msvm_SnapshotCollection "></param>
// <param name="ResultingReferencePointCollection" type="Msvm_ReferencePointCollection "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ResultingReferencePointCollection" type="Msvm_ReferencePointCollection "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionSnapshotService) ConvertToReferencePoint( /* IN */ AffectedSnapshotCollection Msvm_SnapshotCollection,
	/* IN/OUT */ ResultingReferencePointCollection Msvm_ReferencePointCollection,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ConvertToReferencePoint", Action, PercentComplete, Timeout, AffectedSnapshotCollection)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
