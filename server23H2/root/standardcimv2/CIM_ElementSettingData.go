// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_ElementSettingData struct
type CIM_ElementSettingData struct {
	*cim.WmiInstance

	//
	IsCurrent uint16

	//
	IsDefault uint16

	//
	IsNext uint16

	//
	ManagedElement CIM_ManagedElement

	//
	SettingData CIM_SettingData
}

func NewCIM_ElementSettingDataEx1(instance *cim.WmiInstance) (newInstance *CIM_ElementSettingData, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_ElementSettingData{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_ElementSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ElementSettingData, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ElementSettingData{
		WmiInstance: tmp,
	}
	return
}

// SetIsCurrent sets the value of IsCurrent for the instance
func (instance *CIM_ElementSettingData) SetPropertyIsCurrent(value uint16) (err error) {
	return instance.SetProperty("IsCurrent", (value))
}

// GetIsCurrent gets the value of IsCurrent for the instance
func (instance *CIM_ElementSettingData) GetPropertyIsCurrent() (value uint16, err error) {
	retValue, err := instance.GetProperty("IsCurrent")
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

// SetIsDefault sets the value of IsDefault for the instance
func (instance *CIM_ElementSettingData) SetPropertyIsDefault(value uint16) (err error) {
	return instance.SetProperty("IsDefault", (value))
}

// GetIsDefault gets the value of IsDefault for the instance
func (instance *CIM_ElementSettingData) GetPropertyIsDefault() (value uint16, err error) {
	retValue, err := instance.GetProperty("IsDefault")
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

// SetIsNext sets the value of IsNext for the instance
func (instance *CIM_ElementSettingData) SetPropertyIsNext(value uint16) (err error) {
	return instance.SetProperty("IsNext", (value))
}

// GetIsNext gets the value of IsNext for the instance
func (instance *CIM_ElementSettingData) GetPropertyIsNext() (value uint16, err error) {
	retValue, err := instance.GetProperty("IsNext")
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

// SetManagedElement sets the value of ManagedElement for the instance
func (instance *CIM_ElementSettingData) SetPropertyManagedElement(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("ManagedElement", (value))
}

// GetManagedElement gets the value of ManagedElement for the instance
func (instance *CIM_ElementSettingData) GetPropertyManagedElement() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("ManagedElement")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CIM_ManagedElement)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CIM_ManagedElement is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CIM_ManagedElement(valuetmp)

	return
}

// SetSettingData sets the value of SettingData for the instance
func (instance *CIM_ElementSettingData) SetPropertySettingData(value CIM_SettingData) (err error) {
	return instance.SetProperty("SettingData", (value))
}

// GetSettingData gets the value of SettingData for the instance
func (instance *CIM_ElementSettingData) GetPropertySettingData() (value CIM_SettingData, err error) {
	retValue, err := instance.GetProperty("SettingData")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CIM_SettingData)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CIM_SettingData is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CIM_SettingData(valuetmp)

	return
}
