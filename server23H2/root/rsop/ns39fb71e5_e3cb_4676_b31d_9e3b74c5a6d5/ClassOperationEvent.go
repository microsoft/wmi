// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NS39FB71E5_E3CB_4676_B31D_9E3B74C5A6D5
//////////////////////////////////////////////
package ns39fb71e5_e3cb_4676_b31d_9e3b74c5a6d5

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// __ClassOperationEvent struct
type __ClassOperationEvent struct {
	*__Event

	//
	TargetClass interface{}
}

func New__ClassOperationEventEx1(instance *cim.WmiInstance) (newInstance *__ClassOperationEvent, err error) {
	tmp, err := New__EventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__ClassOperationEvent{
		__Event: tmp,
	}
	return
}

func New__ClassOperationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__ClassOperationEvent, err error) {
	tmp, err := New__EventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__ClassOperationEvent{
		__Event: tmp,
	}
	return
}

// SetTargetClass sets the value of TargetClass for the instance
func (instance *__ClassOperationEvent) SetPropertyTargetClass(value interface{}) (err error) {
	return instance.SetProperty("TargetClass", (value))
}

// GetTargetClass gets the value of TargetClass for the instance
func (instance *__ClassOperationEvent) GetPropertyTargetClass() (value interface{}, err error) {
	retValue, err := instance.GetProperty("TargetClass")
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