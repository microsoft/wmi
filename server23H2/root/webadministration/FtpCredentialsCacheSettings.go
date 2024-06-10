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

// FtpCredentialsCacheSettings struct
type FtpCredentialsCacheSettings struct {
	*EmbeddedObject

	//
	CredentialsCache []FtpCredentialsCacheElement
}

func NewFtpCredentialsCacheSettingsEx1(instance *cim.WmiInstance) (newInstance *FtpCredentialsCacheSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpCredentialsCacheSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpCredentialsCacheSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpCredentialsCacheSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpCredentialsCacheSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetCredentialsCache sets the value of CredentialsCache for the instance
func (instance *FtpCredentialsCacheSettings) SetPropertyCredentialsCache(value []FtpCredentialsCacheElement) (err error) {
	return instance.SetProperty("CredentialsCache", (value))
}

// GetCredentialsCache gets the value of CredentialsCache for the instance
func (instance *FtpCredentialsCacheSettings) GetPropertyCredentialsCache() (value []FtpCredentialsCacheElement, err error) {
	retValue, err := instance.GetProperty("CredentialsCache")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(FtpCredentialsCacheElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " FtpCredentialsCacheElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, FtpCredentialsCacheElement(valuetmp))
	}

	return
}
