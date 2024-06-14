// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WebControlsSection struct
type WebControlsSection struct {
	*ConfigurationSection

	//
	ClientScriptsLocation string
}

func NewWebControlsSectionEx1(instance *cim.WmiInstance) (newInstance *WebControlsSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WebControlsSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewWebControlsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WebControlsSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WebControlsSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetClientScriptsLocation sets the value of ClientScriptsLocation for the instance
func (instance *WebControlsSection) SetPropertyClientScriptsLocation(value string) (err error) {
	return instance.SetProperty("ClientScriptsLocation", (value))
}

// GetClientScriptsLocation gets the value of ClientScriptsLocation for the instance
func (instance *WebControlsSection) GetPropertyClientScriptsLocation() (value string, err error) {
	retValue, err := instance.GetProperty("ClientScriptsLocation")
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
