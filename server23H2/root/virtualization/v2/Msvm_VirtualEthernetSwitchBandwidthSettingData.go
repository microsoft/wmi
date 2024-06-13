// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_VirtualEthernetSwitchBandwidthSettingData struct
type Msvm_VirtualEthernetSwitchBandwidthSettingData struct {
	*Msvm_EthernetSwitchFeatureSettingData

	//
	DefaultFlowReservation uint64

	//
	DefaultFlowWeight uint64
}

func NewMsvm_VirtualEthernetSwitchBandwidthSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualEthernetSwitchBandwidthSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchFeatureSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualEthernetSwitchBandwidthSettingData{
		Msvm_EthernetSwitchFeatureSettingData: tmp,
	}
	return
}

func NewMsvm_VirtualEthernetSwitchBandwidthSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualEthernetSwitchBandwidthSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchFeatureSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualEthernetSwitchBandwidthSettingData{
		Msvm_EthernetSwitchFeatureSettingData: tmp,
	}
	return
}

// SetDefaultFlowReservation sets the value of DefaultFlowReservation for the instance
func (instance *Msvm_VirtualEthernetSwitchBandwidthSettingData) SetPropertyDefaultFlowReservation(value uint64) (err error) {
	return instance.SetProperty("DefaultFlowReservation", (value))
}

// GetDefaultFlowReservation gets the value of DefaultFlowReservation for the instance
func (instance *Msvm_VirtualEthernetSwitchBandwidthSettingData) GetPropertyDefaultFlowReservation() (value uint64, err error) {
	retValue, err := instance.GetProperty("DefaultFlowReservation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetDefaultFlowWeight sets the value of DefaultFlowWeight for the instance
func (instance *Msvm_VirtualEthernetSwitchBandwidthSettingData) SetPropertyDefaultFlowWeight(value uint64) (err error) {
	return instance.SetProperty("DefaultFlowWeight", (value))
}

// GetDefaultFlowWeight gets the value of DefaultFlowWeight for the instance
func (instance *Msvm_VirtualEthernetSwitchBandwidthSettingData) GetPropertyDefaultFlowWeight() (value uint64, err error) {
	retValue, err := instance.GetProperty("DefaultFlowWeight")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
func (instance *Msvm_VirtualEthernetSwitchBandwidthSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
