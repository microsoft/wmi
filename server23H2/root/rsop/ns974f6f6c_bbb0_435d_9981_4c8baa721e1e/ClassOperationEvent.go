// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS974F6F6C_BBB0_435D_9981_4C8BAA721E1E
//////////////////////////////////////////////
package ns974f6f6c_bbb0_435d_9981_4c8baa721e1e

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
