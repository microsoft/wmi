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

// Msvm_VirtualEthernetSwitchNicTeamingSettingData struct
type Msvm_VirtualEthernetSwitchNicTeamingSettingData struct {
	*Msvm_EthernetSwitchFeatureSettingData

	//
	LoadBalancingAlgorithm uint32

	//
	TeamingMode uint32
}

func NewMsvm_VirtualEthernetSwitchNicTeamingSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualEthernetSwitchNicTeamingSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchFeatureSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualEthernetSwitchNicTeamingSettingData{
		Msvm_EthernetSwitchFeatureSettingData: tmp,
	}
	return
}

func NewMsvm_VirtualEthernetSwitchNicTeamingSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualEthernetSwitchNicTeamingSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchFeatureSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualEthernetSwitchNicTeamingSettingData{
		Msvm_EthernetSwitchFeatureSettingData: tmp,
	}
	return
}

// SetLoadBalancingAlgorithm sets the value of LoadBalancingAlgorithm for the instance
func (instance *Msvm_VirtualEthernetSwitchNicTeamingSettingData) SetPropertyLoadBalancingAlgorithm(value uint32) (err error) {
	return instance.SetProperty("LoadBalancingAlgorithm", (value))
}

// GetLoadBalancingAlgorithm gets the value of LoadBalancingAlgorithm for the instance
func (instance *Msvm_VirtualEthernetSwitchNicTeamingSettingData) GetPropertyLoadBalancingAlgorithm() (value uint32, err error) {
	retValue, err := instance.GetProperty("LoadBalancingAlgorithm")
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

// SetTeamingMode sets the value of TeamingMode for the instance
func (instance *Msvm_VirtualEthernetSwitchNicTeamingSettingData) SetPropertyTeamingMode(value uint32) (err error) {
	return instance.SetProperty("TeamingMode", (value))
}

// GetTeamingMode gets the value of TeamingMode for the instance
func (instance *Msvm_VirtualEthernetSwitchNicTeamingSettingData) GetPropertyTeamingMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("TeamingMode")
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
func (instance *Msvm_VirtualEthernetSwitchNicTeamingSettingData) GetRelatedVirtualEthernetSwitchSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitchSettingData")
}
