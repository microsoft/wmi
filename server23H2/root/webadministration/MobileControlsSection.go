// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MobileControlsSection struct
type MobileControlsSection struct {
	*ConfigurationSectionWithCollection

	//
	AllowCustomAttributes bool

	//
	CookielessDataDictionaryType string

	//
	MobileControls []DeviceElement

	//
	SessionStateHistorySize int32
}

func NewMobileControlsSectionEx1(instance *cim.WmiInstance) (newInstance *MobileControlsSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MobileControlsSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewMobileControlsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MobileControlsSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MobileControlsSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetAllowCustomAttributes sets the value of AllowCustomAttributes for the instance
func (instance *MobileControlsSection) SetPropertyAllowCustomAttributes(value bool) (err error) {
	return instance.SetProperty("AllowCustomAttributes", (value))
}

// GetAllowCustomAttributes gets the value of AllowCustomAttributes for the instance
func (instance *MobileControlsSection) GetPropertyAllowCustomAttributes() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowCustomAttributes")
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

// SetCookielessDataDictionaryType sets the value of CookielessDataDictionaryType for the instance
func (instance *MobileControlsSection) SetPropertyCookielessDataDictionaryType(value string) (err error) {
	return instance.SetProperty("CookielessDataDictionaryType", (value))
}

// GetCookielessDataDictionaryType gets the value of CookielessDataDictionaryType for the instance
func (instance *MobileControlsSection) GetPropertyCookielessDataDictionaryType() (value string, err error) {
	retValue, err := instance.GetProperty("CookielessDataDictionaryType")
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

// SetMobileControls sets the value of MobileControls for the instance
func (instance *MobileControlsSection) SetPropertyMobileControls(value []DeviceElement) (err error) {
	return instance.SetProperty("MobileControls", (value))
}

// GetMobileControls gets the value of MobileControls for the instance
func (instance *MobileControlsSection) GetPropertyMobileControls() (value []DeviceElement, err error) {
	retValue, err := instance.GetProperty("MobileControls")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(DeviceElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " DeviceElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DeviceElement(valuetmp))
	}

	return
}

// SetSessionStateHistorySize sets the value of SessionStateHistorySize for the instance
func (instance *MobileControlsSection) SetPropertySessionStateHistorySize(value int32) (err error) {
	return instance.SetProperty("SessionStateHistorySize", (value))
}

// GetSessionStateHistorySize gets the value of SessionStateHistorySize for the instance
func (instance *MobileControlsSection) GetPropertySessionStateHistorySize() (value int32, err error) {
	retValue, err := instance.GetProperty("SessionStateHistorySize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}
