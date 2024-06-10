// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.ServiceModel
//
// ////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Contract struct
type Contract struct {
	*cim.WmiInstance

	// The appdomain id of the appdomain that hosts the contract.
	AppDomainId int32

	// The behaviors associated with this contract.
	Behaviors []Behavior

	// The type of callback when the contract is a duplex contract.
	CallbackContract Contract

	// The name of the contract in WSDL.
	Name string

	// The namespace of the <portType> element in WSDL.
	Namespace string

	// The operations of this contract.
	Operations []Operation

	// The process Id of the process that hosts the contract.
	ProcessId int32

	// Indicates whether the contract requires the binding associated with this contract to use channel sessions.
	SessionMode string

	// The type of the contract.
	Type string
}

func NewContractEx1(instance *cim.WmiInstance) (newInstance *Contract, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Contract{
		WmiInstance: tmp,
	}
	return
}

func NewContractEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Contract, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Contract{
		WmiInstance: tmp,
	}
	return
}

// SetAppDomainId sets the value of AppDomainId for the instance
func (instance *Contract) SetPropertyAppDomainId(value int32) (err error) {
	return instance.SetProperty("AppDomainId", (value))
}

// GetAppDomainId gets the value of AppDomainId for the instance
func (instance *Contract) GetPropertyAppDomainId() (value int32, err error) {
	retValue, err := instance.GetProperty("AppDomainId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetBehaviors sets the value of Behaviors for the instance
func (instance *Contract) SetPropertyBehaviors(value []Behavior) (err error) {
	return instance.SetProperty("Behaviors", (value))
}

// GetBehaviors gets the value of Behaviors for the instance
func (instance *Contract) GetPropertyBehaviors() (value []Behavior, err error) {
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

// SetCallbackContract sets the value of CallbackContract for the instance
func (instance *Contract) SetPropertyCallbackContract(value Contract) (err error) {
	return instance.SetProperty("CallbackContract", (value))
}

// GetCallbackContract gets the value of CallbackContract for the instance
func (instance *Contract) GetPropertyCallbackContract() (value Contract, err error) {
	retValue, err := instance.GetProperty("CallbackContract")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Contract)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Contract is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Contract(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *Contract) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *Contract) GetPropertyName() (value string, err error) {
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
func (instance *Contract) SetPropertyNamespace(value string) (err error) {
	return instance.SetProperty("Namespace", (value))
}

// GetNamespace gets the value of Namespace for the instance
func (instance *Contract) GetPropertyNamespace() (value string, err error) {
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

// SetOperations sets the value of Operations for the instance
func (instance *Contract) SetPropertyOperations(value []Operation) (err error) {
	return instance.SetProperty("Operations", (value))
}

// GetOperations gets the value of Operations for the instance
func (instance *Contract) GetPropertyOperations() (value []Operation, err error) {
	retValue, err := instance.GetProperty("Operations")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Operation)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Operation is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Operation(valuetmp))
	}

	return
}

// SetProcessId sets the value of ProcessId for the instance
func (instance *Contract) SetPropertyProcessId(value int32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *Contract) GetPropertyProcessId() (value int32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetSessionMode sets the value of SessionMode for the instance
func (instance *Contract) SetPropertySessionMode(value string) (err error) {
	return instance.SetProperty("SessionMode", (value))
}

// GetSessionMode gets the value of SessionMode for the instance
func (instance *Contract) GetPropertySessionMode() (value string, err error) {
	retValue, err := instance.GetProperty("SessionMode")
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

// SetType sets the value of Type for the instance
func (instance *Contract) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *Contract) GetPropertyType() (value string, err error) {
	retValue, err := instance.GetProperty("Type")
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
