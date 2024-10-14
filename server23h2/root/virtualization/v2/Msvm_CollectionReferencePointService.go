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

// Msvm_CollectionReferencePointService struct
type Msvm_CollectionReferencePointService struct {
	*CIM_Service
}

func NewMsvm_CollectionReferencePointServiceEx1(instance *cim.WmiInstance) (newInstance *Msvm_CollectionReferencePointService, err error) {
	tmp, err := NewCIM_ServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectionReferencePointService{
		CIM_Service: tmp,
	}
	return
}

func NewMsvm_CollectionReferencePointServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_CollectionReferencePointService, err error) {
	tmp, err := NewCIM_ServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectionReferencePointService{
		CIM_Service: tmp,
	}
	return
}

//

// <param name="Collection" type="Msvm_VirtualSystemCollection "></param>
// <param name="ReferencePointSettings" type="string "></param>
// <param name="ReferencePointType" type="CollectionReferencePointService_ReferencePointType "></param>
// <param name="ResultingReferencePointCollection" type="CIM_Collection "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ResultingReferencePointCollection" type="CIM_Collection "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReferencePointService) CreateReferencePoint( /* IN */ Collection Msvm_VirtualSystemCollection,
	/* IN */ ReferencePointSettings string,
	/* IN */ ReferencePointType CollectionReferencePointService_ReferencePointType,
	/* IN/OUT */ ResultingReferencePointCollection CIM_Collection,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("CreateReferencePoint", Action, PercentComplete, Timeout, Collection, ReferencePointSettings, ReferencePointType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AffectedReferencePointCollection" type="CIM_Collection "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReferencePointService) DestroyReferencePoint( /* IN */ AffectedReferencePointCollection CIM_Collection,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("DestroyReferencePoint", Action, PercentComplete, Timeout, AffectedReferencePointCollection)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AffectedReferencePointCollection" type="CIM_Collection "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReferencePointService) RemoveAssociatedData( /* IN */ AffectedReferencePointCollection CIM_Collection,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RemoveAssociatedData", Action, PercentComplete, Timeout, AffectedReferencePointCollection)
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
// <param name="ReferencePointCollection" type="CIM_Collection "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReferencePointService) ExportReferencePoint( /* IN */ ReferencePointCollection CIM_Collection,
	/* IN */ ExportDirectory string,
	/* IN */ ExportSettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ExportReferencePoint", Action, PercentComplete, Timeout, ReferencePointCollection, ExportDirectory, ExportSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

func (instance *Msvm_CollectionReferencePointService) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}
