// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_CollectionReplicationService struct
type Msvm_CollectionReplicationService struct { 
	*CIM_Service
}

	func NewMsvm_CollectionReplicationServiceEx1(instance *cim.WmiInstance) (newInstance *Msvm_CollectionReplicationService, err error) {tmp, err := NewCIM_ServiceEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_CollectionReplicationService {
	CIM_Service: tmp,
	}
	return
	}
	

	func NewMsvm_CollectionReplicationServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_CollectionReplicationService, err error) {tmp, err := NewCIM_ServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_CollectionReplicationService {
	CIM_Service: tmp,
	}
	return
	}
	

// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>
// <param name="CollectionReplicationSettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) CreateReplicationRelationship( /* IN */ Collection CIM_CollectionOfMSEs,
 /* IN */ CollectionReplicationSettingData string,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("CreateReplicationRelationship", Action, PercentComplete, Timeout , Collection, CollectionReplicationSettingData)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>
// <param name="InitialReplicationExportLocation" type="string "></param>
// <param name="InitialReplicationType" type="CollectionReplicationService_InitialReplicationType "></param>
// <param name="StartTime" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) StartReplication( /* IN */ Collection CIM_CollectionOfMSEs,
 /* IN */ InitialReplicationType CollectionReplicationService_InitialReplicationType,
 /* IN */ InitialReplicationExportLocation string,
 /* IN */ StartTime string,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("StartReplication", Action, PercentComplete, Timeout , Collection, InitialReplicationType, InitialReplicationExportLocation, StartTime)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>
// <param name="CollectionReplicationSettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) ModifyReplicationSettings( /* IN */ Collection CIM_CollectionOfMSEs,
 /* IN */ CollectionReplicationSettingData string,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("ModifyReplicationSettings", Action, PercentComplete, Timeout , Collection, CollectionReplicationSettingData)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>
// <param name="CollectionRecoveryPoint" type="Msvm_CollectionRecoveryPoint "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) InitiateFailover( /* IN */ Collection CIM_CollectionOfMSEs,
 /* IN */ CollectionRecoveryPoint Msvm_CollectionRecoveryPoint,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("InitiateFailover", Action, PercentComplete, Timeout , Collection, CollectionRecoveryPoint)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) RevertFailover( /* IN */ Collection CIM_CollectionOfMSEs,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("RevertFailover", Action, PercentComplete, Timeout , Collection)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) CommitFailover( /* IN */ Collection CIM_CollectionOfMSEs,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("CommitFailover", Action, PercentComplete, Timeout , Collection)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) InitiatePlannedFailover( /* IN */ Collection CIM_CollectionOfMSEs,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("InitiatePlannedFailover", Action, PercentComplete, Timeout , Collection)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) RevertPlannedFailover( /* IN */ Collection CIM_CollectionOfMSEs,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("RevertPlannedFailover", Action, PercentComplete, Timeout , Collection)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) RemoveReplicationRelationship( /* IN */ Collection CIM_CollectionOfMSEs,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("RemoveReplicationRelationship", Action, PercentComplete, Timeout , Collection)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) SuspendReplication( /* IN */ Collection CIM_CollectionOfMSEs,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("SuspendReplication", Action, PercentComplete, Timeout , Collection)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) ResumeReplication( /* IN */ Collection CIM_CollectionOfMSEs,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("ResumeReplication", Action, PercentComplete, Timeout , Collection)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) PrepareForReverseReplication( /* IN */ Collection CIM_CollectionOfMSEs,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("PrepareForReverseReplication", Action, PercentComplete, Timeout , Collection)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>
// <param name="CollectionReplicationSettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) ReverseReplicationRelationship( /* IN */ Collection CIM_CollectionOfMSEs,
 /* IN */ CollectionReplicationSettingData string,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("ReverseReplicationRelationship", Action, PercentComplete, Timeout , Collection, CollectionReplicationSettingData)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>
// <param name="IncludeSuspended" type="bool "></param>
// <param name="StartTime" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) Resynchronize( /* IN */ Collection CIM_CollectionOfMSEs,
 /* IN */ StartTime string,
 /* IN */ IncludeSuspended bool,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("Resynchronize", Action, PercentComplete, Timeout , Collection, StartTime, IncludeSuspended)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>
// <param name="CollectionRecoveryPoint" type="Msvm_CollectionRecoveryPoint "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ResultingCollection" type="CIM_CollectionOfMSEs "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) TestReplicaCollection( /* IN */ Collection CIM_CollectionOfMSEs,
 /* IN */ CollectionRecoveryPoint Msvm_CollectionRecoveryPoint,
 /* OUT */ ResultingCollection CIM_CollectionOfMSEs,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("TestReplicaCollection", Action, PercentComplete, Timeout , Collection, CollectionRecoveryPoint)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) CancelInitialReplication( /* IN */ Collection CIM_CollectionOfMSEs,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("CancelInitialReplication", Action, PercentComplete, Timeout , Collection)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) CancelResynchronize( /* IN */ Collection CIM_CollectionOfMSEs,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("CancelResynchronize", Action, PercentComplete, Timeout , Collection)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) RevertTestReplicaCollection( /* IN */ Collection CIM_CollectionOfMSEs,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("RevertTestReplicaCollection", Action, PercentComplete, Timeout , Collection)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) CancelUpdateDiskSet( /* IN */ Collection CIM_CollectionOfMSEs,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("CancelUpdateDiskSet", Action, PercentComplete, Timeout , Collection)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReplicationHealthIssues" type="string []"></param>
// <param name="ReplicationStatistics" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) GetReplicationStatistics( /* IN */ Collection CIM_CollectionOfMSEs,
 /* OUT */ ReplicationStatistics string,
 /* OUT */ ReplicationHealthIssues []string,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("GetReplicationStatistics", Action, PercentComplete, Timeout , Collection)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReplicationService) ResetReplicationStatistics( /* IN */ Collection CIM_CollectionOfMSEs,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("ResetReplicationStatistics", Action, PercentComplete, Timeout , Collection)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}

func  (instance* Msvm_CollectionReplicationService) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_ComputerSystem"); 
	}
	

