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

// FtpProviderDataSettings struct
type FtpProviderDataSettings struct {
	*CollectionElement

	//
	Name string

	//
	ProviderData []FtpProviderDataElement
}

func NewFtpProviderDataSettingsEx1(instance *cim.WmiInstance) (newInstance *FtpProviderDataSettings, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpProviderDataSettings{
		CollectionElement: tmp,
	}
	return
}

func NewFtpProviderDataSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpProviderDataSettings, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpProviderDataSettings{
		CollectionElement: tmp,
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *FtpProviderDataSettings) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *FtpProviderDataSettings) GetPropertyName() (value string, err error) {
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

// SetProviderData sets the value of ProviderData for the instance
func (instance *FtpProviderDataSettings) SetPropertyProviderData(value []FtpProviderDataElement) (err error) {
	return instance.SetProperty("ProviderData", (value))
}

// GetProviderData gets the value of ProviderData for the instance
func (instance *FtpProviderDataSettings) GetPropertyProviderData() (value []FtpProviderDataElement, err error) {
	retValue, err := instance.GetProperty("ProviderData")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(FtpProviderDataElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " FtpProviderDataElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, FtpProviderDataElement(valuetmp))
	}

	return
}
