// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// Msvm_VirtualSystemCollection struct
type Msvm_VirtualSystemCollection struct {
	CIM_CollectionOfMSEs

	//
	FailedOverReplicationType uint16

	//
	LastApplyConsistencyLevel uint16

	//
	LastApplyTime string

	//
	LastApplyVirtualMachineIds []string

	//
	ReplicationMode uint16

	//
	ReplicationState uint16
}

// SetFailedOverReplicationType sets the value of FailedOverReplicationType for the instance
func (instance *Msvm_VirtualSystemCollection) SetPropertyFailedOverReplicationType(value uint16) (err error) {
	return instance.SetProperty("FailedOverReplicationType", value)
}

// GetFailedOverReplicationType gets the value of FailedOverReplicationType for the instance
func (instance *Msvm_VirtualSystemCollection) GetPropertyFailedOverReplicationType() (value uint16, err error) {
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

// SetLastApplyConsistencyLevel sets the value of LastApplyConsistencyLevel for the instance
func (instance *Msvm_VirtualSystemCollection) SetPropertyLastApplyConsistencyLevel(value uint16) (err error) {
	return instance.SetProperty("LastApplyConsistencyLevel", value)
}

// GetLastApplyConsistencyLevel gets the value of LastApplyConsistencyLevel for the instance
func (instance *Msvm_VirtualSystemCollection) GetPropertyLastApplyConsistencyLevel() (value uint16, err error) {
	retValue, err := instance.GetProperty("LastApplyConsistencyLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastApplyTime sets the value of LastApplyTime for the instance
func (instance *Msvm_VirtualSystemCollection) SetPropertyLastApplyTime(value string) (err error) {
	return instance.SetProperty("LastApplyTime", value)
}

// GetLastApplyTime gets the value of LastApplyTime for the instance
func (instance *Msvm_VirtualSystemCollection) GetPropertyLastApplyTime() (value string, err error) {
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

// SetLastApplyVirtualMachineIds sets the value of LastApplyVirtualMachineIds for the instance
func (instance *Msvm_VirtualSystemCollection) SetPropertyLastApplyVirtualMachineIds(value []string) (err error) {
	return instance.SetProperty("LastApplyVirtualMachineIds", value)
}

// GetLastApplyVirtualMachineIds gets the value of LastApplyVirtualMachineIds for the instance
func (instance *Msvm_VirtualSystemCollection) GetPropertyLastApplyVirtualMachineIds() (value []string, err error) {
	retValue, err := instance.GetProperty("LastApplyVirtualMachineIds")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplicationMode sets the value of ReplicationMode for the instance
func (instance *Msvm_VirtualSystemCollection) SetPropertyReplicationMode(value uint16) (err error) {
	return instance.SetProperty("ReplicationMode", value)
}

// GetReplicationMode gets the value of ReplicationMode for the instance
func (instance *Msvm_VirtualSystemCollection) GetPropertyReplicationMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationMode")
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
func (instance *Msvm_VirtualSystemCollection) SetPropertyReplicationState(value uint16) (err error) {
	return instance.SetProperty("ReplicationState", value)
}

// GetReplicationState gets the value of ReplicationState for the instance
func (instance *Msvm_VirtualSystemCollection) GetPropertyReplicationState() (value uint16, err error) {
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
