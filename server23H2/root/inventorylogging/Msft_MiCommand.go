// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msft_MiCommand struct
type Msft_MiCommand struct {
	*Msft_MiStream

	//
	ClassName string

	//
	Input interface{}

	//
	MethodName string

	//
	NamespaceName string

	//
	Parameters interface{}
}

func NewMsft_MiCommandEx1(instance *cim.WmiInstance) (newInstance *Msft_MiCommand, err error) {
	tmp, err := NewMsft_MiStreamEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msft_MiCommand{
		Msft_MiStream: tmp,
	}
	return
}

func NewMsft_MiCommandEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msft_MiCommand, err error) {
	tmp, err := NewMsft_MiStreamEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msft_MiCommand{
		Msft_MiStream: tmp,
	}
	return
}

// SetClassName sets the value of ClassName for the instance
func (instance *Msft_MiCommand) SetPropertyClassName(value string) (err error) {
	return instance.SetProperty("ClassName", (value))
}

// GetClassName gets the value of ClassName for the instance
func (instance *Msft_MiCommand) GetPropertyClassName() (value string, err error) {
	retValue, err := instance.GetProperty("ClassName")
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

// SetInput sets the value of Input for the instance
func (instance *Msft_MiCommand) SetPropertyInput(value interface{}) (err error) {
	return instance.SetProperty("Input", (value))
}

// GetInput gets the value of Input for the instance
func (instance *Msft_MiCommand) GetPropertyInput() (value interface{}, err error) {
	retValue, err := instance.GetProperty("Input")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetMethodName sets the value of MethodName for the instance
func (instance *Msft_MiCommand) SetPropertyMethodName(value string) (err error) {
	return instance.SetProperty("MethodName", (value))
}

// GetMethodName gets the value of MethodName for the instance
func (instance *Msft_MiCommand) GetPropertyMethodName() (value string, err error) {
	retValue, err := instance.GetProperty("MethodName")
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

// SetNamespaceName sets the value of NamespaceName for the instance
func (instance *Msft_MiCommand) SetPropertyNamespaceName(value string) (err error) {
	return instance.SetProperty("NamespaceName", (value))
}

// GetNamespaceName gets the value of NamespaceName for the instance
func (instance *Msft_MiCommand) GetPropertyNamespaceName() (value string, err error) {
	retValue, err := instance.GetProperty("NamespaceName")
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

// SetParameters sets the value of Parameters for the instance
func (instance *Msft_MiCommand) SetPropertyParameters(value interface{}) (err error) {
	return instance.SetProperty("Parameters", (value))
}

// GetParameters gets the value of Parameters for the instance
func (instance *Msft_MiCommand) GetPropertyParameters() (value interface{}, err error) {
	retValue, err := instance.GetProperty("Parameters")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}
