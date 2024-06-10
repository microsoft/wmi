// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_EthernetSwitchPortTeamMappingSettingData struct
type Msvm_EthernetSwitchPortTeamMappingSettingData struct {
	*Msvm_EthernetSwitchPortFeatureSettingData

	//
	DisableOnFailover uint32

	//
	NetAdapterDeviceId string

	//
	NetAdapterName string
}

func NewMsvm_EthernetSwitchPortTeamMappingSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchPortTeamMappingSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortTeamMappingSettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchPortTeamMappingSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchPortTeamMappingSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortTeamMappingSettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

// SetDisableOnFailover sets the value of DisableOnFailover for the instance
func (instance *Msvm_EthernetSwitchPortTeamMappingSettingData) SetPropertyDisableOnFailover(value uint32) (err error) {
	return instance.SetProperty("DisableOnFailover", (value))
}

// GetDisableOnFailover gets the value of DisableOnFailover for the instance
func (instance *Msvm_EthernetSwitchPortTeamMappingSettingData) GetPropertyDisableOnFailover() (value uint32, err error) {
	retValue, err := instance.GetProperty("DisableOnFailover")
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

// SetNetAdapterDeviceId sets the value of NetAdapterDeviceId for the instance
func (instance *Msvm_EthernetSwitchPortTeamMappingSettingData) SetPropertyNetAdapterDeviceId(value string) (err error) {
	return instance.SetProperty("NetAdapterDeviceId", (value))
}

// GetNetAdapterDeviceId gets the value of NetAdapterDeviceId for the instance
func (instance *Msvm_EthernetSwitchPortTeamMappingSettingData) GetPropertyNetAdapterDeviceId() (value string, err error) {
	retValue, err := instance.GetProperty("NetAdapterDeviceId")
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

// SetNetAdapterName sets the value of NetAdapterName for the instance
func (instance *Msvm_EthernetSwitchPortTeamMappingSettingData) SetPropertyNetAdapterName(value string) (err error) {
	return instance.SetProperty("NetAdapterName", (value))
}

// GetNetAdapterName gets the value of NetAdapterName for the instance
func (instance *Msvm_EthernetSwitchPortTeamMappingSettingData) GetPropertyNetAdapterName() (value string, err error) {
	retValue, err := instance.GetProperty("NetAdapterName")
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
func (instance *Msvm_EthernetSwitchPortTeamMappingSettingData) GetRelatedEthernetPortAllocationSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetPortAllocationSettingData")
}
