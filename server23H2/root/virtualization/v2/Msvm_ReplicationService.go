// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ReplicationService struct
type Msvm_ReplicationService struct {
	*CIM_Service
}

func NewMsvm_ReplicationServiceEx1(instance *cim.WmiInstance) (newInstance *Msvm_ReplicationService, err error) {
	tmp, err := NewCIM_ServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReplicationService{
		CIM_Service: tmp,
	}
	return
}

func NewMsvm_ReplicationServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ReplicationService, err error) {
	tmp, err := NewCIM_ServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReplicationService{
		CIM_Service: tmp,
	}
	return
}

//

// <param name="SettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) ModifyServiceSettings( /* IN */ SettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifyServiceSettings", Action, PercentComplete, Timeout, SettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AuthorizationEntry" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) AddAuthorizationEntry( /* IN */ AuthorizationEntry string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("AddAuthorizationEntry", Action, PercentComplete, Timeout, AuthorizationEntry)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AuthorizationEntry" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) ModifyAuthorizationEntry( /* IN */ AuthorizationEntry string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifyAuthorizationEntry", Action, PercentComplete, Timeout, AuthorizationEntry)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AllowedPrimaryHostSystem" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) RemoveAuthorizationEntry( /* IN */ AllowedPrimaryHostSystem string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RemoveAuthorizationEntry", Action, PercentComplete, Timeout, AllowedPrimaryHostSystem)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AuthenticationType" type="uint16 "></param>
// <param name="BypassProxyServer" type="bool "></param>
// <param name="CertificateThumbPrint" type="string "></param>
// <param name="RecoveryConnectionPoint" type="string "></param>
// <param name="RecoveryServerPortNumber" type="uint16 "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) TestReplicationConnection( /* IN */ RecoveryConnectionPoint string,
	/* IN */ RecoveryServerPortNumber uint16,
	/* IN */ AuthenticationType uint16,
	/* IN */ CertificateThumbPrint string,
	/* IN */ BypassProxyServer bool,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("TestReplicationConnection", Action, PercentComplete, Timeout, RecoveryConnectionPoint, RecoveryServerPortNumber, AuthenticationType, CertificateThumbPrint, BypassProxyServer)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="ReplicationSettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) CreateReplicationRelationship( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ ReplicationSettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("CreateReplicationRelationship", Action, PercentComplete, Timeout, ComputerSystem, ReplicationSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="ReplicationSettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) ModifyReplicationSettings( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ ReplicationSettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifyReplicationSettings", Action, PercentComplete, Timeout, ComputerSystem, ReplicationSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) RemoveReplicationRelationship( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RemoveReplicationRelationship", Action, PercentComplete, Timeout, ComputerSystem)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="ReplicationRelationship" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) RemoveReplicationRelationshipEx( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ ReplicationRelationship string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RemoveReplicationRelationshipEx", Action, PercentComplete, Timeout, ComputerSystem, ReplicationRelationship)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="InitialReplicationExportLocation" type="string "></param>
// <param name="InitialReplicationType" type="ReplicationService_InitialReplicationType "></param>
// <param name="StartTime" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) StartReplication( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ InitialReplicationType ReplicationService_InitialReplicationType,
	/* IN */ InitialReplicationExportLocation string,
	/* IN */ StartTime string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("StartReplication", Action, PercentComplete, Timeout, ComputerSystem, InitialReplicationType, InitialReplicationExportLocation, StartTime)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="InitialReplicationImportLocation" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) ImportInitialReplica( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ InitialReplicationImportLocation string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ImportInitialReplica", Action, PercentComplete, Timeout, ComputerSystem, InitialReplicationImportLocation)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="ReplicationSettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) ReverseReplicationRelationship( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ ReplicationSettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ReverseReplicationRelationship", Action, PercentComplete, Timeout, ComputerSystem, ReplicationSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="SnapshotSettingData" type="CIM_VirtualSystemSettingData "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) InitiateFailover( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ SnapshotSettingData CIM_VirtualSystemSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("InitiateFailover", Action, PercentComplete, Timeout, ComputerSystem, SnapshotSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) RevertFailover( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RevertFailover", Action, PercentComplete, Timeout, ComputerSystem)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) CommitFailover( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("CommitFailover", Action, PercentComplete, Timeout, ComputerSystem)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="SnapshotSettingData" type="CIM_VirtualSystemSettingData "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ResultingSystem" type="CIM_ComputerSystem "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) TestReplicaSystem( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ SnapshotSettingData CIM_VirtualSystemSettingData,
	/* OUT */ ResultingSystem CIM_ComputerSystem,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("TestReplicaSystem", Action, PercentComplete, Timeout, ComputerSystem, SnapshotSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="StartTime" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) Resynchronize( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ StartTime string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("Resynchronize", Action, PercentComplete, Timeout, ComputerSystem, StartTime)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="NetworkSettings" type="string []"></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) SetFailoverNetworkAdapterSettings( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ NetworkSettings []string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("SetFailoverNetworkAdapterSettings", Action, PercentComplete, Timeout, ComputerSystem, NetworkSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReplicationHealthIssues" type="string []"></param>
// <param name="ReplicationStatistics" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) GetReplicationStatistics( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* OUT */ ReplicationStatistics []string,
	/* OUT */ ReplicationHealthIssues []string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("GetReplicationStatistics", Action, PercentComplete, Timeout, ComputerSystem)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="ReplicationRelationship" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReplicationHealthIssues" type="string []"></param>
// <param name="ReplicationStatistics" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) GetReplicationStatisticsEx( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ ReplicationRelationship string,
	/* OUT */ ReplicationStatistics []string,
	/* OUT */ ReplicationHealthIssues []string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("GetReplicationStatisticsEx", Action, PercentComplete, Timeout, ComputerSystem, ReplicationRelationship)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) ResetReplicationStatistics( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ResetReplicationStatistics", Action, PercentComplete, Timeout, ComputerSystem)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="ReplicationRelationship" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) ResetReplicationStatisticsEx( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ ReplicationRelationship string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ResetReplicationStatisticsEx", Action, PercentComplete, Timeout, ComputerSystem, ReplicationRelationship)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AuthorizationEntry" type="string "></param>
// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) SetAuthorizationEntry( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ AuthorizationEntry string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("SetAuthorizationEntry", Action, PercentComplete, Timeout, ComputerSystem, AuthorizationEntry)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="EncodedCertificates" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) GetSystemCertificates( /* OUT */ EncodedCertificates []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSystemCertificates")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="ReplicationRelationship" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) ChangeReplicationModeToPrimary( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ ReplicationRelationship string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ChangeReplicationModeToPrimary", Action, PercentComplete, Timeout, ComputerSystem, ReplicationRelationship)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="RecoveryPointIdentifier" type="string "></param>
// <param name="ReplicationSettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ReplicationService) InitiateFailback( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ ReplicationSettingData string,
	/* IN */ RecoveryPointIdentifier string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("InitiateFailback", Action, PercentComplete, Timeout, ComputerSystem, ReplicationSettingData, RecoveryPointIdentifier)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

func (instance *Msvm_ReplicationService) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_ReplicationService) GetRelatedReplicationServiceSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ReplicationServiceSettingData")
}
