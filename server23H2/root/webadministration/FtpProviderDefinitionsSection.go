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

// FtpProviderDefinitionsSection struct
type FtpProviderDefinitionsSection struct {
	*ConfigurationSectionWithCollection

	//
	Activation FtpProviderDefinitionElementActivation

	//
	ProviderDefinitions []FtpProviderDefinitionElement
}

func NewFtpProviderDefinitionsSectionEx1(instance *cim.WmiInstance) (newInstance *FtpProviderDefinitionsSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpProviderDefinitionsSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewFtpProviderDefinitionsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpProviderDefinitionsSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpProviderDefinitionsSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetActivation sets the value of Activation for the instance
func (instance *FtpProviderDefinitionsSection) SetPropertyActivation(value FtpProviderDefinitionElementActivation) (err error) {
	return instance.SetProperty("Activation", (value))
}

// GetActivation gets the value of Activation for the instance
func (instance *FtpProviderDefinitionsSection) GetPropertyActivation() (value FtpProviderDefinitionElementActivation, err error) {
	retValue, err := instance.GetProperty("Activation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpProviderDefinitionElementActivation)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpProviderDefinitionElementActivation is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpProviderDefinitionElementActivation(valuetmp)

	return
}

// SetProviderDefinitions sets the value of ProviderDefinitions for the instance
func (instance *FtpProviderDefinitionsSection) SetPropertyProviderDefinitions(value []FtpProviderDefinitionElement) (err error) {
	return instance.SetProperty("ProviderDefinitions", (value))
}

// GetProviderDefinitions gets the value of ProviderDefinitions for the instance
func (instance *FtpProviderDefinitionsSection) GetPropertyProviderDefinitions() (value []FtpProviderDefinitionElement, err error) {
	retValue, err := instance.GetProperty("ProviderDefinitions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(FtpProviderDefinitionElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " FtpProviderDefinitionElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, FtpProviderDefinitionElement(valuetmp))
	}

	return
}
