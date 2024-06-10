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

// Msvm_EthernetSwitchPortRdmaSettingData struct
type Msvm_EthernetSwitchPortRdmaSettingData struct {
	*Msvm_EthernetSwitchPortFeatureSettingData

	//
	RdmaOffloadWeight uint32
}

func NewMsvm_EthernetSwitchPortRdmaSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchPortRdmaSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortRdmaSettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchPortRdmaSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchPortRdmaSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortRdmaSettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

// SetRdmaOffloadWeight sets the value of RdmaOffloadWeight for the instance
func (instance *Msvm_EthernetSwitchPortRdmaSettingData) SetPropertyRdmaOffloadWeight(value uint32) (err error) {
	return instance.SetProperty("RdmaOffloadWeight", (value))
}

// GetRdmaOffloadWeight gets the value of RdmaOffloadWeight for the instance
func (instance *Msvm_EthernetSwitchPortRdmaSettingData) GetPropertyRdmaOffloadWeight() (value uint32, err error) {
	retValue, err := instance.GetProperty("RdmaOffloadWeight")
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
func (instance *Msvm_EthernetSwitchPortRdmaSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
