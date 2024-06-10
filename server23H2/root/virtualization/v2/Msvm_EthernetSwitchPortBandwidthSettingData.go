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

// Msvm_EthernetSwitchPortBandwidthSettingData struct
type Msvm_EthernetSwitchPortBandwidthSettingData struct {
	*Msvm_EthernetSwitchPortFeatureSettingData

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

func NewMsvm_EthernetSwitchPortBandwidthSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchPortBandwidthSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortBandwidthSettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchPortBandwidthSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchPortBandwidthSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortBandwidthSettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

// SetBurstLimit sets the value of BurstLimit for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) SetPropertyBurstLimit(value uint64) (err error) {
	return instance.SetProperty("BurstLimit", (value))
}

// GetBurstLimit gets the value of BurstLimit for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) GetPropertyBurstLimit() (value uint64, err error) {
	retValue, err := instance.GetProperty("BurstLimit")
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

// SetBurstSize sets the value of BurstSize for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) SetPropertyBurstSize(value uint64) (err error) {
	return instance.SetProperty("BurstSize", (value))
}

// GetBurstSize gets the value of BurstSize for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) GetPropertyBurstSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("BurstSize")
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

// SetLimit sets the value of Limit for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) SetPropertyLimit(value uint64) (err error) {
	return instance.SetProperty("Limit", (value))
}

// GetLimit gets the value of Limit for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) GetPropertyLimit() (value uint64, err error) {
	retValue, err := instance.GetProperty("Limit")
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

// SetReservation sets the value of Reservation for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) SetPropertyReservation(value uint64) (err error) {
	return instance.SetProperty("Reservation", (value))
}

// GetReservation gets the value of Reservation for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) GetPropertyReservation() (value uint64, err error) {
	retValue, err := instance.GetProperty("Reservation")
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

// SetWeight sets the value of Weight for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) SetPropertyWeight(value uint64) (err error) {
	return instance.SetProperty("Weight", (value))
}

// GetWeight gets the value of Weight for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) GetPropertyWeight() (value uint64, err error) {
	retValue, err := instance.GetProperty("Weight")
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
func (instance *Msvm_EthernetSwitchPortBandwidthSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
