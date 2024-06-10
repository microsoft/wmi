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
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ServiceThrottlingBehavior struct
type ServiceThrottlingBehavior struct {
	*Behavior

	// The maximum number of messages actively processing across all dispatcher objects in a ServiceHost.
	MaxConcurrentCalls int32

	// The maximum number of service objects that can execute at one time.
	MaxConcurrentInstances int32

	// The maximum number of sessions a host can accept at one time.
	MaxConcurrentSessions int32
}

func NewServiceThrottlingBehaviorEx1(instance *cim.WmiInstance) (newInstance *ServiceThrottlingBehavior, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServiceThrottlingBehavior{
		Behavior: tmp,
	}
	return
}

func NewServiceThrottlingBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServiceThrottlingBehavior, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServiceThrottlingBehavior{
		Behavior: tmp,
	}
	return
}

// SetMaxConcurrentCalls sets the value of MaxConcurrentCalls for the instance
func (instance *ServiceThrottlingBehavior) SetPropertyMaxConcurrentCalls(value int32) (err error) {
	return instance.SetProperty("MaxConcurrentCalls", (value))
}

// GetMaxConcurrentCalls gets the value of MaxConcurrentCalls for the instance
func (instance *ServiceThrottlingBehavior) GetPropertyMaxConcurrentCalls() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxConcurrentCalls")
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

// SetMaxConcurrentInstances sets the value of MaxConcurrentInstances for the instance
func (instance *ServiceThrottlingBehavior) SetPropertyMaxConcurrentInstances(value int32) (err error) {
	return instance.SetProperty("MaxConcurrentInstances", (value))
}

// GetMaxConcurrentInstances gets the value of MaxConcurrentInstances for the instance
func (instance *ServiceThrottlingBehavior) GetPropertyMaxConcurrentInstances() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxConcurrentInstances")
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

// SetMaxConcurrentSessions sets the value of MaxConcurrentSessions for the instance
func (instance *ServiceThrottlingBehavior) SetPropertyMaxConcurrentSessions(value int32) (err error) {
	return instance.SetProperty("MaxConcurrentSessions", (value))
}

// GetMaxConcurrentSessions gets the value of MaxConcurrentSessions for the instance
func (instance *ServiceThrottlingBehavior) GetPropertyMaxConcurrentSessions() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxConcurrentSessions")
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
