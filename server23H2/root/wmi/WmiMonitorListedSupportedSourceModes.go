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

// WmiMonitorListedSupportedSourceModes struct
type WmiMonitorListedSupportedSourceModes struct {
	*MSMonitorClass

	//
	Active bool

	//
	InstanceName string

	//
	MonitorSourceModes []VideoModeDescriptor

	//
	NumOfMonitorSourceModes uint16

	//
	PreferredMonitorSourceModeIndex uint16
}

func NewWmiMonitorListedSupportedSourceModesEx1(instance *cim.WmiInstance) (newInstance *WmiMonitorListedSupportedSourceModes, err error) {
	tmp, err := NewMSMonitorClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WmiMonitorListedSupportedSourceModes{
		MSMonitorClass: tmp,
	}
	return
}

func NewWmiMonitorListedSupportedSourceModesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WmiMonitorListedSupportedSourceModes, err error) {
	tmp, err := NewMSMonitorClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WmiMonitorListedSupportedSourceModes{
		MSMonitorClass: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *WmiMonitorListedSupportedSourceModes) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *WmiMonitorListedSupportedSourceModes) GetPropertyActive() (value bool, err error) {
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
func (instance *WmiMonitorListedSupportedSourceModes) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *WmiMonitorListedSupportedSourceModes) GetPropertyInstanceName() (value string, err error) {
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

// SetMonitorSourceModes sets the value of MonitorSourceModes for the instance
func (instance *WmiMonitorListedSupportedSourceModes) SetPropertyMonitorSourceModes(value []VideoModeDescriptor) (err error) {
	return instance.SetProperty("MonitorSourceModes", (value))
}

// GetMonitorSourceModes gets the value of MonitorSourceModes for the instance
func (instance *WmiMonitorListedSupportedSourceModes) GetPropertyMonitorSourceModes() (value []VideoModeDescriptor, err error) {
	retValue, err := instance.GetProperty("MonitorSourceModes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(VideoModeDescriptor)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " VideoModeDescriptor is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, VideoModeDescriptor(valuetmp))
	}

	return
}

// SetNumOfMonitorSourceModes sets the value of NumOfMonitorSourceModes for the instance
func (instance *WmiMonitorListedSupportedSourceModes) SetPropertyNumOfMonitorSourceModes(value uint16) (err error) {
	return instance.SetProperty("NumOfMonitorSourceModes", (value))
}

// GetNumOfMonitorSourceModes gets the value of NumOfMonitorSourceModes for the instance
func (instance *WmiMonitorListedSupportedSourceModes) GetPropertyNumOfMonitorSourceModes() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumOfMonitorSourceModes")
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

// SetPreferredMonitorSourceModeIndex sets the value of PreferredMonitorSourceModeIndex for the instance
func (instance *WmiMonitorListedSupportedSourceModes) SetPropertyPreferredMonitorSourceModeIndex(value uint16) (err error) {
	return instance.SetProperty("PreferredMonitorSourceModeIndex", (value))
}

// GetPreferredMonitorSourceModeIndex gets the value of PreferredMonitorSourceModeIndex for the instance
func (instance *WmiMonitorListedSupportedSourceModes) GetPropertyPreferredMonitorSourceModeIndex() (value uint16, err error) {
	retValue, err := instance.GetProperty("PreferredMonitorSourceModeIndex")
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
