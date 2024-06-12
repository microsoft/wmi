// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_AbstractResourcePoolSettingData struct
type Msvm_AbstractResourcePoolSettingData struct {
	*CIM_SettingData

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

func NewMsvm_AbstractResourcePoolSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_AbstractResourcePoolSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_AbstractResourcePoolSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_AbstractResourcePoolSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_AbstractResourcePoolSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_AbstractResourcePoolSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetLoadBalancingBehavior sets the value of LoadBalancingBehavior for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) SetPropertyLoadBalancingBehavior(value uint16) (err error) {
	return instance.SetProperty("LoadBalancingBehavior", (value))
}

// GetLoadBalancingBehavior gets the value of LoadBalancingBehavior for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) GetPropertyLoadBalancingBehavior() (value uint16, err error) {
	retValue, err := instance.GetProperty("LoadBalancingBehavior")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetMappingBehavior sets the value of MappingBehavior for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) SetPropertyMappingBehavior(value uint16) (err error) {
	return instance.SetProperty("MappingBehavior", (value))
}

// GetMappingBehavior gets the value of MappingBehavior for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) GetPropertyMappingBehavior() (value uint16, err error) {
	retValue, err := instance.GetProperty("MappingBehavior")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetMappingOrder sets the value of MappingOrder for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) SetPropertyMappingOrder(value []string) (err error) {
	return instance.SetProperty("MappingOrder", (value))
}

// GetMappingOrder gets the value of MappingOrder for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) GetPropertyMappingOrder() (value []string, err error) {
	retValue, err := instance.GetProperty("MappingOrder")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetNotes sets the value of Notes for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) SetPropertyNotes(value string) (err error) {
	return instance.SetProperty("Notes", (value))
}

// GetNotes gets the value of Notes for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) GetPropertyNotes() (value string, err error) {
	retValue, err := instance.GetProperty("Notes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetOtherResourceType sets the value of OtherResourceType for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) SetPropertyOtherResourceType(value string) (err error) {
	return instance.SetProperty("OtherResourceType", (value))
}

// GetOtherResourceType gets the value of OtherResourceType for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) GetPropertyOtherResourceType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherResourceType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetPoolID sets the value of PoolID for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) SetPropertyPoolID(value string) (err error) {
	return instance.SetProperty("PoolID", (value))
}

// GetPoolID gets the value of PoolID for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) GetPropertyPoolID() (value string, err error) {
	retValue, err := instance.GetProperty("PoolID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetResourceSubType sets the value of ResourceSubType for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) SetPropertyResourceSubType(value string) (err error) {
	return instance.SetProperty("ResourceSubType", (value))
}

// GetResourceSubType gets the value of ResourceSubType for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) GetPropertyResourceSubType() (value string, err error) {
	retValue, err := instance.GetProperty("ResourceSubType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetResourceType sets the value of ResourceType for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) SetPropertyResourceType(value uint16) (err error) {
	return instance.SetProperty("ResourceType", (value))
}

// GetResourceType gets the value of ResourceType for the instance
func (instance *Msvm_AbstractResourcePoolSettingData) GetPropertyResourceType() (value uint16, err error) {
	retValue, err := instance.GetProperty("ResourceType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
func (instance *Msvm_AbstractResourcePoolSettingData) GetRelatedResourcePool() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePool")
}
