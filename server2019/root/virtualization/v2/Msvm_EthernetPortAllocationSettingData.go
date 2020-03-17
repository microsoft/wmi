// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_EthernetPortAllocationSettingData struct
type Msvm_EthernetPortAllocationSettingData struct {
	CIM_EthernetPortAllocationSettingData

	//
	CompartmentGuid string

	// EnabledState is an integer enumeration that indicates whether the allocation request is enabled or disabled. When an allocation request is marked as Disabled (3), then the allocation is not processed. The EnabledState for an active configuration is always marked as Enabled (2).
	EnabledState EthernetPortAllocationSettingData_EnabledState

	// The last known friendly name of the switch this port had a hard affinity to, if any.
	LastKnownSwitchName string

	//
	RequiredFeatureHints []string

	//
	RequiredFeatures []string

	//
	TestReplicaPoolID string

	//
	TestReplicaSwitchName string
}

// SetCompartmentGuid sets the value of CompartmentGuid for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) SetPropertyCompartmentGuid(value string) (err error) {
	return instance.SetProperty("CompartmentGuid", value)
}

// GetCompartmentGuid gets the value of CompartmentGuid for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) GetPropertyCompartmentGuid() (value string, err error) {
	retValue, err := instance.GetProperty("CompartmentGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabledState sets the value of EnabledState for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) SetPropertyEnabledState(value EthernetPortAllocationSettingData_EnabledState) (err error) {
	return instance.SetProperty("EnabledState", value)
}

// GetEnabledState gets the value of EnabledState for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) GetPropertyEnabledState() (value EthernetPortAllocationSettingData_EnabledState, err error) {
	retValue, err := instance.GetProperty("EnabledState")
	if err != nil {
		return
	}
	value, ok := retValue.(EthernetPortAllocationSettingData_EnabledState)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastKnownSwitchName sets the value of LastKnownSwitchName for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) SetPropertyLastKnownSwitchName(value string) (err error) {
	return instance.SetProperty("LastKnownSwitchName", value)
}

// GetLastKnownSwitchName gets the value of LastKnownSwitchName for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) GetPropertyLastKnownSwitchName() (value string, err error) {
	retValue, err := instance.GetProperty("LastKnownSwitchName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequiredFeatureHints sets the value of RequiredFeatureHints for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) SetPropertyRequiredFeatureHints(value []string) (err error) {
	return instance.SetProperty("RequiredFeatureHints", value)
}

// GetRequiredFeatureHints gets the value of RequiredFeatureHints for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) GetPropertyRequiredFeatureHints() (value []string, err error) {
	retValue, err := instance.GetProperty("RequiredFeatureHints")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequiredFeatures sets the value of RequiredFeatures for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) SetPropertyRequiredFeatures(value []string) (err error) {
	return instance.SetProperty("RequiredFeatures", value)
}

// GetRequiredFeatures gets the value of RequiredFeatures for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) GetPropertyRequiredFeatures() (value []string, err error) {
	retValue, err := instance.GetProperty("RequiredFeatures")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTestReplicaPoolID sets the value of TestReplicaPoolID for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) SetPropertyTestReplicaPoolID(value string) (err error) {
	return instance.SetProperty("TestReplicaPoolID", value)
}

// GetTestReplicaPoolID gets the value of TestReplicaPoolID for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) GetPropertyTestReplicaPoolID() (value string, err error) {
	retValue, err := instance.GetProperty("TestReplicaPoolID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTestReplicaSwitchName sets the value of TestReplicaSwitchName for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) SetPropertyTestReplicaSwitchName(value string) (err error) {
	return instance.SetProperty("TestReplicaSwitchName", value)
}

// GetTestReplicaSwitchName gets the value of TestReplicaSwitchName for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) GetPropertyTestReplicaSwitchName() (value string, err error) {
	retValue, err := instance.GetProperty("TestReplicaSwitchName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_EthernetPortAllocationSettingData) GetRelatedVirtualEthernetSwitchSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitchSettingData")
}

func (instance *Msvm_EthernetPortAllocationSettingData) GetRelatedEthernetSwitchPort() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchPort")
}
