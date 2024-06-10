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

// FtpCustomAuthentication struct
type FtpCustomAuthentication struct {
	*EmbeddedObject

	//
	Providers FtpProviderSettings
}

func NewFtpCustomAuthenticationEx1(instance *cim.WmiInstance) (newInstance *FtpCustomAuthentication, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpCustomAuthentication{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpCustomAuthenticationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpCustomAuthentication, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpCustomAuthentication{
		EmbeddedObject: tmp,
	}
	return
}

// SetProviders sets the value of Providers for the instance
func (instance *FtpCustomAuthentication) SetPropertyProviders(value FtpProviderSettings) (err error) {
	return instance.SetProperty("Providers", (value))
}

// GetProviders gets the value of Providers for the instance
func (instance *FtpCustomAuthentication) GetPropertyProviders() (value FtpProviderSettings, err error) {
	retValue, err := instance.GetProperty("Providers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpProviderSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpProviderSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpProviderSettings(valuetmp)

	return
}
