// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
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

// Operation struct
type Operation struct {
	*cim.WmiInstance

	// The WS-Adressing action of the request message.
	Action string

	// Indicates that an operation is implemented asynchronously using a Begin<> and End<> method pair in a service contract.
	AsyncPattern bool

	// The bhaviors associated with this operation.
	Behaviors []Behavior

	// True when the operation is a callback operation.
	IsCallback bool

	// Indicates whether the method implements an operation that can initiate a session on the server.
	IsInitiating bool

	// Indicates whether an operation returns a reply message.
	IsOneWay bool

	// Indicates whether an operation returns a reply message.
	IsTerminating bool

	// The method signature of the operation.
	MethodSignature string

	// The name of the operation.
	Name string

	// The types of the parameters of the operation.
	ParameterTypes []string

	// The value of the SOAP action for the reply message of the operation.
	ReplyAction string

	// The return type of the operation.
	ReturnType string
}

func NewOperationEx1(instance *cim.WmiInstance) (newInstance *Operation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Operation{
		WmiInstance: tmp,
	}
	return
}

func NewOperationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Operation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Operation{
		WmiInstance: tmp,
	}
	return
}

// SetAction sets the value of Action for the instance
func (instance *Operation) SetPropertyAction(value string) (err error) {
	return instance.SetProperty("Action", (value))
}

// GetAction gets the value of Action for the instance
func (instance *Operation) GetPropertyAction() (value string, err error) {
	retValue, err := instance.GetProperty("Action")
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

// SetAsyncPattern sets the value of AsyncPattern for the instance
func (instance *Operation) SetPropertyAsyncPattern(value bool) (err error) {
	return instance.SetProperty("AsyncPattern", (value))
}

// GetAsyncPattern gets the value of AsyncPattern for the instance
func (instance *Operation) GetPropertyAsyncPattern() (value bool, err error) {
	retValue, err := instance.GetProperty("AsyncPattern")
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

// SetBehaviors sets the value of Behaviors for the instance
func (instance *Operation) SetPropertyBehaviors(value []Behavior) (err error) {
	return instance.SetProperty("Behaviors", (value))
}

// GetBehaviors gets the value of Behaviors for the instance
func (instance *Operation) GetPropertyBehaviors() (value []Behavior, err error) {
	retValue, err := instance.GetProperty("Behaviors")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Behavior)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Behavior is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Behavior(valuetmp))
	}

	return
}

// SetIsCallback sets the value of IsCallback for the instance
func (instance *Operation) SetPropertyIsCallback(value bool) (err error) {
	return instance.SetProperty("IsCallback", (value))
}

// GetIsCallback gets the value of IsCallback for the instance
func (instance *Operation) GetPropertyIsCallback() (value bool, err error) {
	retValue, err := instance.GetProperty("IsCallback")
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

// SetIsInitiating sets the value of IsInitiating for the instance
func (instance *Operation) SetPropertyIsInitiating(value bool) (err error) {
	return instance.SetProperty("IsInitiating", (value))
}

// GetIsInitiating gets the value of IsInitiating for the instance
func (instance *Operation) GetPropertyIsInitiating() (value bool, err error) {
	retValue, err := instance.GetProperty("IsInitiating")
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

// SetIsOneWay sets the value of IsOneWay for the instance
func (instance *Operation) SetPropertyIsOneWay(value bool) (err error) {
	return instance.SetProperty("IsOneWay", (value))
}

// GetIsOneWay gets the value of IsOneWay for the instance
func (instance *Operation) GetPropertyIsOneWay() (value bool, err error) {
	retValue, err := instance.GetProperty("IsOneWay")
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

// SetIsTerminating sets the value of IsTerminating for the instance
func (instance *Operation) SetPropertyIsTerminating(value bool) (err error) {
	return instance.SetProperty("IsTerminating", (value))
}

// GetIsTerminating gets the value of IsTerminating for the instance
func (instance *Operation) GetPropertyIsTerminating() (value bool, err error) {
	retValue, err := instance.GetProperty("IsTerminating")
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

// SetMethodSignature sets the value of MethodSignature for the instance
func (instance *Operation) SetPropertyMethodSignature(value string) (err error) {
	return instance.SetProperty("MethodSignature", (value))
}

// GetMethodSignature gets the value of MethodSignature for the instance
func (instance *Operation) GetPropertyMethodSignature() (value string, err error) {
	retValue, err := instance.GetProperty("MethodSignature")
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
func (instance *Operation) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *Operation) GetPropertyName() (value string, err error) {
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

// SetParameterTypes sets the value of ParameterTypes for the instance
func (instance *Operation) SetPropertyParameterTypes(value []string) (err error) {
	return instance.SetProperty("ParameterTypes", (value))
}

// GetParameterTypes gets the value of ParameterTypes for the instance
func (instance *Operation) GetPropertyParameterTypes() (value []string, err error) {
	retValue, err := instance.GetProperty("ParameterTypes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetReplyAction sets the value of ReplyAction for the instance
func (instance *Operation) SetPropertyReplyAction(value string) (err error) {
	return instance.SetProperty("ReplyAction", (value))
}

// GetReplyAction gets the value of ReplyAction for the instance
func (instance *Operation) GetPropertyReplyAction() (value string, err error) {
	retValue, err := instance.GetProperty("ReplyAction")
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

// SetReturnType sets the value of ReturnType for the instance
func (instance *Operation) SetPropertyReturnType(value string) (err error) {
	return instance.SetProperty("ReturnType", (value))
}

// GetReturnType gets the value of ReturnType for the instance
func (instance *Operation) GetPropertyReturnType() (value string, err error) {
	retValue, err := instance.GetProperty("ReturnType")
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
