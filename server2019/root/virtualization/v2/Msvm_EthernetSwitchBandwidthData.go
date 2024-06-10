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

// Msvm_EthernetSwitchBandwidthData struct
type Msvm_EthernetSwitchBandwidthData struct {
	*Msvm_EthernetSwitchData

	//
	Capacity uint64

	//
	DefaultFlowReservation uint64

	//
	DefaultFlowReservationPercentage uint32

	//
	DefaultFlowWeight uint64

	//
	Reservation uint64
}

func NewMsvm_EthernetSwitchBandwidthDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchBandwidthData, err error) {
	tmp, err := NewMsvm_EthernetSwitchDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchBandwidthData{
		Msvm_EthernetSwitchData: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchBandwidthDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchBandwidthData, err error) {
	tmp, err := NewMsvm_EthernetSwitchDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchBandwidthData{
		Msvm_EthernetSwitchData: tmp,
	}
	return
}

// SetCapacity sets the value of Capacity for the instance
func (instance *Msvm_EthernetSwitchBandwidthData) SetPropertyCapacity(value uint64) (err error) {
	return instance.SetProperty("Capacity", (value))
}

// GetCapacity gets the value of Capacity for the instance
func (instance *Msvm_EthernetSwitchBandwidthData) GetPropertyCapacity() (value uint64, err error) {
	retValue, err := instance.GetProperty("Capacity")
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

// SetDefaultFlowReservation sets the value of DefaultFlowReservation for the instance
func (instance *Msvm_EthernetSwitchBandwidthData) SetPropertyDefaultFlowReservation(value uint64) (err error) {
	return instance.SetProperty("DefaultFlowReservation", (value))
}

// GetDefaultFlowReservation gets the value of DefaultFlowReservation for the instance
func (instance *Msvm_EthernetSwitchBandwidthData) GetPropertyDefaultFlowReservation() (value uint64, err error) {
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

// SetDefaultFlowReservationPercentage sets the value of DefaultFlowReservationPercentage for the instance
func (instance *Msvm_EthernetSwitchBandwidthData) SetPropertyDefaultFlowReservationPercentage(value uint32) (err error) {
	return instance.SetProperty("DefaultFlowReservationPercentage", (value))
}

// GetDefaultFlowReservationPercentage gets the value of DefaultFlowReservationPercentage for the instance
func (instance *Msvm_EthernetSwitchBandwidthData) GetPropertyDefaultFlowReservationPercentage() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultFlowReservationPercentage")
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

// SetDefaultFlowWeight sets the value of DefaultFlowWeight for the instance
func (instance *Msvm_EthernetSwitchBandwidthData) SetPropertyDefaultFlowWeight(value uint64) (err error) {
	return instance.SetProperty("DefaultFlowWeight", (value))
}

// GetDefaultFlowWeight gets the value of DefaultFlowWeight for the instance
func (instance *Msvm_EthernetSwitchBandwidthData) GetPropertyDefaultFlowWeight() (value uint64, err error) {
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

// SetReservation sets the value of Reservation for the instance
func (instance *Msvm_EthernetSwitchBandwidthData) SetPropertyReservation(value uint64) (err error) {
	return instance.SetProperty("Reservation", (value))
}

// GetReservation gets the value of Reservation for the instance
func (instance *Msvm_EthernetSwitchBandwidthData) GetPropertyReservation() (value uint64, err error) {
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
func (instance *Msvm_EthernetSwitchBandwidthData) GetRelatedVirtualEthernetSwitch() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitch")
}
