// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_BIOSElement struct
type Msvm_BIOSElement struct {
	*CIM_BIOSElement

	//
	BaseBoardSerialNumber string

	//
	BIOSGUID string

	//
	BIOSNumLock bool

	//
	BIOSSerialNumber string

	//
	BootOrder []uint16

	//
	ChassisAssetTag string

	//
	ChassisSerialNumber string

	//
	EnableHibernation bool
}

func NewMsvm_BIOSElementEx1(instance *cim.WmiInstance) (newInstance *Msvm_BIOSElement, err error) {
	tmp, err := NewCIM_BIOSElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_BIOSElement{
		CIM_BIOSElement: tmp,
	}
	return
}

func NewMsvm_BIOSElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_BIOSElement, err error) {
	tmp, err := NewCIM_BIOSElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_BIOSElement{
		CIM_BIOSElement: tmp,
	}
	return
}

// SetBaseBoardSerialNumber sets the value of BaseBoardSerialNumber for the instance
func (instance *Msvm_BIOSElement) SetPropertyBaseBoardSerialNumber(value string) (err error) {
	return instance.SetProperty("BaseBoardSerialNumber", (value))
}

// GetBaseBoardSerialNumber gets the value of BaseBoardSerialNumber for the instance
func (instance *Msvm_BIOSElement) GetPropertyBaseBoardSerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("BaseBoardSerialNumber")
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

// SetBIOSGUID sets the value of BIOSGUID for the instance
func (instance *Msvm_BIOSElement) SetPropertyBIOSGUID(value string) (err error) {
	return instance.SetProperty("BIOSGUID", (value))
}

// GetBIOSGUID gets the value of BIOSGUID for the instance
func (instance *Msvm_BIOSElement) GetPropertyBIOSGUID() (value string, err error) {
	retValue, err := instance.GetProperty("BIOSGUID")
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

// SetBIOSNumLock sets the value of BIOSNumLock for the instance
func (instance *Msvm_BIOSElement) SetPropertyBIOSNumLock(value bool) (err error) {
	return instance.SetProperty("BIOSNumLock", (value))
}

// GetBIOSNumLock gets the value of BIOSNumLock for the instance
func (instance *Msvm_BIOSElement) GetPropertyBIOSNumLock() (value bool, err error) {
	retValue, err := instance.GetProperty("BIOSNumLock")
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

// SetBIOSSerialNumber sets the value of BIOSSerialNumber for the instance
func (instance *Msvm_BIOSElement) SetPropertyBIOSSerialNumber(value string) (err error) {
	return instance.SetProperty("BIOSSerialNumber", (value))
}

// GetBIOSSerialNumber gets the value of BIOSSerialNumber for the instance
func (instance *Msvm_BIOSElement) GetPropertyBIOSSerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("BIOSSerialNumber")
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

// SetBootOrder sets the value of BootOrder for the instance
func (instance *Msvm_BIOSElement) SetPropertyBootOrder(value []uint16) (err error) {
	return instance.SetProperty("BootOrder", (value))
}

// GetBootOrder gets the value of BootOrder for the instance
func (instance *Msvm_BIOSElement) GetPropertyBootOrder() (value []uint16, err error) {
	retValue, err := instance.GetProperty("BootOrder")
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

// SetChassisAssetTag sets the value of ChassisAssetTag for the instance
func (instance *Msvm_BIOSElement) SetPropertyChassisAssetTag(value string) (err error) {
	return instance.SetProperty("ChassisAssetTag", (value))
}

// GetChassisAssetTag gets the value of ChassisAssetTag for the instance
func (instance *Msvm_BIOSElement) GetPropertyChassisAssetTag() (value string, err error) {
	retValue, err := instance.GetProperty("ChassisAssetTag")
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

// SetChassisSerialNumber sets the value of ChassisSerialNumber for the instance
func (instance *Msvm_BIOSElement) SetPropertyChassisSerialNumber(value string) (err error) {
	return instance.SetProperty("ChassisSerialNumber", (value))
}

// GetChassisSerialNumber gets the value of ChassisSerialNumber for the instance
func (instance *Msvm_BIOSElement) GetPropertyChassisSerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("ChassisSerialNumber")
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

// SetEnableHibernation sets the value of EnableHibernation for the instance
func (instance *Msvm_BIOSElement) SetPropertyEnableHibernation(value bool) (err error) {
	return instance.SetProperty("EnableHibernation", (value))
}

// GetEnableHibernation gets the value of EnableHibernation for the instance
func (instance *Msvm_BIOSElement) GetPropertyEnableHibernation() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableHibernation")
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
func (instance *Msvm_BIOSElement) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}
