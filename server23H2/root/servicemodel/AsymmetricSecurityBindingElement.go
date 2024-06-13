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

// AsymmetricSecurityBindingElement struct
type AsymmetricSecurityBindingElement struct {
	*SecurityBindingElement

	// The order of message encryption and signing for this binding.
	MessageProtectionOrder string

	// Whether the binding requires signature confirmation.
	RequireSignatureConfirmation bool
}

func NewAsymmetricSecurityBindingElementEx1(instance *cim.WmiInstance) (newInstance *AsymmetricSecurityBindingElement, err error) {
	tmp, err := NewSecurityBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AsymmetricSecurityBindingElement{
		SecurityBindingElement: tmp,
	}
	return
}

func NewAsymmetricSecurityBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AsymmetricSecurityBindingElement, err error) {
	tmp, err := NewSecurityBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AsymmetricSecurityBindingElement{
		SecurityBindingElement: tmp,
	}
	return
}

// SetMessageProtectionOrder sets the value of MessageProtectionOrder for the instance
func (instance *AsymmetricSecurityBindingElement) SetPropertyMessageProtectionOrder(value string) (err error) {
	return instance.SetProperty("MessageProtectionOrder", (value))
}

// GetMessageProtectionOrder gets the value of MessageProtectionOrder for the instance
func (instance *AsymmetricSecurityBindingElement) GetPropertyMessageProtectionOrder() (value string, err error) {
	retValue, err := instance.GetProperty("MessageProtectionOrder")
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

// SetRequireSignatureConfirmation sets the value of RequireSignatureConfirmation for the instance
func (instance *AsymmetricSecurityBindingElement) SetPropertyRequireSignatureConfirmation(value bool) (err error) {
	return instance.SetProperty("RequireSignatureConfirmation", (value))
}

// GetRequireSignatureConfirmation gets the value of RequireSignatureConfirmation for the instance
func (instance *AsymmetricSecurityBindingElement) GetPropertyRequireSignatureConfirmation() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireSignatureConfirmation")
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
