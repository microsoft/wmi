// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WmiMonitorListedFrequencyRanges struct
type WmiMonitorListedFrequencyRanges struct {
	*MSMonitorClass

	//
	Active bool

	//
	InstanceName string

	//
	MonitorFreqRanges []FrequencyRangeDescriptor

	//
	NumOfMonitorFreqRanges uint16
}

func NewWmiMonitorListedFrequencyRangesEx1(instance *cim.WmiInstance) (newInstance *WmiMonitorListedFrequencyRanges, err error) {
	tmp, err := NewMSMonitorClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WmiMonitorListedFrequencyRanges{
		MSMonitorClass: tmp,
	}
	return
}

func NewWmiMonitorListedFrequencyRangesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WmiMonitorListedFrequencyRanges, err error) {
	tmp, err := NewMSMonitorClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WmiMonitorListedFrequencyRanges{
		MSMonitorClass: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *WmiMonitorListedFrequencyRanges) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *WmiMonitorListedFrequencyRanges) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *WmiMonitorListedFrequencyRanges) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *WmiMonitorListedFrequencyRanges) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetMonitorFreqRanges sets the value of MonitorFreqRanges for the instance
func (instance *WmiMonitorListedFrequencyRanges) SetPropertyMonitorFreqRanges(value []FrequencyRangeDescriptor) (err error) {
	return instance.SetProperty("MonitorFreqRanges", (value))
}

// GetMonitorFreqRanges gets the value of MonitorFreqRanges for the instance
func (instance *WmiMonitorListedFrequencyRanges) GetPropertyMonitorFreqRanges() (value []FrequencyRangeDescriptor, err error) {
	retValue, err := instance.GetProperty("MonitorFreqRanges")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(FrequencyRangeDescriptor)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " FrequencyRangeDescriptor is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, FrequencyRangeDescriptor(valuetmp))
	}

	return
}

// SetNumOfMonitorFreqRanges sets the value of NumOfMonitorFreqRanges for the instance
func (instance *WmiMonitorListedFrequencyRanges) SetPropertyNumOfMonitorFreqRanges(value uint16) (err error) {
	return instance.SetProperty("NumOfMonitorFreqRanges", (value))
}

// GetNumOfMonitorFreqRanges gets the value of NumOfMonitorFreqRanges for the instance
func (instance *WmiMonitorListedFrequencyRanges) GetPropertyNumOfMonitorFreqRanges() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumOfMonitorFreqRanges")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
