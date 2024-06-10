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

// SystemWindowsFormsSection struct
type SystemWindowsFormsSection struct {
	*ConfigurationSection

	//
	JitDebugging bool
}

func NewSystemWindowsFormsSectionEx1(instance *cim.WmiInstance) (newInstance *SystemWindowsFormsSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemWindowsFormsSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewSystemWindowsFormsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemWindowsFormsSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemWindowsFormsSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetJitDebugging sets the value of JitDebugging for the instance
func (instance *SystemWindowsFormsSection) SetPropertyJitDebugging(value bool) (err error) {
	return instance.SetProperty("JitDebugging", (value))
}

// GetJitDebugging gets the value of JitDebugging for the instance
func (instance *SystemWindowsFormsSection) GetPropertyJitDebugging() (value bool, err error) {
	retValue, err := instance.GetProperty("JitDebugging")
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
