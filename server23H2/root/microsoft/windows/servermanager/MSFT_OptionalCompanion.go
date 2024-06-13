// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_OptionalCompanion struct
type MSFT_OptionalCompanion struct {
	*cim.WmiInstance

	//
	CompanionComponentName string

	//
	CompanionType uint8

	//
	PrerequisiteEnabled bool
}

func NewMSFT_OptionalCompanionEx1(instance *cim.WmiInstance) (newInstance *MSFT_OptionalCompanion, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_OptionalCompanion{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_OptionalCompanionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_OptionalCompanion, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_OptionalCompanion{
		WmiInstance: tmp,
	}
	return
}

// SetCompanionComponentName sets the value of CompanionComponentName for the instance
func (instance *MSFT_OptionalCompanion) SetPropertyCompanionComponentName(value string) (err error) {
	return instance.SetProperty("CompanionComponentName", (value))
}

// GetCompanionComponentName gets the value of CompanionComponentName for the instance
func (instance *MSFT_OptionalCompanion) GetPropertyCompanionComponentName() (value string, err error) {
	retValue, err := instance.GetProperty("CompanionComponentName")
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

// SetCompanionType sets the value of CompanionType for the instance
func (instance *MSFT_OptionalCompanion) SetPropertyCompanionType(value uint8) (err error) {
	return instance.SetProperty("CompanionType", (value))
}

// GetCompanionType gets the value of CompanionType for the instance
func (instance *MSFT_OptionalCompanion) GetPropertyCompanionType() (value uint8, err error) {
	retValue, err := instance.GetProperty("CompanionType")
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

// SetPrerequisiteEnabled sets the value of PrerequisiteEnabled for the instance
func (instance *MSFT_OptionalCompanion) SetPropertyPrerequisiteEnabled(value bool) (err error) {
	return instance.SetProperty("PrerequisiteEnabled", (value))
}

// GetPrerequisiteEnabled gets the value of PrerequisiteEnabled for the instance
func (instance *MSFT_OptionalCompanion) GetPropertyPrerequisiteEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("PrerequisiteEnabled")
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
