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

// FtpProviderSettings struct
type FtpProviderSettings struct {
	*EmbeddedObject

	//
	Providers []FtpProviderElement
}

func NewFtpProviderSettingsEx1(instance *cim.WmiInstance) (newInstance *FtpProviderSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpProviderSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpProviderSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpProviderSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpProviderSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetProviders sets the value of Providers for the instance
func (instance *FtpProviderSettings) SetPropertyProviders(value []FtpProviderElement) (err error) {
	return instance.SetProperty("Providers", (value))
}

// GetProviders gets the value of Providers for the instance
func (instance *FtpProviderSettings) GetPropertyProviders() (value []FtpProviderElement, err error) {
	retValue, err := instance.GetProperty("Providers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(FtpProviderElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " FtpProviderElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, FtpProviderElement(valuetmp))
	}

	return
}
