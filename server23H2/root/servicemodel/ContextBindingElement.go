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

// ContextBindingElement struct
type ContextBindingElement struct {
	*BindingElement

	// The ContextExchangeMechanism for ContextBindingElement
	ContextExchangeMechanism string

	// Whether automatic context management is enabled for ContextBindingElement
	ContextManagementEnabled bool

	// The ProtectionLevel forContextBindingElement
	ProtectionLevel string
}

func NewContextBindingElementEx1(instance *cim.WmiInstance) (newInstance *ContextBindingElement, err error) {
	tmp, err := NewBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ContextBindingElement{
		BindingElement: tmp,
	}
	return
}

func NewContextBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ContextBindingElement, err error) {
	tmp, err := NewBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ContextBindingElement{
		BindingElement: tmp,
	}
	return
}

// SetContextExchangeMechanism sets the value of ContextExchangeMechanism for the instance
func (instance *ContextBindingElement) SetPropertyContextExchangeMechanism(value string) (err error) {
	return instance.SetProperty("ContextExchangeMechanism", (value))
}

// GetContextExchangeMechanism gets the value of ContextExchangeMechanism for the instance
func (instance *ContextBindingElement) GetPropertyContextExchangeMechanism() (value string, err error) {
	retValue, err := instance.GetProperty("ContextExchangeMechanism")
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

// SetContextManagementEnabled sets the value of ContextManagementEnabled for the instance
func (instance *ContextBindingElement) SetPropertyContextManagementEnabled(value bool) (err error) {
	return instance.SetProperty("ContextManagementEnabled", (value))
}

// GetContextManagementEnabled gets the value of ContextManagementEnabled for the instance
func (instance *ContextBindingElement) GetPropertyContextManagementEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("ContextManagementEnabled")
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

// SetProtectionLevel sets the value of ProtectionLevel for the instance
func (instance *ContextBindingElement) SetPropertyProtectionLevel(value string) (err error) {
	return instance.SetProperty("ProtectionLevel", (value))
}

// GetProtectionLevel gets the value of ProtectionLevel for the instance
func (instance *ContextBindingElement) GetPropertyProtectionLevel() (value string, err error) {
	retValue, err := instance.GetProperty("ProtectionLevel")
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
