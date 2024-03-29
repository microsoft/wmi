// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.DeliveryOptimization
//////////////////////////////////////////////
package deliveryoptimization

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __MethodInvocationEvent struct
type __MethodInvocationEvent struct {
	*__InstanceOperationEvent

	//
	Method string

	//
	Parameters interface{}

	//
	PreCall bool
}

func New__MethodInvocationEventEx1(instance *cim.WmiInstance) (newInstance *__MethodInvocationEvent, err error) {
	tmp, err := New__InstanceOperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__MethodInvocationEvent{
		__InstanceOperationEvent: tmp,
	}
	return
}

func New__MethodInvocationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__MethodInvocationEvent, err error) {
	tmp, err := New__InstanceOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__MethodInvocationEvent{
		__InstanceOperationEvent: tmp,
	}
	return
}

// SetMethod sets the value of Method for the instance
func (instance *__MethodInvocationEvent) SetPropertyMethod(value string) (err error) {
	return instance.SetProperty("Method", value)
}

// GetMethod gets the value of Method for the instance
func (instance *__MethodInvocationEvent) GetPropertyMethod() (value string, err error) {
	retValue, err := instance.GetProperty("Method")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParameters sets the value of Parameters for the instance
func (instance *__MethodInvocationEvent) SetPropertyParameters(value interface{}) (err error) {
	return instance.SetProperty("Parameters", value)
}

// GetParameters gets the value of Parameters for the instance
func (instance *__MethodInvocationEvent) GetPropertyParameters() (value interface{}, err error) {
	retValue, err := instance.GetProperty("Parameters")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPreCall sets the value of PreCall for the instance
func (instance *__MethodInvocationEvent) SetPropertyPreCall(value bool) (err error) {
	return instance.SetProperty("PreCall", value)
}

// GetPreCall gets the value of PreCall for the instance
func (instance *__MethodInvocationEvent) GetPropertyPreCall() (value bool, err error) {
	retValue, err := instance.GetProperty("PreCall")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
