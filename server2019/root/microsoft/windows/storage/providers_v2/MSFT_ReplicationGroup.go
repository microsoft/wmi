// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

// MSFT_ReplicationGroup struct
type MSFT_ReplicationGroup struct {
	MSFT_StorageObject

	// A user-friendly string representing the description of the replication group.
	Description string

	// A user-friendly string representing the name of the replication group.
	FriendlyName string

	// Denotes the current health status of the replication group. Health of a group is derived from the health of the backing storage replicas.
	/// 0 - 'Healthy': All replicas are in a healthy state.
	///1 - 'Warning': The majority of replicas are healthy, but one or more may be not fully synchronized.
	///2 - 'Unhealthy': The majority of replicas are unhealthy or in a failed state.
	HealthStatus ReplicationGroup_HealthStatus

	// Indicates the current operating conditions of the group. Unlike HealthStatus, this field indicates the status of hardware, software, and infrastructure issues related to this group, and can contain multiple values.
	OperationalStatus []ReplicationGroup_OperationalStatus
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_ReplicationGroup) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_ReplicationGroup) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_ReplicationGroup) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", value)
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_ReplicationGroup) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *MSFT_ReplicationGroup) SetPropertyHealthStatus(value ReplicationGroup_HealthStatus) (err error) {
	return instance.SetProperty("HealthStatus", value)
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSFT_ReplicationGroup) GetPropertyHealthStatus() (value ReplicationGroup_HealthStatus, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(ReplicationGroup_HealthStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_ReplicationGroup) SetPropertyOperationalStatus(value []ReplicationGroup_OperationalStatus) (err error) {
	return instance.SetProperty("OperationalStatus", value)
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_ReplicationGroup) GetPropertyOperationalStatus() (value []ReplicationGroup_OperationalStatus, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]ReplicationGroup_OperationalStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="FriendlyName" type="string "></param>
// <param name="RecoveryPointObjective" type="uint32 "></param>
// <param name="ReplicationSettings" type="MSFT_ReplicationSettings "></param>
// <param name="SyncType" type="uint16 "></param>
// <param name="TargetGroupObjectId" type="string "></param>
// <param name="TargetStoragePoolObjectId" type="string "></param>
// <param name="TargetStorageSubsystem" type="MSFT_ReplicaPeer "></param>

// <param name="CreatedReplicaPeer" type="MSFT_ReplicaPeer "></param>
// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ReplicationGroup) CreateReplica( /* IN */ FriendlyName string,
	/* IN */ TargetStorageSubsystem MSFT_ReplicaPeer,
	/* IN */ TargetGroupObjectId string,
	/* IN */ TargetStoragePoolObjectId string,
	/* IN */ RecoveryPointObjective uint32,
	/* IN */ ReplicationSettings MSFT_ReplicationSettings,
	/* IN */ SyncType uint16,
	/* OUT */ CreatedReplicaPeer MSFT_ReplicaPeer,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateReplica", FriendlyName, TargetStorageSubsystem, TargetGroupObjectId, TargetStoragePoolObjectId, RecoveryPointObjective, ReplicationSettings, SyncType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Operation" type="uint16 "></param>
// <param name="SourceStorageObjects" type="MSFT_StorageObject []"></param>
// <param name="SyncPairs" type="MSFT_Synchronized []"></param>
// <param name="TargetGroup" type="MSFT_ReplicaPeer "></param>
// <param name="TargetStorageObjects" type="MSFT_StorageObject []"></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ReplicationGroup) SetReplicationRelationship( /* IN */ Operation uint16,
	/* IN */ TargetGroup MSFT_ReplicaPeer,
	/* IN */ SourceStorageObjects []MSFT_StorageObject,
	/* IN */ TargetStorageObjects []MSFT_StorageObject,
	/* IN */ SyncPairs []MSFT_Synchronized,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetReplicationRelationship", Operation, TargetGroup, SourceStorageObjects, TargetStorageObjects, SyncPairs)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReplicationSettings" type="MSFT_ReplicationSettings "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ReplicationGroup) SetReplicationSettings( /* IN */ ReplicationSettings MSFT_ReplicationSettings,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetReplicationSettings", ReplicationSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReplicationSettings" type="MSFT_ReplicationSettings "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ReplicationGroup) GetReplicationSettings( /* OUT */ ReplicationSettings MSFT_ReplicationSettings,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetReplicationSettings")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="StorageObjects" type="MSFT_StorageObject []"></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ReplicationGroup) AddMember( /* IN */ StorageObjects []MSFT_StorageObject,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddMember", StorageObjects)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="StorageObjects" type="MSFT_StorageObject []"></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ReplicationGroup) RemoveMember( /* IN */ StorageObjects []MSFT_StorageObject,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemoveMember", StorageObjects)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="FriendlyName" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ReplicationGroup) SetFriendlyName( /* IN */ FriendlyName string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetFriendlyName", FriendlyName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ReplicationGroup) DeleteObject( /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DeleteObject")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
