// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WmiMonitorSupportedDisplayFeatures struct
type WmiMonitorSupportedDisplayFeatures struct {
	*cim.WmiInstance

	//
	ActiveOffSupported bool

	//
	DisplayType uint8

	//
	GTFSupported bool

	//
	HasPreferredTimingMode bool

	//
	sRGBSupported bool

	//
	StandbySupported bool

	//
	SuspendSupported bool
}

func NewWmiMonitorSupportedDisplayFeaturesEx1(instance *cim.WmiInstance) (newInstance *WmiMonitorSupportedDisplayFeatures, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &WmiMonitorSupportedDisplayFeatures{
		WmiInstance: tmp,
	}
	return
}

func NewWmiMonitorSupportedDisplayFeaturesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WmiMonitorSupportedDisplayFeatures, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WmiMonitorSupportedDisplayFeatures{
		WmiInstance: tmp,
	}
	return
}

// SetActiveOffSupported sets the value of ActiveOffSupported for the instance
func (instance *WmiMonitorSupportedDisplayFeatures) SetPropertyActiveOffSupported(value bool) (err error) {
	return instance.SetProperty("ActiveOffSupported", (value))
}

// GetActiveOffSupported gets the value of ActiveOffSupported for the instance
func (instance *WmiMonitorSupportedDisplayFeatures) GetPropertyActiveOffSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("ActiveOffSupported")
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

// SetDisplayType sets the value of DisplayType for the instance
func (instance *WmiMonitorSupportedDisplayFeatures) SetPropertyDisplayType(value uint8) (err error) {
	return instance.SetProperty("DisplayType", (value))
}

// GetDisplayType gets the value of DisplayType for the instance
func (instance *WmiMonitorSupportedDisplayFeatures) GetPropertyDisplayType() (value uint8, err error) {
	retValue, err := instance.GetProperty("DisplayType")
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

// SetGTFSupported sets the value of GTFSupported for the instance
func (instance *WmiMonitorSupportedDisplayFeatures) SetPropertyGTFSupported(value bool) (err error) {
	return instance.SetProperty("GTFSupported", (value))
}

// GetGTFSupported gets the value of GTFSupported for the instance
func (instance *WmiMonitorSupportedDisplayFeatures) GetPropertyGTFSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("GTFSupported")
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

// SetHasPreferredTimingMode sets the value of HasPreferredTimingMode for the instance
func (instance *WmiMonitorSupportedDisplayFeatures) SetPropertyHasPreferredTimingMode(value bool) (err error) {
	return instance.SetProperty("HasPreferredTimingMode", (value))
}

// GetHasPreferredTimingMode gets the value of HasPreferredTimingMode for the instance
func (instance *WmiMonitorSupportedDisplayFeatures) GetPropertyHasPreferredTimingMode() (value bool, err error) {
	retValue, err := instance.GetProperty("HasPreferredTimingMode")
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

// SetsRGBSupported sets the value of sRGBSupported for the instance
func (instance *WmiMonitorSupportedDisplayFeatures) SetPropertysRGBSupported(value bool) (err error) {
	return instance.SetProperty("sRGBSupported", (value))
}

// GetsRGBSupported gets the value of sRGBSupported for the instance
func (instance *WmiMonitorSupportedDisplayFeatures) GetPropertysRGBSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("sRGBSupported")
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

// SetStandbySupported sets the value of StandbySupported for the instance
func (instance *WmiMonitorSupportedDisplayFeatures) SetPropertyStandbySupported(value bool) (err error) {
	return instance.SetProperty("StandbySupported", (value))
}

// GetStandbySupported gets the value of StandbySupported for the instance
func (instance *WmiMonitorSupportedDisplayFeatures) GetPropertyStandbySupported() (value bool, err error) {
	retValue, err := instance.GetProperty("StandbySupported")
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

// SetSuspendSupported sets the value of SuspendSupported for the instance
func (instance *WmiMonitorSupportedDisplayFeatures) SetPropertySuspendSupported(value bool) (err error) {
	return instance.SetProperty("SuspendSupported", (value))
}

// GetSuspendSupported gets the value of SuspendSupported for the instance
func (instance *WmiMonitorSupportedDisplayFeatures) GetPropertySuspendSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("SuspendSupported")
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
