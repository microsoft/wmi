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

// IsapiCgiRestrictionSection struct
type IsapiCgiRestrictionSection struct {
	*ConfigurationSectionWithCollection

	//
	IsapiCgiRestriction []IsapiCgiRestrictionElement

	//
	NotListedCgisAllowed bool

	//
	NotListedIsapisAllowed bool
}

func NewIsapiCgiRestrictionSectionEx1(instance *cim.WmiInstance) (newInstance *IsapiCgiRestrictionSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &IsapiCgiRestrictionSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewIsapiCgiRestrictionSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *IsapiCgiRestrictionSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &IsapiCgiRestrictionSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetIsapiCgiRestriction sets the value of IsapiCgiRestriction for the instance
func (instance *IsapiCgiRestrictionSection) SetPropertyIsapiCgiRestriction(value []IsapiCgiRestrictionElement) (err error) {
	return instance.SetProperty("IsapiCgiRestriction", (value))
}

// GetIsapiCgiRestriction gets the value of IsapiCgiRestriction for the instance
func (instance *IsapiCgiRestrictionSection) GetPropertyIsapiCgiRestriction() (value []IsapiCgiRestrictionElement, err error) {
	retValue, err := instance.GetProperty("IsapiCgiRestriction")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(IsapiCgiRestrictionElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " IsapiCgiRestrictionElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, IsapiCgiRestrictionElement(valuetmp))
	}

	return
}

// SetNotListedCgisAllowed sets the value of NotListedCgisAllowed for the instance
func (instance *IsapiCgiRestrictionSection) SetPropertyNotListedCgisAllowed(value bool) (err error) {
	return instance.SetProperty("NotListedCgisAllowed", (value))
}

// GetNotListedCgisAllowed gets the value of NotListedCgisAllowed for the instance
func (instance *IsapiCgiRestrictionSection) GetPropertyNotListedCgisAllowed() (value bool, err error) {
	retValue, err := instance.GetProperty("NotListedCgisAllowed")
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

// SetNotListedIsapisAllowed sets the value of NotListedIsapisAllowed for the instance
func (instance *IsapiCgiRestrictionSection) SetPropertyNotListedIsapisAllowed(value bool) (err error) {
	return instance.SetProperty("NotListedIsapisAllowed", (value))
}

// GetNotListedIsapisAllowed gets the value of NotListedIsapisAllowed for the instance
func (instance *IsapiCgiRestrictionSection) GetPropertyNotListedIsapisAllowed() (value bool, err error) {
	retValue, err := instance.GetProperty("NotListedIsapisAllowed")
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
