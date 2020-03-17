// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// Msvm_CollectionReplicationStatistics struct
type Msvm_CollectionReplicationStatistics struct {
	CIM_ManagedElement

	//
	ApplicationConsistentSnapshotFailureCount uint32

	//
	EndStatisticTime string

	//
	LastTestFailoverTime string

	//
	LastWALReplicationTime string

	//
	MaxReplicationLatency uint32

	//
	MaxReplicationSize uint64

	//
	NetworkFailureCount uint32

	//
	PendingReplicationSize uint64

	//
	ReplicationFailureCount uint32

	//
	ReplicationHealth uint32

	//
	ReplicationSize uint64

	//
	ReplicationWALMissCount uint32

	//
	ReplicationWALSuccessCount uint32

	//
	StartStatisticTime string
}

// SetApplicationConsistentSnapshotFailureCount sets the value of ApplicationConsistentSnapshotFailureCount for the instance
func (instance *Msvm_CollectionReplicationStatistics) SetPropertyApplicationConsistentSnapshotFailureCount(value uint32) (err error) {
	return instance.SetProperty("ApplicationConsistentSnapshotFailureCount", value)
}

// GetApplicationConsistentSnapshotFailureCount gets the value of ApplicationConsistentSnapshotFailureCount for the instance
func (instance *Msvm_CollectionReplicationStatistics) GetPropertyApplicationConsistentSnapshotFailureCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ApplicationConsistentSnapshotFailureCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEndStatisticTime sets the value of EndStatisticTime for the instance
func (instance *Msvm_CollectionReplicationStatistics) SetPropertyEndStatisticTime(value string) (err error) {
	return instance.SetProperty("EndStatisticTime", value)
}

// GetEndStatisticTime gets the value of EndStatisticTime for the instance
func (instance *Msvm_CollectionReplicationStatistics) GetPropertyEndStatisticTime() (value string, err error) {
	retValue, err := instance.GetProperty("EndStatisticTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastTestFailoverTime sets the value of LastTestFailoverTime for the instance
func (instance *Msvm_CollectionReplicationStatistics) SetPropertyLastTestFailoverTime(value string) (err error) {
	return instance.SetProperty("LastTestFailoverTime", value)
}

// GetLastTestFailoverTime gets the value of LastTestFailoverTime for the instance
func (instance *Msvm_CollectionReplicationStatistics) GetPropertyLastTestFailoverTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastTestFailoverTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastWALReplicationTime sets the value of LastWALReplicationTime for the instance
func (instance *Msvm_CollectionReplicationStatistics) SetPropertyLastWALReplicationTime(value string) (err error) {
	return instance.SetProperty("LastWALReplicationTime", value)
}

// GetLastWALReplicationTime gets the value of LastWALReplicationTime for the instance
func (instance *Msvm_CollectionReplicationStatistics) GetPropertyLastWALReplicationTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastWALReplicationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxReplicationLatency sets the value of MaxReplicationLatency for the instance
func (instance *Msvm_CollectionReplicationStatistics) SetPropertyMaxReplicationLatency(value uint32) (err error) {
	return instance.SetProperty("MaxReplicationLatency", value)
}

// GetMaxReplicationLatency gets the value of MaxReplicationLatency for the instance
func (instance *Msvm_CollectionReplicationStatistics) GetPropertyMaxReplicationLatency() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxReplicationLatency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxReplicationSize sets the value of MaxReplicationSize for the instance
func (instance *Msvm_CollectionReplicationStatistics) SetPropertyMaxReplicationSize(value uint64) (err error) {
	return instance.SetProperty("MaxReplicationSize", value)
}

// GetMaxReplicationSize gets the value of MaxReplicationSize for the instance
func (instance *Msvm_CollectionReplicationStatistics) GetPropertyMaxReplicationSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxReplicationSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkFailureCount sets the value of NetworkFailureCount for the instance
func (instance *Msvm_CollectionReplicationStatistics) SetPropertyNetworkFailureCount(value uint32) (err error) {
	return instance.SetProperty("NetworkFailureCount", value)
}

// GetNetworkFailureCount gets the value of NetworkFailureCount for the instance
func (instance *Msvm_CollectionReplicationStatistics) GetPropertyNetworkFailureCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("NetworkFailureCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPendingReplicationSize sets the value of PendingReplicationSize for the instance
func (instance *Msvm_CollectionReplicationStatistics) SetPropertyPendingReplicationSize(value uint64) (err error) {
	return instance.SetProperty("PendingReplicationSize", value)
}

// GetPendingReplicationSize gets the value of PendingReplicationSize for the instance
func (instance *Msvm_CollectionReplicationStatistics) GetPropertyPendingReplicationSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("PendingReplicationSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplicationFailureCount sets the value of ReplicationFailureCount for the instance
func (instance *Msvm_CollectionReplicationStatistics) SetPropertyReplicationFailureCount(value uint32) (err error) {
	return instance.SetProperty("ReplicationFailureCount", value)
}

// GetReplicationFailureCount gets the value of ReplicationFailureCount for the instance
func (instance *Msvm_CollectionReplicationStatistics) GetPropertyReplicationFailureCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReplicationFailureCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplicationHealth sets the value of ReplicationHealth for the instance
func (instance *Msvm_CollectionReplicationStatistics) SetPropertyReplicationHealth(value uint32) (err error) {
	return instance.SetProperty("ReplicationHealth", value)
}

// GetReplicationHealth gets the value of ReplicationHealth for the instance
func (instance *Msvm_CollectionReplicationStatistics) GetPropertyReplicationHealth() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReplicationHealth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplicationSize sets the value of ReplicationSize for the instance
func (instance *Msvm_CollectionReplicationStatistics) SetPropertyReplicationSize(value uint64) (err error) {
	return instance.SetProperty("ReplicationSize", value)
}

// GetReplicationSize gets the value of ReplicationSize for the instance
func (instance *Msvm_CollectionReplicationStatistics) GetPropertyReplicationSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReplicationSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplicationWALMissCount sets the value of ReplicationWALMissCount for the instance
func (instance *Msvm_CollectionReplicationStatistics) SetPropertyReplicationWALMissCount(value uint32) (err error) {
	return instance.SetProperty("ReplicationWALMissCount", value)
}

// GetReplicationWALMissCount gets the value of ReplicationWALMissCount for the instance
func (instance *Msvm_CollectionReplicationStatistics) GetPropertyReplicationWALMissCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReplicationWALMissCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplicationWALSuccessCount sets the value of ReplicationWALSuccessCount for the instance
func (instance *Msvm_CollectionReplicationStatistics) SetPropertyReplicationWALSuccessCount(value uint32) (err error) {
	return instance.SetProperty("ReplicationWALSuccessCount", value)
}

// GetReplicationWALSuccessCount gets the value of ReplicationWALSuccessCount for the instance
func (instance *Msvm_CollectionReplicationStatistics) GetPropertyReplicationWALSuccessCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReplicationWALSuccessCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartStatisticTime sets the value of StartStatisticTime for the instance
func (instance *Msvm_CollectionReplicationStatistics) SetPropertyStartStatisticTime(value string) (err error) {
	return instance.SetProperty("StartStatisticTime", value)
}

// GetStartStatisticTime gets the value of StartStatisticTime for the instance
func (instance *Msvm_CollectionReplicationStatistics) GetPropertyStartStatisticTime() (value string, err error) {
	retValue, err := instance.GetProperty("StartStatisticTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
