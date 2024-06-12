// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SystemConfig_V2_Platform struct
type SystemConfig_V2_Platform struct {
	*SystemConfig_V2

	//
	BiosDate string

	//
	BiosVersion string

	//
	SystemManufacturer string

	//
	SystemProductName string
}

func NewSystemConfig_V2_PlatformEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_Platform, err error) {
	tmp, err := NewSystemConfig_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_Platform{
		SystemConfig_V2: tmp,
	}
	return
}

func NewSystemConfig_V2_PlatformEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V2_Platform, err error) {
	tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_Platform{
		SystemConfig_V2: tmp,
	}
	return
}

// SetBiosDate sets the value of BiosDate for the instance
func (instance *SystemConfig_V2_Platform) SetPropertyBiosDate(value string) (err error) {
	return instance.SetProperty("BiosDate", (value))
}

// GetBiosDate gets the value of BiosDate for the instance
func (instance *SystemConfig_V2_Platform) GetPropertyBiosDate() (value string, err error) {
	retValue, err := instance.GetProperty("BiosDate")
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

// SetBiosVersion sets the value of BiosVersion for the instance
func (instance *SystemConfig_V2_Platform) SetPropertyBiosVersion(value string) (err error) {
	return instance.SetProperty("BiosVersion", (value))
}

// GetBiosVersion gets the value of BiosVersion for the instance
func (instance *SystemConfig_V2_Platform) GetPropertyBiosVersion() (value string, err error) {
	retValue, err := instance.GetProperty("BiosVersion")
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

// SetSystemManufacturer sets the value of SystemManufacturer for the instance
func (instance *SystemConfig_V2_Platform) SetPropertySystemManufacturer(value string) (err error) {
	return instance.SetProperty("SystemManufacturer", (value))
}

// GetSystemManufacturer gets the value of SystemManufacturer for the instance
func (instance *SystemConfig_V2_Platform) GetPropertySystemManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("SystemManufacturer")
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

// SetSystemProductName sets the value of SystemProductName for the instance
func (instance *SystemConfig_V2_Platform) SetPropertySystemProductName(value string) (err error) {
	return instance.SetProperty("SystemProductName", (value))
}

// GetSystemProductName gets the value of SystemProductName for the instance
func (instance *SystemConfig_V2_Platform) GetPropertySystemProductName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemProductName")
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
