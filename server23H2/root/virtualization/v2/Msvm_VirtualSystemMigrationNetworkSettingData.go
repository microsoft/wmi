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

// Msvm_VirtualSystemMigrationNetworkSettingData struct
type Msvm_VirtualSystemMigrationNetworkSettingData struct {
	*CIM_SettingData

	//
	Metric uint32

	//
	PrefixLength uint8

	//
	SubnetNumber string

	//
	Tags []string
}

func NewMsvm_VirtualSystemMigrationNetworkSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemMigrationNetworkSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemMigrationNetworkSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_VirtualSystemMigrationNetworkSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemMigrationNetworkSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemMigrationNetworkSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetMetric sets the value of Metric for the instance
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) SetPropertyMetric(value uint32) (err error) {
	return instance.SetProperty("Metric", (value))
}

// GetMetric gets the value of Metric for the instance
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) GetPropertyMetric() (value uint32, err error) {
	retValue, err := instance.GetProperty("Metric")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPrefixLength sets the value of PrefixLength for the instance
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) SetPropertyPrefixLength(value uint8) (err error) {
	return instance.SetProperty("PrefixLength", (value))
}

// GetPrefixLength gets the value of PrefixLength for the instance
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) GetPropertyPrefixLength() (value uint8, err error) {
	retValue, err := instance.GetProperty("PrefixLength")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSubnetNumber sets the value of SubnetNumber for the instance
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) SetPropertySubnetNumber(value string) (err error) {
	return instance.SetProperty("SubnetNumber", (value))
}

// GetSubnetNumber gets the value of SubnetNumber for the instance
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) GetPropertySubnetNumber() (value string, err error) {
	retValue, err := instance.GetProperty("SubnetNumber")
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

// SetTags sets the value of Tags for the instance
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) SetPropertyTags(value []string) (err error) {
	return instance.SetProperty("Tags", (value))
}

// GetTags gets the value of Tags for the instance
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) GetPropertyTags() (value []string, err error) {
	retValue, err := instance.GetProperty("Tags")
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
func (instance *Msvm_VirtualSystemMigrationNetworkSettingData) GetRelatedVirtualSystemMigrationServiceSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemMigrationServiceSettingData")
}
