// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_VirtualLogicalUnitSettingData struct
type Msvm_VirtualLogicalUnitSettingData struct {
	*CIM_ResourceAllocationSettingData

	//
	StorageSubsystemType string
}

func NewMsvm_VirtualLogicalUnitSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualLogicalUnitSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualLogicalUnitSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_VirtualLogicalUnitSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualLogicalUnitSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualLogicalUnitSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetStorageSubsystemType sets the value of StorageSubsystemType for the instance
func (instance *Msvm_VirtualLogicalUnitSettingData) SetPropertyStorageSubsystemType(value string) (err error) {
	return instance.SetProperty("StorageSubsystemType", (value))
}

// GetStorageSubsystemType gets the value of StorageSubsystemType for the instance
func (instance *Msvm_VirtualLogicalUnitSettingData) GetPropertyStorageSubsystemType() (value string, err error) {
	retValue, err := instance.GetProperty("StorageSubsystemType")
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
func (instance *Msvm_VirtualLogicalUnitSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
