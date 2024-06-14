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

// FtpCustomAuthorization struct
type FtpCustomAuthorization struct {
	*EmbeddedObject

	//
	provider FtpProviderElement
}

func NewFtpCustomAuthorizationEx1(instance *cim.WmiInstance) (newInstance *FtpCustomAuthorization, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpCustomAuthorization{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpCustomAuthorizationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpCustomAuthorization, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpCustomAuthorization{
		EmbeddedObject: tmp,
	}
	return
}

// Setprovider sets the value of provider for the instance
func (instance *FtpCustomAuthorization) SetPropertyprovider(value FtpProviderElement) (err error) {
	return instance.SetProperty("provider", (value))
}

// Getprovider gets the value of provider for the instance
func (instance *FtpCustomAuthorization) GetPropertyprovider() (value FtpProviderElement, err error) {
	retValue, err := instance.GetProperty("provider")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpProviderElement)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpProviderElement is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpProviderElement(valuetmp)

	return
}
