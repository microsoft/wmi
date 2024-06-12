// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WorkflowRuntimeBehavior struct
type WorkflowRuntimeBehavior struct {
	*Behavior

	// Specifies the interval after which idle workflow instances in memory are terminated.
	CachedInstanceExpiration string
}

func NewWorkflowRuntimeBehaviorEx1(instance *cim.WmiInstance) (newInstance *WorkflowRuntimeBehavior, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WorkflowRuntimeBehavior{
		Behavior: tmp,
	}
	return
}

func NewWorkflowRuntimeBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WorkflowRuntimeBehavior, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WorkflowRuntimeBehavior{
		Behavior: tmp,
	}
	return
}

// SetCachedInstanceExpiration sets the value of CachedInstanceExpiration for the instance
func (instance *WorkflowRuntimeBehavior) SetPropertyCachedInstanceExpiration(value string) (err error) {
	return instance.SetProperty("CachedInstanceExpiration", (value))
}

// GetCachedInstanceExpiration gets the value of CachedInstanceExpiration for the instance
func (instance *WorkflowRuntimeBehavior) GetPropertyCachedInstanceExpiration() (value string, err error) {
	retValue, err := instance.GetProperty("CachedInstanceExpiration")
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
