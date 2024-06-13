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

// WmiMonitorID struct
type WmiMonitorID struct {
	*MSMonitorClass

	//
	Active bool

	//
	InstanceName string

	//
	ManufacturerName []uint16

	//
	ProductCodeID []uint16

	//
	SerialNumberID []uint16

	//
	UserFriendlyName []uint16

	//
	UserFriendlyNameLength uint16

	//
	WeekOfManufacture uint8

	//
	YearOfManufacture uint16
}

func NewWmiMonitorIDEx1(instance *cim.WmiInstance) (newInstance *WmiMonitorID, err error) {
	tmp, err := NewMSMonitorClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WmiMonitorID{
		MSMonitorClass: tmp,
	}
	return
}

func NewWmiMonitorIDEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WmiMonitorID, err error) {
	tmp, err := NewMSMonitorClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WmiMonitorID{
		MSMonitorClass: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *WmiMonitorID) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *WmiMonitorID) GetPropertyActive() (value bool, err error) {
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
func (instance *WmiMonitorID) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *WmiMonitorID) GetPropertyInstanceName() (value string, err error) {
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

// SetManufacturerName sets the value of ManufacturerName for the instance
func (instance *WmiMonitorID) SetPropertyManufacturerName(value []uint16) (err error) {
	return instance.SetProperty("ManufacturerName", (value))
}

// GetManufacturerName gets the value of ManufacturerName for the instance
func (instance *WmiMonitorID) GetPropertyManufacturerName() (value []uint16, err error) {
	retValue, err := instance.GetProperty("ManufacturerName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetProductCodeID sets the value of ProductCodeID for the instance
func (instance *WmiMonitorID) SetPropertyProductCodeID(value []uint16) (err error) {
	return instance.SetProperty("ProductCodeID", (value))
}

// GetProductCodeID gets the value of ProductCodeID for the instance
func (instance *WmiMonitorID) GetPropertyProductCodeID() (value []uint16, err error) {
	retValue, err := instance.GetProperty("ProductCodeID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetSerialNumberID sets the value of SerialNumberID for the instance
func (instance *WmiMonitorID) SetPropertySerialNumberID(value []uint16) (err error) {
	return instance.SetProperty("SerialNumberID", (value))
}

// GetSerialNumberID gets the value of SerialNumberID for the instance
func (instance *WmiMonitorID) GetPropertySerialNumberID() (value []uint16, err error) {
	retValue, err := instance.GetProperty("SerialNumberID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetUserFriendlyName sets the value of UserFriendlyName for the instance
func (instance *WmiMonitorID) SetPropertyUserFriendlyName(value []uint16) (err error) {
	return instance.SetProperty("UserFriendlyName", (value))
}

// GetUserFriendlyName gets the value of UserFriendlyName for the instance
func (instance *WmiMonitorID) GetPropertyUserFriendlyName() (value []uint16, err error) {
	retValue, err := instance.GetProperty("UserFriendlyName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetUserFriendlyNameLength sets the value of UserFriendlyNameLength for the instance
func (instance *WmiMonitorID) SetPropertyUserFriendlyNameLength(value uint16) (err error) {
	return instance.SetProperty("UserFriendlyNameLength", (value))
}

// GetUserFriendlyNameLength gets the value of UserFriendlyNameLength for the instance
func (instance *WmiMonitorID) GetPropertyUserFriendlyNameLength() (value uint16, err error) {
	retValue, err := instance.GetProperty("UserFriendlyNameLength")
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

// SetWeekOfManufacture sets the value of WeekOfManufacture for the instance
func (instance *WmiMonitorID) SetPropertyWeekOfManufacture(value uint8) (err error) {
	return instance.SetProperty("WeekOfManufacture", (value))
}

// GetWeekOfManufacture gets the value of WeekOfManufacture for the instance
func (instance *WmiMonitorID) GetPropertyWeekOfManufacture() (value uint8, err error) {
	retValue, err := instance.GetProperty("WeekOfManufacture")
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

// SetYearOfManufacture sets the value of YearOfManufacture for the instance
func (instance *WmiMonitorID) SetPropertyYearOfManufacture(value uint16) (err error) {
	return instance.SetProperty("YearOfManufacture", (value))
}

// GetYearOfManufacture gets the value of YearOfManufacture for the instance
func (instance *WmiMonitorID) GetPropertyYearOfManufacture() (value uint16, err error) {
	retValue, err := instance.GetProperty("YearOfManufacture")
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
