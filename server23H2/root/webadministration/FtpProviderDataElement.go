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

// FtpProviderDataElement struct
type FtpProviderDataElement struct {
	*CollectionElement

	//
	EncryptedValue string

	//
	key string

	//
	Value string
}

func NewFtpProviderDataElementEx1(instance *cim.WmiInstance) (newInstance *FtpProviderDataElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpProviderDataElement{
		CollectionElement: tmp,
	}
	return
}

func NewFtpProviderDataElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpProviderDataElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpProviderDataElement{
		CollectionElement: tmp,
	}
	return
}

// SetEncryptedValue sets the value of EncryptedValue for the instance
func (instance *FtpProviderDataElement) SetPropertyEncryptedValue(value string) (err error) {
	return instance.SetProperty("EncryptedValue", (value))
}

// GetEncryptedValue gets the value of EncryptedValue for the instance
func (instance *FtpProviderDataElement) GetPropertyEncryptedValue() (value string, err error) {
	retValue, err := instance.GetProperty("EncryptedValue")
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

// Setkey sets the value of key for the instance
func (instance *FtpProviderDataElement) SetPropertykey(value string) (err error) {
	return instance.SetProperty("key", (value))
}

// Getkey gets the value of key for the instance
func (instance *FtpProviderDataElement) GetPropertykey() (value string, err error) {
	retValue, err := instance.GetProperty("key")
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

// SetValue sets the value of Value for the instance
func (instance *FtpProviderDataElement) SetPropertyValue(value string) (err error) {
	return instance.SetProperty("Value", (value))
}

// GetValue gets the value of Value for the instance
func (instance *FtpProviderDataElement) GetPropertyValue() (value string, err error) {
	retValue, err := instance.GetProperty("Value")
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
