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

// WmiMonitorColorCharacteristics struct
type WmiMonitorColorCharacteristics struct {
	*MSMonitorClass

	//
	Active bool

	//
	Blue WmiMonitorColorXYZinCIE

	//
	DefaultWhite WmiMonitorColorXYZinCIE

	//
	Green WmiMonitorColorXYZinCIE

	//
	InstanceName string

	//
	Red WmiMonitorColorXYZinCIE
}

func NewWmiMonitorColorCharacteristicsEx1(instance *cim.WmiInstance) (newInstance *WmiMonitorColorCharacteristics, err error) {
	tmp, err := NewMSMonitorClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WmiMonitorColorCharacteristics{
		MSMonitorClass: tmp,
	}
	return
}

func NewWmiMonitorColorCharacteristicsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WmiMonitorColorCharacteristics, err error) {
	tmp, err := NewMSMonitorClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WmiMonitorColorCharacteristics{
		MSMonitorClass: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *WmiMonitorColorCharacteristics) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *WmiMonitorColorCharacteristics) GetPropertyActive() (value bool, err error) {
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

// SetBlue sets the value of Blue for the instance
func (instance *WmiMonitorColorCharacteristics) SetPropertyBlue(value WmiMonitorColorXYZinCIE) (err error) {
	return instance.SetProperty("Blue", (value))
}

// GetBlue gets the value of Blue for the instance
func (instance *WmiMonitorColorCharacteristics) GetPropertyBlue() (value WmiMonitorColorXYZinCIE, err error) {
	retValue, err := instance.GetProperty("Blue")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(WmiMonitorColorXYZinCIE)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " WmiMonitorColorXYZinCIE is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = WmiMonitorColorXYZinCIE(valuetmp)

	return
}

// SetDefaultWhite sets the value of DefaultWhite for the instance
func (instance *WmiMonitorColorCharacteristics) SetPropertyDefaultWhite(value WmiMonitorColorXYZinCIE) (err error) {
	return instance.SetProperty("DefaultWhite", (value))
}

// GetDefaultWhite gets the value of DefaultWhite for the instance
func (instance *WmiMonitorColorCharacteristics) GetPropertyDefaultWhite() (value WmiMonitorColorXYZinCIE, err error) {
	retValue, err := instance.GetProperty("DefaultWhite")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(WmiMonitorColorXYZinCIE)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " WmiMonitorColorXYZinCIE is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = WmiMonitorColorXYZinCIE(valuetmp)

	return
}

// SetGreen sets the value of Green for the instance
func (instance *WmiMonitorColorCharacteristics) SetPropertyGreen(value WmiMonitorColorXYZinCIE) (err error) {
	return instance.SetProperty("Green", (value))
}

// GetGreen gets the value of Green for the instance
func (instance *WmiMonitorColorCharacteristics) GetPropertyGreen() (value WmiMonitorColorXYZinCIE, err error) {
	retValue, err := instance.GetProperty("Green")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(WmiMonitorColorXYZinCIE)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " WmiMonitorColorXYZinCIE is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = WmiMonitorColorXYZinCIE(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *WmiMonitorColorCharacteristics) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *WmiMonitorColorCharacteristics) GetPropertyInstanceName() (value string, err error) {
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

// SetRed sets the value of Red for the instance
func (instance *WmiMonitorColorCharacteristics) SetPropertyRed(value WmiMonitorColorXYZinCIE) (err error) {
	return instance.SetProperty("Red", (value))
}

// GetRed gets the value of Red for the instance
func (instance *WmiMonitorColorCharacteristics) GetPropertyRed() (value WmiMonitorColorXYZinCIE, err error) {
	retValue, err := instance.GetProperty("Red")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(WmiMonitorColorXYZinCIE)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " WmiMonitorColorXYZinCIE is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = WmiMonitorColorXYZinCIE(valuetmp)

	return
}
