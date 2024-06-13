// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Binding struct
type Binding struct {
	*cim.WmiInstance

	// The collection of binding elements implemented by the binding.
	BindingElements []BindingElement

	// The interval of time provided for a close operation to complete.
	CloseTimeout string

	// The name of the binding.
	Name string

	// The XML namespace of the binding.
	Namespace string

	// The interval of time provided for an open operation to complete.
	OpenTimeout string

	// The interval of time provided for a receive operation to complete.
	ReceiveTimeout string

	// The URI transport scheme that is used by the channel and listener factories that are built by the binding.
	Scheme string

	// The interval of time provided for a send operation to complete.
	SendTimeout string
}

func NewBindingEx1(instance *cim.WmiInstance) (newInstance *Binding, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Binding{
		WmiInstance: tmp,
	}
	return
}

func NewBindingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Binding, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Binding{
		WmiInstance: tmp,
	}
	return
}

// SetBindingElements sets the value of BindingElements for the instance
func (instance *Binding) SetPropertyBindingElements(value []BindingElement) (err error) {
	return instance.SetProperty("BindingElements", (value))
}

// GetBindingElements gets the value of BindingElements for the instance
func (instance *Binding) GetPropertyBindingElements() (value []BindingElement, err error) {
	retValue, err := instance.GetProperty("BindingElements")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(BindingElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " BindingElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, BindingElement(valuetmp))
	}

	return
}

// SetCloseTimeout sets the value of CloseTimeout for the instance
func (instance *Binding) SetPropertyCloseTimeout(value string) (err error) {
	return instance.SetProperty("CloseTimeout", (value))
}

// GetCloseTimeout gets the value of CloseTimeout for the instance
func (instance *Binding) GetPropertyCloseTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("CloseTimeout")
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

// SetName sets the value of Name for the instance
func (instance *Binding) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *Binding) GetPropertyName() (value string, err error) {
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

// SetNamespace sets the value of Namespace for the instance
func (instance *Binding) SetPropertyNamespace(value string) (err error) {
	return instance.SetProperty("Namespace", (value))
}

// GetNamespace gets the value of Namespace for the instance
func (instance *Binding) GetPropertyNamespace() (value string, err error) {
	retValue, err := instance.GetProperty("Namespace")
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

// SetOpenTimeout sets the value of OpenTimeout for the instance
func (instance *Binding) SetPropertyOpenTimeout(value string) (err error) {
	return instance.SetProperty("OpenTimeout", (value))
}

// GetOpenTimeout gets the value of OpenTimeout for the instance
func (instance *Binding) GetPropertyOpenTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("OpenTimeout")
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

// SetReceiveTimeout sets the value of ReceiveTimeout for the instance
func (instance *Binding) SetPropertyReceiveTimeout(value string) (err error) {
	return instance.SetProperty("ReceiveTimeout", (value))
}

// GetReceiveTimeout gets the value of ReceiveTimeout for the instance
func (instance *Binding) GetPropertyReceiveTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("ReceiveTimeout")
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

// SetScheme sets the value of Scheme for the instance
func (instance *Binding) SetPropertyScheme(value string) (err error) {
	return instance.SetProperty("Scheme", (value))
}

// GetScheme gets the value of Scheme for the instance
func (instance *Binding) GetPropertyScheme() (value string, err error) {
	retValue, err := instance.GetProperty("Scheme")
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

// SetSendTimeout sets the value of SendTimeout for the instance
func (instance *Binding) SetPropertySendTimeout(value string) (err error) {
	return instance.SetProperty("SendTimeout", (value))
}

// GetSendTimeout gets the value of SendTimeout for the instance
func (instance *Binding) GetPropertySendTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("SendTimeout")
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
