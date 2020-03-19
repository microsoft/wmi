// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
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
	return instance.SetProperty("CompanionComponentName", value)
}

// GetCompanionComponentName gets the value of CompanionComponentName for the instance
func (instance *MSFT_OptionalCompanion) GetPropertyCompanionComponentName() (value string, err error) {
	retValue, err := instance.GetProperty("CompanionComponentName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompanionType sets the value of CompanionType for the instance
func (instance *MSFT_OptionalCompanion) SetPropertyCompanionType(value uint8) (err error) {
	return instance.SetProperty("CompanionType", value)
}

// GetCompanionType gets the value of CompanionType for the instance
func (instance *MSFT_OptionalCompanion) GetPropertyCompanionType() (value uint8, err error) {
	retValue, err := instance.GetProperty("CompanionType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrerequisiteEnabled sets the value of PrerequisiteEnabled for the instance
func (instance *MSFT_OptionalCompanion) SetPropertyPrerequisiteEnabled(value bool) (err error) {
	return instance.SetProperty("PrerequisiteEnabled", value)
}

// GetPrerequisiteEnabled gets the value of PrerequisiteEnabled for the instance
func (instance *MSFT_OptionalCompanion) GetPropertyPrerequisiteEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("PrerequisiteEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
