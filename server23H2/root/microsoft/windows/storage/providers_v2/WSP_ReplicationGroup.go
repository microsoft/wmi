// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WSP_ReplicationGroup struct
type WSP_ReplicationGroup struct {
	*MSFT_ReplicationGroup
}

func NewWSP_ReplicationGroupEx1(instance *cim.WmiInstance) (newInstance *WSP_ReplicationGroup, err error) {
	tmp, err := NewMSFT_ReplicationGroupEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_ReplicationGroup{
		MSFT_ReplicationGroup: tmp,
	}
	return
}

func NewWSP_ReplicationGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_ReplicationGroup, err error) {
	tmp, err := NewMSFT_ReplicationGroupEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_ReplicationGroup{
		MSFT_ReplicationGroup: tmp,
	}
	return
}

//

// <param name="Description" type="string "></param>
// <param name="FriendlyName" type="string "></param>
// <param name="LogDevice" type="string "></param>
// <param name="LogSizeInBytes" type="uint64 "></param>
// <param name="ReplicationQuorum" type="uint16 "></param>
// <param name="StorageElements" type="string []"></param>
// <param name="SyncMode" type="uint16 "></param>

// <param name="CreatedReplicationGroup" type="string "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *WSP_ReplicationGroup) CreateReplicationGroup( /* IN */ FriendlyName string,
	/* IN */ Description string,
	/* IN */ StorageElements []string,
	/* IN */ LogDevice string,
	/* IN */ LogSizeInBytes uint64,
	/* IN */ ReplicationQuorum uint16,
	/* IN */ SyncMode uint16,
	/* OUT */ CreatedReplicationGroup string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateReplicationGroup", FriendlyName, Description, StorageElements, LogDevice, LogSizeInBytes, ReplicationQuorum, SyncMode)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RecoveryPointObjective" type="uint16 "></param>
// <param name="SourceGroupSettings" type="MSFT_ReplicationSettings "></param>
// <param name="SourceReplicationGroupDescription" type="string "></param>
// <param name="SourceReplicationGroupFriendlyName" type="string "></param>
// <param name="SourceStorageElements" type="MSFT_StorageObject []"></param>
// <param name="TargetGroupSettings" type="MSFT_ReplicationSettings "></param>
// <param name="TargetReplicationGroupDescription" type="string "></param>
// <param name="TargetReplicationGroupFriendlyName" type="string "></param>
// <param name="TargetStorageElements" type="MSFT_StorageObject []"></param>
// <param name="TargetStoragePool" type="MSFT_StoragePool "></param>
// <param name="TargetStorageSubsystem" type="MSFT_ReplicaPeer "></param>

// <param name="CreatedReplicaPeer" type="MSFT_ReplicaPeer "></param>
// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SourceGroup" type="MSFT_ReplicationGroup "></param>
// <param name="TargetGroup" type="MSFT_ReplicationGroup "></param>
func (instance *WSP_ReplicationGroup) CreateReplicationRelationship( /* IN */ TargetStorageSubsystem MSFT_ReplicaPeer,
	/* IN */ SourceReplicationGroupFriendlyName string,
	/* IN */ SourceReplicationGroupDescription string,
	/* IN */ SourceStorageElements []MSFT_StorageObject,
	/* IN */ SourceGroupSettings MSFT_ReplicationSettings,
	/* IN */ TargetReplicationGroupFriendlyName string,
	/* IN */ TargetReplicationGroupDescription string,
	/* IN */ TargetStorageElements []MSFT_StorageObject,
	/* IN */ TargetStoragePool MSFT_StoragePool,
	/* IN */ TargetGroupSettings MSFT_ReplicationSettings,
	/* IN */ RecoveryPointObjective uint16,
	/* OUT */ SourceGroup MSFT_ReplicationGroup,
	/* OUT */ TargetGroup MSFT_ReplicationGroup,
	/* OUT */ CreatedReplicaPeer MSFT_ReplicaPeer,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateReplicationRelationship", TargetStorageSubsystem, SourceReplicationGroupFriendlyName, SourceReplicationGroupDescription, SourceStorageElements, SourceGroupSettings, TargetReplicationGroupFriendlyName, TargetReplicationGroupDescription, TargetStorageElements, TargetStoragePool, TargetGroupSettings, RecoveryPointObjective)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="SourceReplicationGroup" type="MSFT_ReplicationGroup "></param>
// <param name="TargetGroupReplicaPeer" type="MSFT_ReplicaPeer "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *WSP_ReplicationGroup) DeleteReplicationRelationship( /* IN */ SourceReplicationGroup MSFT_ReplicationGroup,
	/* IN */ TargetGroupReplicaPeer MSFT_ReplicaPeer,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DeleteReplicationRelationship", SourceReplicationGroup, TargetGroupReplicaPeer)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
