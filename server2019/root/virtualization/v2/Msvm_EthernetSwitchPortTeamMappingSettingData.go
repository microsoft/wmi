// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
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
	return instance.SetProperty("DisableOnFailover", value)
}

// GetDisableOnFailover gets the value of DisableOnFailover for the instance
func (instance *Msvm_EthernetSwitchPortTeamMappingSettingData) GetPropertyDisableOnFailover() (value uint32, err error) {
	retValue, err := instance.GetProperty("DisableOnFailover")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetAdapterDeviceId sets the value of NetAdapterDeviceId for the instance
func (instance *Msvm_EthernetSwitchPortTeamMappingSettingData) SetPropertyNetAdapterDeviceId(value string) (err error) {
	return instance.SetProperty("NetAdapterDeviceId", value)
}

// GetNetAdapterDeviceId gets the value of NetAdapterDeviceId for the instance
func (instance *Msvm_EthernetSwitchPortTeamMappingSettingData) GetPropertyNetAdapterDeviceId() (value string, err error) {
	retValue, err := instance.GetProperty("NetAdapterDeviceId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetAdapterName sets the value of NetAdapterName for the instance
func (instance *Msvm_EthernetSwitchPortTeamMappingSettingData) SetPropertyNetAdapterName(value string) (err error) {
	return instance.SetProperty("NetAdapterName", value)
}

// GetNetAdapterName gets the value of NetAdapterName for the instance
func (instance *Msvm_EthernetSwitchPortTeamMappingSettingData) GetPropertyNetAdapterName() (value string, err error) {
	retValue, err := instance.GetProperty("NetAdapterName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_EthernetSwitchPortTeamMappingSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
