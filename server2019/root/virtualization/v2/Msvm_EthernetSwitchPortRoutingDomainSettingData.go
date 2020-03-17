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

// Msvm_EthernetSwitchPortRoutingDomainSettingData struct
type Msvm_EthernetSwitchPortRoutingDomainSettingData struct {
	Msvm_EthernetSwitchPortFeatureSettingData

	//
	IsolationIdList []uint32

	//
	IsolationIdNameList []string

	//
	RoutingDomainGuid string

	//
	RoutingDomainName string
}

// SetIsolationIdList sets the value of IsolationIdList for the instance
func (instance *Msvm_EthernetSwitchPortRoutingDomainSettingData) SetPropertyIsolationIdList(value []uint32) (err error) {
	return instance.SetProperty("IsolationIdList", value)
}

// GetIsolationIdList gets the value of IsolationIdList for the instance
func (instance *Msvm_EthernetSwitchPortRoutingDomainSettingData) GetPropertyIsolationIdList() (value []uint32, err error) {
	retValue, err := instance.GetProperty("IsolationIdList")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsolationIdNameList sets the value of IsolationIdNameList for the instance
func (instance *Msvm_EthernetSwitchPortRoutingDomainSettingData) SetPropertyIsolationIdNameList(value []string) (err error) {
	return instance.SetProperty("IsolationIdNameList", value)
}

// GetIsolationIdNameList gets the value of IsolationIdNameList for the instance
func (instance *Msvm_EthernetSwitchPortRoutingDomainSettingData) GetPropertyIsolationIdNameList() (value []string, err error) {
	retValue, err := instance.GetProperty("IsolationIdNameList")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoutingDomainGuid sets the value of RoutingDomainGuid for the instance
func (instance *Msvm_EthernetSwitchPortRoutingDomainSettingData) SetPropertyRoutingDomainGuid(value string) (err error) {
	return instance.SetProperty("RoutingDomainGuid", value)
}

// GetRoutingDomainGuid gets the value of RoutingDomainGuid for the instance
func (instance *Msvm_EthernetSwitchPortRoutingDomainSettingData) GetPropertyRoutingDomainGuid() (value string, err error) {
	retValue, err := instance.GetProperty("RoutingDomainGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoutingDomainName sets the value of RoutingDomainName for the instance
func (instance *Msvm_EthernetSwitchPortRoutingDomainSettingData) SetPropertyRoutingDomainName(value string) (err error) {
	return instance.SetProperty("RoutingDomainName", value)
}

// GetRoutingDomainName gets the value of RoutingDomainName for the instance
func (instance *Msvm_EthernetSwitchPortRoutingDomainSettingData) GetPropertyRoutingDomainName() (value string, err error) {
	retValue, err := instance.GetProperty("RoutingDomainName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_EthernetSwitchPortRoutingDomainSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
