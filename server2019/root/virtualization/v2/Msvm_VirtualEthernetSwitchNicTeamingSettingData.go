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

// Msvm_VirtualEthernetSwitchNicTeamingSettingData struct
type Msvm_VirtualEthernetSwitchNicTeamingSettingData struct {
	Msvm_EthernetSwitchFeatureSettingData

	//
	LoadBalancingAlgorithm uint32

	//
	TeamingMode uint32
}

// SetLoadBalancingAlgorithm sets the value of LoadBalancingAlgorithm for the instance
func (instance *Msvm_VirtualEthernetSwitchNicTeamingSettingData) SetPropertyLoadBalancingAlgorithm(value uint32) (err error) {
	return instance.SetProperty("LoadBalancingAlgorithm", value)
}

// GetLoadBalancingAlgorithm gets the value of LoadBalancingAlgorithm for the instance
func (instance *Msvm_VirtualEthernetSwitchNicTeamingSettingData) GetPropertyLoadBalancingAlgorithm() (value uint32, err error) {
	retValue, err := instance.GetProperty("LoadBalancingAlgorithm")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTeamingMode sets the value of TeamingMode for the instance
func (instance *Msvm_VirtualEthernetSwitchNicTeamingSettingData) SetPropertyTeamingMode(value uint32) (err error) {
	return instance.SetProperty("TeamingMode", value)
}

// GetTeamingMode gets the value of TeamingMode for the instance
func (instance *Msvm_VirtualEthernetSwitchNicTeamingSettingData) GetPropertyTeamingMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("TeamingMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_VirtualEthernetSwitchNicTeamingSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
