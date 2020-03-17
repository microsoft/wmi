// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// Msvm_CollectionRecoveryPoint struct
type Msvm_CollectionRecoveryPoint struct {
	CIM_ManagedElement

	//
	CollectionID string

	//
	ConsistencyLevel uint16

	//
	CreationTime string

	//
	RecoveryPointIds []string

	//
	VirtualMachineIds []string
}

// SetCollectionID sets the value of CollectionID for the instance
func (instance *Msvm_CollectionRecoveryPoint) SetPropertyCollectionID(value string) (err error) {
	return instance.SetProperty("CollectionID", value)
}

// GetCollectionID gets the value of CollectionID for the instance
func (instance *Msvm_CollectionRecoveryPoint) GetPropertyCollectionID() (value string, err error) {
	retValue, err := instance.GetProperty("CollectionID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConsistencyLevel sets the value of ConsistencyLevel for the instance
func (instance *Msvm_CollectionRecoveryPoint) SetPropertyConsistencyLevel(value uint16) (err error) {
	return instance.SetProperty("ConsistencyLevel", value)
}

// GetConsistencyLevel gets the value of ConsistencyLevel for the instance
func (instance *Msvm_CollectionRecoveryPoint) GetPropertyConsistencyLevel() (value uint16, err error) {
	retValue, err := instance.GetProperty("ConsistencyLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCreationTime sets the value of CreationTime for the instance
func (instance *Msvm_CollectionRecoveryPoint) SetPropertyCreationTime(value string) (err error) {
	return instance.SetProperty("CreationTime", value)
}

// GetCreationTime gets the value of CreationTime for the instance
func (instance *Msvm_CollectionRecoveryPoint) GetPropertyCreationTime() (value string, err error) {
	retValue, err := instance.GetProperty("CreationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecoveryPointIds sets the value of RecoveryPointIds for the instance
func (instance *Msvm_CollectionRecoveryPoint) SetPropertyRecoveryPointIds(value []string) (err error) {
	return instance.SetProperty("RecoveryPointIds", value)
}

// GetRecoveryPointIds gets the value of RecoveryPointIds for the instance
func (instance *Msvm_CollectionRecoveryPoint) GetPropertyRecoveryPointIds() (value []string, err error) {
	retValue, err := instance.GetProperty("RecoveryPointIds")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualMachineIds sets the value of VirtualMachineIds for the instance
func (instance *Msvm_CollectionRecoveryPoint) SetPropertyVirtualMachineIds(value []string) (err error) {
	return instance.SetProperty("VirtualMachineIds", value)
}

// GetVirtualMachineIds gets the value of VirtualMachineIds for the instance
func (instance *Msvm_CollectionRecoveryPoint) GetPropertyVirtualMachineIds() (value []string, err error) {
	retValue, err := instance.GetProperty("VirtualMachineIds")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
