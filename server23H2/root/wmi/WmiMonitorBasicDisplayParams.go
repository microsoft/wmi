// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WmiMonitorBasicDisplayParams struct
type WmiMonitorBasicDisplayParams struct {
	*MSMonitorClass

	//
	Active bool

	//
	DisplayTransferCharacteristic uint8

	//
	InstanceName string

	//
	MaxHorizontalImageSize uint8

	//
	MaxVerticalImageSize uint8

	//
	SupportedDisplayFeatures WmiMonitorSupportedDisplayFeatures

	//
	VideoInputType uint8
}

func NewWmiMonitorBasicDisplayParamsEx1(instance *cim.WmiInstance) (newInstance *WmiMonitorBasicDisplayParams, err error) {
	tmp, err := NewMSMonitorClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WmiMonitorBasicDisplayParams{
		MSMonitorClass: tmp,
	}
	return
}

func NewWmiMonitorBasicDisplayParamsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WmiMonitorBasicDisplayParams, err error) {
	tmp, err := NewMSMonitorClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WmiMonitorBasicDisplayParams{
		MSMonitorClass: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *WmiMonitorBasicDisplayParams) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *WmiMonitorBasicDisplayParams) GetPropertyActive() (value bool, err error) {
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

// SetDisplayTransferCharacteristic sets the value of DisplayTransferCharacteristic for the instance
func (instance *WmiMonitorBasicDisplayParams) SetPropertyDisplayTransferCharacteristic(value uint8) (err error) {
	return instance.SetProperty("DisplayTransferCharacteristic", (value))
}

// GetDisplayTransferCharacteristic gets the value of DisplayTransferCharacteristic for the instance
func (instance *WmiMonitorBasicDisplayParams) GetPropertyDisplayTransferCharacteristic() (value uint8, err error) {
	retValue, err := instance.GetProperty("DisplayTransferCharacteristic")
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
func (instance *WmiMonitorBasicDisplayParams) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *WmiMonitorBasicDisplayParams) GetPropertyInstanceName() (value string, err error) {
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

// SetMaxHorizontalImageSize sets the value of MaxHorizontalImageSize for the instance
func (instance *WmiMonitorBasicDisplayParams) SetPropertyMaxHorizontalImageSize(value uint8) (err error) {
	return instance.SetProperty("MaxHorizontalImageSize", (value))
}

// GetMaxHorizontalImageSize gets the value of MaxHorizontalImageSize for the instance
func (instance *WmiMonitorBasicDisplayParams) GetPropertyMaxHorizontalImageSize() (value uint8, err error) {
	retValue, err := instance.GetProperty("MaxHorizontalImageSize")
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

// SetMaxVerticalImageSize sets the value of MaxVerticalImageSize for the instance
func (instance *WmiMonitorBasicDisplayParams) SetPropertyMaxVerticalImageSize(value uint8) (err error) {
	return instance.SetProperty("MaxVerticalImageSize", (value))
}

// GetMaxVerticalImageSize gets the value of MaxVerticalImageSize for the instance
func (instance *WmiMonitorBasicDisplayParams) GetPropertyMaxVerticalImageSize() (value uint8, err error) {
	retValue, err := instance.GetProperty("MaxVerticalImageSize")
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

// SetSupportedDisplayFeatures sets the value of SupportedDisplayFeatures for the instance
func (instance *WmiMonitorBasicDisplayParams) SetPropertySupportedDisplayFeatures(value WmiMonitorSupportedDisplayFeatures) (err error) {
	return instance.SetProperty("SupportedDisplayFeatures", (value))
}

// GetSupportedDisplayFeatures gets the value of SupportedDisplayFeatures for the instance
func (instance *WmiMonitorBasicDisplayParams) GetPropertySupportedDisplayFeatures() (value WmiMonitorSupportedDisplayFeatures, err error) {
	retValue, err := instance.GetProperty("SupportedDisplayFeatures")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(WmiMonitorSupportedDisplayFeatures)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " WmiMonitorSupportedDisplayFeatures is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = WmiMonitorSupportedDisplayFeatures(valuetmp)

	return
}

// SetVideoInputType sets the value of VideoInputType for the instance
func (instance *WmiMonitorBasicDisplayParams) SetPropertyVideoInputType(value uint8) (err error) {
	return instance.SetProperty("VideoInputType", (value))
}

// GetVideoInputType gets the value of VideoInputType for the instance
func (instance *WmiMonitorBasicDisplayParams) GetPropertyVideoInputType() (value uint8, err error) {
	retValue, err := instance.GetProperty("VideoInputType")
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
