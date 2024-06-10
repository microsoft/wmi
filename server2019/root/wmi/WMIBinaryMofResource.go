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
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WMIBinaryMofResource struct
type WMIBinaryMofResource struct {
	*cim.WmiInstance

	//
	HighDateTime uint32

	//
	LowDateTime uint32

	//
	MofProcessed bool

	//
	Name string
}

func NewWMIBinaryMofResourceEx1(instance *cim.WmiInstance) (newInstance *WMIBinaryMofResource, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &WMIBinaryMofResource{
		WmiInstance: tmp,
	}
	return
}

func NewWMIBinaryMofResourceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WMIBinaryMofResource, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WMIBinaryMofResource{
		WmiInstance: tmp,
	}
	return
}

// SetHighDateTime sets the value of HighDateTime for the instance
func (instance *WMIBinaryMofResource) SetPropertyHighDateTime(value uint32) (err error) {
	return instance.SetProperty("HighDateTime", (value))
}

// GetHighDateTime gets the value of HighDateTime for the instance
func (instance *WMIBinaryMofResource) GetPropertyHighDateTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("HighDateTime")
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

// SetLowDateTime sets the value of LowDateTime for the instance
func (instance *WMIBinaryMofResource) SetPropertyLowDateTime(value uint32) (err error) {
	return instance.SetProperty("LowDateTime", (value))
}

// GetLowDateTime gets the value of LowDateTime for the instance
func (instance *WMIBinaryMofResource) GetPropertyLowDateTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("LowDateTime")
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

// SetMofProcessed sets the value of MofProcessed for the instance
func (instance *WMIBinaryMofResource) SetPropertyMofProcessed(value bool) (err error) {
	return instance.SetProperty("MofProcessed", (value))
}

// GetMofProcessed gets the value of MofProcessed for the instance
func (instance *WMIBinaryMofResource) GetPropertyMofProcessed() (value bool, err error) {
	retValue, err := instance.GetProperty("MofProcessed")
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

// SetName sets the value of Name for the instance
func (instance *WMIBinaryMofResource) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *WMIBinaryMofResource) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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
