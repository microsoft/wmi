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

// Msvm_VirtualSystemMigrationNetworkSettingData struct
type Msvm_VirtualSystemMigrationNetworkSettingData struct {
	CIM_SettingData

	//
	Metric uint32

	//
	PrefixLength uint8

	//
	SubnetNumber string

	//
	Tags []string
}

// SetMetric sets the value of Metric for the instance
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) SetPropertyMetric(value uint32) (err error) {
	return instance.SetProperty("Metric", value)
}

// GetMetric gets the value of Metric for the instance
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) GetPropertyMetric() (value uint32, err error) {
	retValue, err := instance.GetProperty("Metric")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrefixLength sets the value of PrefixLength for the instance
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) SetPropertyPrefixLength(value uint8) (err error) {
	return instance.SetProperty("PrefixLength", value)
}

// GetPrefixLength gets the value of PrefixLength for the instance
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) GetPropertyPrefixLength() (value uint8, err error) {
	retValue, err := instance.GetProperty("PrefixLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSubnetNumber sets the value of SubnetNumber for the instance
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) SetPropertySubnetNumber(value string) (err error) {
	return instance.SetProperty("SubnetNumber", value)
}

// GetSubnetNumber gets the value of SubnetNumber for the instance
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) GetPropertySubnetNumber() (value string, err error) {
	retValue, err := instance.GetProperty("SubnetNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTags sets the value of Tags for the instance
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) SetPropertyTags(value []string) (err error) {
	return instance.SetProperty("Tags", value)
}

// GetTags gets the value of Tags for the instance
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) GetPropertyTags() (value []string, err error) {
	retValue, err := instance.GetProperty("Tags")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) GetRelatedVirtualSystemMigrationServiceSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemMigrationServiceSettingData")
}
