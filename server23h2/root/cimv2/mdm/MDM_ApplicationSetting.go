// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_ApplicationSetting struct
type MDM_ApplicationSetting struct {
	*cim.WmiInstance

	//
	PackageFamilyName string

	//
	SettingName string

	//
	SettingType uint32

	//
	SettingValue string
}

func NewMDM_ApplicationSettingEx1(instance *cim.WmiInstance) (newInstance *MDM_ApplicationSetting, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_ApplicationSetting{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_ApplicationSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_ApplicationSetting, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_ApplicationSetting{
		WmiInstance: tmp,
	}
	return
}

// SetPackageFamilyName sets the value of PackageFamilyName for the instance
func (instance *MDM_ApplicationSetting) SetPropertyPackageFamilyName(value string) (err error) {
	return instance.SetProperty("PackageFamilyName", (value))
}

// GetPackageFamilyName gets the value of PackageFamilyName for the instance
func (instance *MDM_ApplicationSetting) GetPropertyPackageFamilyName() (value string, err error) {
	retValue, err := instance.GetProperty("PackageFamilyName")
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

// SetSettingName sets the value of SettingName for the instance
func (instance *MDM_ApplicationSetting) SetPropertySettingName(value string) (err error) {
	return instance.SetProperty("SettingName", (value))
}

// GetSettingName gets the value of SettingName for the instance
func (instance *MDM_ApplicationSetting) GetPropertySettingName() (value string, err error) {
	retValue, err := instance.GetProperty("SettingName")
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

// SetSettingType sets the value of SettingType for the instance
func (instance *MDM_ApplicationSetting) SetPropertySettingType(value uint32) (err error) {
	return instance.SetProperty("SettingType", (value))
}

// GetSettingType gets the value of SettingType for the instance
func (instance *MDM_ApplicationSetting) GetPropertySettingType() (value uint32, err error) {
	retValue, err := instance.GetProperty("SettingType")
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

// SetSettingValue sets the value of SettingValue for the instance
func (instance *MDM_ApplicationSetting) SetPropertySettingValue(value string) (err error) {
	return instance.SetProperty("SettingValue", (value))
}

// GetSettingValue gets the value of SettingValue for the instance
func (instance *MDM_ApplicationSetting) GetPropertySettingValue() (value string, err error) {
	retValue, err := instance.GetProperty("SettingValue")
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
