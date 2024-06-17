// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// Endpoint struct
type Endpoint struct {
	*cim.WmiInstance

	// A Uri that contains the address of the endpoint.
	Address string

	//  The collection of address headers attached to this endpoint.
	AddressHeaders []string

	// The identity of the endpoint.
	AddressIdentity string

	// The appdomain id of the appdomain that hosts the endpoint.
	AppDomainId int32

	// The collection of behaviors implemented by this endpoint.
	Behaviors []Behavior

	// The binding used by this endpoint.
	Binding Binding

	// The contract this endpoint is exposing.
	Contract Contract

	// A string that specifies which contract this endpoint is exposing.
	ContractName string

	// The name of the instance of performance counters of the endpoint.
	CounterInstanceName string

	// The Uri the endpoint listens on.
	ListenUri string

	// The unique name of this endpoint.
	Name string

	// The process Id of the process that hosts the endpoint.
	ProcessId int32
}

func NewEndpointEx1(instance *cim.WmiInstance) (newInstance *Endpoint, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Endpoint{
		WmiInstance: tmp,
	}
	return
}

func NewEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Endpoint, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Endpoint{
		WmiInstance: tmp,
	}
	return
}

// SetAddress sets the value of Address for the instance
func (instance *Endpoint) SetPropertyAddress(value string) (err error) {
	return instance.SetProperty("Address", (value))
}

// GetAddress gets the value of Address for the instance
func (instance *Endpoint) GetPropertyAddress() (value string, err error) {
	retValue, err := instance.GetProperty("Address")
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

// SetAddressHeaders sets the value of AddressHeaders for the instance
func (instance *Endpoint) SetPropertyAddressHeaders(value []string) (err error) {
	return instance.SetProperty("AddressHeaders", (value))
}

// GetAddressHeaders gets the value of AddressHeaders for the instance
func (instance *Endpoint) GetPropertyAddressHeaders() (value []string, err error) {
	retValue, err := instance.GetProperty("AddressHeaders")
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

// SetAddressIdentity sets the value of AddressIdentity for the instance
func (instance *Endpoint) SetPropertyAddressIdentity(value string) (err error) {
	return instance.SetProperty("AddressIdentity", (value))
}

// GetAddressIdentity gets the value of AddressIdentity for the instance
func (instance *Endpoint) GetPropertyAddressIdentity() (value string, err error) {
	retValue, err := instance.GetProperty("AddressIdentity")
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

// SetAppDomainId sets the value of AppDomainId for the instance
func (instance *Endpoint) SetPropertyAppDomainId(value int32) (err error) {
	return instance.SetProperty("AppDomainId", (value))
}

// GetAppDomainId gets the value of AppDomainId for the instance
func (instance *Endpoint) GetPropertyAppDomainId() (value int32, err error) {
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
func (instance *Endpoint) SetPropertyBehaviors(value []Behavior) (err error) {
	return instance.SetProperty("Behaviors", (value))
}

// GetBehaviors gets the value of Behaviors for the instance
func (instance *Endpoint) GetPropertyBehaviors() (value []Behavior, err error) {
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

// SetBinding sets the value of Binding for the instance
func (instance *Endpoint) SetPropertyBinding(value Binding) (err error) {
	return instance.SetProperty("Binding", (value))
}

// GetBinding gets the value of Binding for the instance
func (instance *Endpoint) GetPropertyBinding() (value Binding, err error) {
	retValue, err := instance.GetProperty("Binding")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Binding)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Binding is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Binding(valuetmp)

	return
}

// SetContract sets the value of Contract for the instance
func (instance *Endpoint) SetPropertyContract(value Contract) (err error) {
	return instance.SetProperty("Contract", (value))
}

// GetContract gets the value of Contract for the instance
func (instance *Endpoint) GetPropertyContract() (value Contract, err error) {
	retValue, err := instance.GetProperty("Contract")
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

// SetContractName sets the value of ContractName for the instance
func (instance *Endpoint) SetPropertyContractName(value string) (err error) {
	return instance.SetProperty("ContractName", (value))
}

// GetContractName gets the value of ContractName for the instance
func (instance *Endpoint) GetPropertyContractName() (value string, err error) {
	retValue, err := instance.GetProperty("ContractName")
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

// SetCounterInstanceName sets the value of CounterInstanceName for the instance
func (instance *Endpoint) SetPropertyCounterInstanceName(value string) (err error) {
	return instance.SetProperty("CounterInstanceName", (value))
}

// GetCounterInstanceName gets the value of CounterInstanceName for the instance
func (instance *Endpoint) GetPropertyCounterInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("CounterInstanceName")
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

// SetListenUri sets the value of ListenUri for the instance
func (instance *Endpoint) SetPropertyListenUri(value string) (err error) {
	return instance.SetProperty("ListenUri", (value))
}

// GetListenUri gets the value of ListenUri for the instance
func (instance *Endpoint) GetPropertyListenUri() (value string, err error) {
	retValue, err := instance.GetProperty("ListenUri")
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
func (instance *Endpoint) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *Endpoint) GetPropertyName() (value string, err error) {
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *Endpoint) SetPropertyProcessId(value int32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *Endpoint) GetPropertyProcessId() (value int32, err error) {
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

// Retrieves operation performance counter instance name

// <param name="Operation" type="string "></param>

// <param name="ReturnValue" type="string "></param>
func (instance *Endpoint) GetOperationCounterInstanceName( /* IN */ Operation string) (result string, err error) {
	retVal, err := instance.InvokeMethodWithReturn("GetOperationCounterInstanceName", Operation)
	if err != nil {
		return
	}
	result = string(retVal)
	return

}
