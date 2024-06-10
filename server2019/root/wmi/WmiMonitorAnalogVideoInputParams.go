// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WmiMonitorAnalogVideoInputParams struct
type WmiMonitorAnalogVideoInputParams struct {
	*MSMonitorClass

	//
	Active bool

	//
	CompositeSyncSupported uint8

	//
	InstanceName string

	//
	SeparateSyncsSupported uint8

	//
	SerrationOfVsyncRequired uint8

	//
	SetupExpected uint8

	//
	SignalLevelStandard uint8

	//
	SyncOnGreenVideoSupported uint8
}

func NewWmiMonitorAnalogVideoInputParamsEx1(instance *cim.WmiInstance) (newInstance *WmiMonitorAnalogVideoInputParams, err error) {
	tmp, err := NewMSMonitorClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WmiMonitorAnalogVideoInputParams{
		MSMonitorClass: tmp,
	}
	return
}

func NewWmiMonitorAnalogVideoInputParamsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WmiMonitorAnalogVideoInputParams, err error) {
	tmp, err := NewMSMonitorClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WmiMonitorAnalogVideoInputParams{
		MSMonitorClass: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *WmiMonitorAnalogVideoInputParams) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *WmiMonitorAnalogVideoInputParams) GetPropertyActive() (value bool, err error) {
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

// SetCompositeSyncSupported sets the value of CompositeSyncSupported for the instance
func (instance *WmiMonitorAnalogVideoInputParams) SetPropertyCompositeSyncSupported(value uint8) (err error) {
	return instance.SetProperty("CompositeSyncSupported", (value))
}

// GetCompositeSyncSupported gets the value of CompositeSyncSupported for the instance
func (instance *WmiMonitorAnalogVideoInputParams) GetPropertyCompositeSyncSupported() (value uint8, err error) {
	retValue, err := instance.GetProperty("CompositeSyncSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *WmiMonitorAnalogVideoInputParams) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *WmiMonitorAnalogVideoInputParams) GetPropertyInstanceName() (value string, err error) {
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

// SetSeparateSyncsSupported sets the value of SeparateSyncsSupported for the instance
func (instance *WmiMonitorAnalogVideoInputParams) SetPropertySeparateSyncsSupported(value uint8) (err error) {
	return instance.SetProperty("SeparateSyncsSupported", (value))
}

// GetSeparateSyncsSupported gets the value of SeparateSyncsSupported for the instance
func (instance *WmiMonitorAnalogVideoInputParams) GetPropertySeparateSyncsSupported() (value uint8, err error) {
	retValue, err := instance.GetProperty("SeparateSyncsSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSerrationOfVsyncRequired sets the value of SerrationOfVsyncRequired for the instance
func (instance *WmiMonitorAnalogVideoInputParams) SetPropertySerrationOfVsyncRequired(value uint8) (err error) {
	return instance.SetProperty("SerrationOfVsyncRequired", (value))
}

// GetSerrationOfVsyncRequired gets the value of SerrationOfVsyncRequired for the instance
func (instance *WmiMonitorAnalogVideoInputParams) GetPropertySerrationOfVsyncRequired() (value uint8, err error) {
	retValue, err := instance.GetProperty("SerrationOfVsyncRequired")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSetupExpected sets the value of SetupExpected for the instance
func (instance *WmiMonitorAnalogVideoInputParams) SetPropertySetupExpected(value uint8) (err error) {
	return instance.SetProperty("SetupExpected", (value))
}

// GetSetupExpected gets the value of SetupExpected for the instance
func (instance *WmiMonitorAnalogVideoInputParams) GetPropertySetupExpected() (value uint8, err error) {
	retValue, err := instance.GetProperty("SetupExpected")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSignalLevelStandard sets the value of SignalLevelStandard for the instance
func (instance *WmiMonitorAnalogVideoInputParams) SetPropertySignalLevelStandard(value uint8) (err error) {
	return instance.SetProperty("SignalLevelStandard", (value))
}

// GetSignalLevelStandard gets the value of SignalLevelStandard for the instance
func (instance *WmiMonitorAnalogVideoInputParams) GetPropertySignalLevelStandard() (value uint8, err error) {
	retValue, err := instance.GetProperty("SignalLevelStandard")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSyncOnGreenVideoSupported sets the value of SyncOnGreenVideoSupported for the instance
func (instance *WmiMonitorAnalogVideoInputParams) SetPropertySyncOnGreenVideoSupported(value uint8) (err error) {
	return instance.SetProperty("SyncOnGreenVideoSupported", (value))
}

// GetSyncOnGreenVideoSupported gets the value of SyncOnGreenVideoSupported for the instance
func (instance *WmiMonitorAnalogVideoInputParams) GetPropertySyncOnGreenVideoSupported() (value uint8, err error) {
	retValue, err := instance.GetProperty("SyncOnGreenVideoSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}
