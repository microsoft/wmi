// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SslStreamSecurityBindingElement struct
type SslStreamSecurityBindingElement struct {
	*BindingElement

	// Specifies if a client certificate is required for this binding.
	RequireClientCertificate bool
}

func NewSslStreamSecurityBindingElementEx1(instance *cim.WmiInstance) (newInstance *SslStreamSecurityBindingElement, err error) {
	tmp, err := NewBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SslStreamSecurityBindingElement{
		BindingElement: tmp,
	}
	return
}

func NewSslStreamSecurityBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SslStreamSecurityBindingElement, err error) {
	tmp, err := NewBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SslStreamSecurityBindingElement{
		BindingElement: tmp,
	}
	return
}

// SetRequireClientCertificate sets the value of RequireClientCertificate for the instance
func (instance *SslStreamSecurityBindingElement) SetPropertyRequireClientCertificate(value bool) (err error) {
	return instance.SetProperty("RequireClientCertificate", (value))
}

// GetRequireClientCertificate gets the value of RequireClientCertificate for the instance
func (instance *SslStreamSecurityBindingElement) GetPropertyRequireClientCertificate() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireClientCertificate")
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
