// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.ServerManager
//
// ////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ServerServiceDetail struct
type MSFT_ServerServiceDetail struct {
	*cim.WmiInstance

	//
	DependentServices []string

	//
	Description string

	//
	DisplayName string

	//
	ExitCode uint64

	//
	IsDelayedAutoStart bool

	//
	IsTriggered bool

	//
	Name string

	//
	StartupType uint32

	//
	Status uint32

	//
	SupportedControlCodes uint32
}

func NewMSFT_ServerServiceDetailEx1(instance *cim.WmiInstance) (newInstance *MSFT_ServerServiceDetail, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerServiceDetail{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ServerServiceDetailEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ServerServiceDetail, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerServiceDetail{
		WmiInstance: tmp,
	}
	return
}

// SetDependentServices sets the value of DependentServices for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertyDependentServices(value []string) (err error) {
	return instance.SetProperty("DependentServices", (value))
}

// GetDependentServices gets the value of DependentServices for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyDependentServices() (value []string, err error) {
	retValue, err := instance.GetProperty("DependentServices")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
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

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", (value))
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
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

// SetExitCode sets the value of ExitCode for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertyExitCode(value uint64) (err error) {
	return instance.SetProperty("ExitCode", (value))
}

// GetExitCode gets the value of ExitCode for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyExitCode() (value uint64, err error) {
	retValue, err := instance.GetProperty("ExitCode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetIsDelayedAutoStart sets the value of IsDelayedAutoStart for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertyIsDelayedAutoStart(value bool) (err error) {
	return instance.SetProperty("IsDelayedAutoStart", (value))
}

// GetIsDelayedAutoStart gets the value of IsDelayedAutoStart for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyIsDelayedAutoStart() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDelayedAutoStart")
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

// SetIsTriggered sets the value of IsTriggered for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertyIsTriggered(value bool) (err error) {
	return instance.SetProperty("IsTriggered", (value))
}

// GetIsTriggered gets the value of IsTriggered for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyIsTriggered() (value bool, err error) {
	retValue, err := instance.GetProperty("IsTriggered")
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
func (instance *MSFT_ServerServiceDetail) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyName() (value string, err error) {
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

// SetStartupType sets the value of StartupType for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertyStartupType(value uint32) (err error) {
	return instance.SetProperty("StartupType", (value))
}

// GetStartupType gets the value of StartupType for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyStartupType() (value uint32, err error) {
	retValue, err := instance.GetProperty("StartupType")
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

// SetStatus sets the value of Status for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
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

// SetSupportedControlCodes sets the value of SupportedControlCodes for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertySupportedControlCodes(value uint32) (err error) {
	return instance.SetProperty("SupportedControlCodes", (value))
}

// GetSupportedControlCodes gets the value of SupportedControlCodes for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertySupportedControlCodes() (value uint32, err error) {
	retValue, err := instance.GetProperty("SupportedControlCodes")
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
