// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_ReplicationAuthorizationSettingData struct
type Msvm_ReplicationAuthorizationSettingData struct {
	CIM_SettingData

	//
	AllowedPrimaryHostSystem string

	//
	ReplicaStorageLocation string

	//
	TrustGroup string
}

// SetAllowedPrimaryHostSystem sets the value of AllowedPrimaryHostSystem for the instance
func (instance *Msvm_ReplicationAuthorizationSettingData) SetPropertyAllowedPrimaryHostSystem(value string) (err error) {
	return instance.SetProperty("AllowedPrimaryHostSystem", value)
}

// GetAllowedPrimaryHostSystem gets the value of AllowedPrimaryHostSystem for the instance
func (instance *Msvm_ReplicationAuthorizationSettingData) GetPropertyAllowedPrimaryHostSystem() (value string, err error) {
	retValue, err := instance.GetProperty("AllowedPrimaryHostSystem")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplicaStorageLocation sets the value of ReplicaStorageLocation for the instance
func (instance *Msvm_ReplicationAuthorizationSettingData) SetPropertyReplicaStorageLocation(value string) (err error) {
	return instance.SetProperty("ReplicaStorageLocation", value)
}

// GetReplicaStorageLocation gets the value of ReplicaStorageLocation for the instance
func (instance *Msvm_ReplicationAuthorizationSettingData) GetPropertyReplicaStorageLocation() (value string, err error) {
	retValue, err := instance.GetProperty("ReplicaStorageLocation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustGroup sets the value of TrustGroup for the instance
func (instance *Msvm_ReplicationAuthorizationSettingData) SetPropertyTrustGroup(value string) (err error) {
	return instance.SetProperty("TrustGroup", value)
}

// GetTrustGroup gets the value of TrustGroup for the instance
func (instance *Msvm_ReplicationAuthorizationSettingData) GetPropertyTrustGroup() (value string, err error) {
	retValue, err := instance.GetProperty("TrustGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
