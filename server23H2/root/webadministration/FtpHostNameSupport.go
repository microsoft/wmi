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

// FtpHostNameSupport struct
type FtpHostNameSupport struct {
	*EmbeddedObject

	//
	UseDomainNameAsHostName bool
}

func NewFtpHostNameSupportEx1(instance *cim.WmiInstance) (newInstance *FtpHostNameSupport, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpHostNameSupport{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpHostNameSupportEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpHostNameSupport, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpHostNameSupport{
		EmbeddedObject: tmp,
	}
	return
}

// SetUseDomainNameAsHostName sets the value of UseDomainNameAsHostName for the instance
func (instance *FtpHostNameSupport) SetPropertyUseDomainNameAsHostName(value bool) (err error) {
	return instance.SetProperty("UseDomainNameAsHostName", (value))
}

// GetUseDomainNameAsHostName gets the value of UseDomainNameAsHostName for the instance
func (instance *FtpHostNameSupport) GetPropertyUseDomainNameAsHostName() (value bool, err error) {
	retValue, err := instance.GetProperty("UseDomainNameAsHostName")
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
