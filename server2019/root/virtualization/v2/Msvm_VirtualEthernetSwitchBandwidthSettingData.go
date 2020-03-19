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
	return instance.SetProperty("DefaultFlowReservation", value)
}

// GetDefaultFlowReservation gets the value of DefaultFlowReservation for the instance
func (instance *Msvm_VirtualEthernetSwitchBandwidthSettingData) GetPropertyDefaultFlowReservation() (value uint64, err error) {
	retValue, err := instance.GetProperty("DefaultFlowReservation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultFlowWeight sets the value of DefaultFlowWeight for the instance
func (instance *Msvm_VirtualEthernetSwitchBandwidthSettingData) SetPropertyDefaultFlowWeight(value uint64) (err error) {
	return instance.SetProperty("DefaultFlowWeight", value)
}

// GetDefaultFlowWeight gets the value of DefaultFlowWeight for the instance
func (instance *Msvm_VirtualEthernetSwitchBandwidthSettingData) GetPropertyDefaultFlowWeight() (value uint64, err error) {
	retValue, err := instance.GetProperty("DefaultFlowWeight")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_VirtualEthernetSwitchBandwidthSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
