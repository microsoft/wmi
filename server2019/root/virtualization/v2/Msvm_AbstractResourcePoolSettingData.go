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

// Msvm_AbstractResourcePoolSettingData struct
type Msvm_AbstractResourcePoolSettingData struct {
	CIM_SettingData

	//
	LoadBalancingBehavior uint16

	//
	MappingBehavior uint16

	//
	MappingOrder []string

	// End-user supplied notes that are related to this ResourcePool.
	Notes string

	//
	OtherResourceType string

	//
	PoolID string

	//
	ResourceSubType string

	//
	ResourceType uint16
}

// SetLoadBalancingBehavior sets the value of LoadBalancingBehavior for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) SetPropertyLoadBalancingBehavior(value uint16) (err error) {
	return instance.SetProperty("LoadBalancingBehavior", value)
}

// GetLoadBalancingBehavior gets the value of LoadBalancingBehavior for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) GetPropertyLoadBalancingBehavior() (value uint16, err error) {
	retValue, err := instance.GetProperty("LoadBalancingBehavior")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMappingBehavior sets the value of MappingBehavior for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) SetPropertyMappingBehavior(value uint16) (err error) {
	return instance.SetProperty("MappingBehavior", value)
}

// GetMappingBehavior gets the value of MappingBehavior for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) GetPropertyMappingBehavior() (value uint16, err error) {
	retValue, err := instance.GetProperty("MappingBehavior")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMappingOrder sets the value of MappingOrder for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) SetPropertyMappingOrder(value []string) (err error) {
	return instance.SetProperty("MappingOrder", value)
}

// GetMappingOrder gets the value of MappingOrder for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) GetPropertyMappingOrder() (value []string, err error) {
	retValue, err := instance.GetProperty("MappingOrder")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNotes sets the value of Notes for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) SetPropertyNotes(value string) (err error) {
	return instance.SetProperty("Notes", value)
}

// GetNotes gets the value of Notes for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) GetPropertyNotes() (value string, err error) {
	retValue, err := instance.GetProperty("Notes")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherResourceType sets the value of OtherResourceType for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) SetPropertyOtherResourceType(value string) (err error) {
	return instance.SetProperty("OtherResourceType", value)
}

// GetOtherResourceType gets the value of OtherResourceType for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) GetPropertyOtherResourceType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherResourceType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPoolID sets the value of PoolID for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) SetPropertyPoolID(value string) (err error) {
	return instance.SetProperty("PoolID", value)
}

// GetPoolID gets the value of PoolID for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) GetPropertyPoolID() (value string, err error) {
	retValue, err := instance.GetProperty("PoolID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceSubType sets the value of ResourceSubType for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) SetPropertyResourceSubType(value string) (err error) {
	return instance.SetProperty("ResourceSubType", value)
}

// GetResourceSubType gets the value of ResourceSubType for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) GetPropertyResourceSubType() (value string, err error) {
	retValue, err := instance.GetProperty("ResourceSubType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceType sets the value of ResourceType for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) SetPropertyResourceType(value uint16) (err error) {
	return instance.SetProperty("ResourceType", value)
}

// GetResourceType gets the value of ResourceType for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) GetPropertyResourceType() (value uint16, err error) {
	retValue, err := instance.GetProperty("ResourceType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_AbstractResourcePoolSettingData) GetRelatedResourcePool() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePool")
}
