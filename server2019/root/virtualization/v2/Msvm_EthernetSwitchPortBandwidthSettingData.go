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

// Msvm_EthernetSwitchPortBandwidthSettingData struct
type Msvm_EthernetSwitchPortBandwidthSettingData struct {
	Msvm_EthernetSwitchPortFeatureSettingData

	//
	BurstLimit uint64

	//
	BurstSize uint64

	//
	Limit uint64

	//
	Reservation uint64

	//
	Weight uint64
}

// SetBurstLimit sets the value of BurstLimit for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) SetPropertyBurstLimit(value uint64) (err error) {
	return instance.SetProperty("BurstLimit", value)
}

// GetBurstLimit gets the value of BurstLimit for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) GetPropertyBurstLimit() (value uint64, err error) {
	retValue, err := instance.GetProperty("BurstLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBurstSize sets the value of BurstSize for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) SetPropertyBurstSize(value uint64) (err error) {
	return instance.SetProperty("BurstSize", value)
}

// GetBurstSize gets the value of BurstSize for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) GetPropertyBurstSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("BurstSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLimit sets the value of Limit for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) SetPropertyLimit(value uint64) (err error) {
	return instance.SetProperty("Limit", value)
}

// GetLimit gets the value of Limit for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) GetPropertyLimit() (value uint64, err error) {
	retValue, err := instance.GetProperty("Limit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReservation sets the value of Reservation for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) SetPropertyReservation(value uint64) (err error) {
	return instance.SetProperty("Reservation", value)
}

// GetReservation gets the value of Reservation for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) GetPropertyReservation() (value uint64, err error) {
	retValue, err := instance.GetProperty("Reservation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWeight sets the value of Weight for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) SetPropertyWeight(value uint64) (err error) {
	return instance.SetProperty("Weight", value)
}

// GetWeight gets the value of Weight for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) GetPropertyWeight() (value uint64, err error) {
	retValue, err := instance.GetProperty("Weight")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
