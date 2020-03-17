// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_ReplicationRelationship struct
type Msvm_ReplicationRelationship struct {
	CIM_ManagedSystemElement

	//
	FailedOverReplicationType uint16

	//
	LastApplicationConsistentReplicationTime string

	//
	LastApplyTime string

	//
	LastReplicationTime string

	//
	LastReplicationType uint16

	//
	ReplicationHealth uint16

	//
	ReplicationState uint16
}

// SetFailedOverReplicationType sets the value of FailedOverReplicationType for the instance
func (instance *Msvm_ReplicationRelationship) SetPropertyFailedOverReplicationType(value uint16) (err error) {
	return instance.SetProperty("FailedOverReplicationType", value)
}

// GetFailedOverReplicationType gets the value of FailedOverReplicationType for the instance
func (instance *Msvm_ReplicationRelationship) GetPropertyFailedOverReplicationType() (value uint16, err error) {
	retValue, err := instance.GetProperty("FailedOverReplicationType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastApplicationConsistentReplicationTime sets the value of LastApplicationConsistentReplicationTime for the instance
func (instance *Msvm_ReplicationRelationship) SetPropertyLastApplicationConsistentReplicationTime(value string) (err error) {
	return instance.SetProperty("LastApplicationConsistentReplicationTime", value)
}

// GetLastApplicationConsistentReplicationTime gets the value of LastApplicationConsistentReplicationTime for the instance
func (instance *Msvm_ReplicationRelationship) GetPropertyLastApplicationConsistentReplicationTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastApplicationConsistentReplicationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastApplyTime sets the value of LastApplyTime for the instance
func (instance *Msvm_ReplicationRelationship) SetPropertyLastApplyTime(value string) (err error) {
	return instance.SetProperty("LastApplyTime", value)
}

// GetLastApplyTime gets the value of LastApplyTime for the instance
func (instance *Msvm_ReplicationRelationship) GetPropertyLastApplyTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastApplyTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastReplicationTime sets the value of LastReplicationTime for the instance
func (instance *Msvm_ReplicationRelationship) SetPropertyLastReplicationTime(value string) (err error) {
	return instance.SetProperty("LastReplicationTime", value)
}

// GetLastReplicationTime gets the value of LastReplicationTime for the instance
func (instance *Msvm_ReplicationRelationship) GetPropertyLastReplicationTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastReplicationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastReplicationType sets the value of LastReplicationType for the instance
func (instance *Msvm_ReplicationRelationship) SetPropertyLastReplicationType(value uint16) (err error) {
	return instance.SetProperty("LastReplicationType", value)
}

// GetLastReplicationType gets the value of LastReplicationType for the instance
func (instance *Msvm_ReplicationRelationship) GetPropertyLastReplicationType() (value uint16, err error) {
	retValue, err := instance.GetProperty("LastReplicationType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplicationHealth sets the value of ReplicationHealth for the instance
func (instance *Msvm_ReplicationRelationship) SetPropertyReplicationHealth(value uint16) (err error) {
	return instance.SetProperty("ReplicationHealth", value)
}

// GetReplicationHealth gets the value of ReplicationHealth for the instance
func (instance *Msvm_ReplicationRelationship) GetPropertyReplicationHealth() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationHealth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplicationState sets the value of ReplicationState for the instance
func (instance *Msvm_ReplicationRelationship) SetPropertyReplicationState(value uint16) (err error) {
	return instance.SetProperty("ReplicationState", value)
}

// GetReplicationState gets the value of ReplicationState for the instance
func (instance *Msvm_ReplicationRelationship) GetPropertyReplicationState() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
