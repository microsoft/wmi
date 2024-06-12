// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS77DB142C_6A49_489B_8538_882E7276E951
//////////////////////////////////////////////
package ns77db142c_6a49_489b_8538_882e7276e951

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// __InstanceModificationEvent struct
type __InstanceModificationEvent struct {
	*__InstanceOperationEvent

	//
	PreviousInstance interface{}
}

func New__InstanceModificationEventEx1(instance *cim.WmiInstance) (newInstance *__InstanceModificationEvent, err error) {
	tmp, err := New__InstanceOperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__InstanceModificationEvent{
		__InstanceOperationEvent: tmp,
	}
	return
}

func New__InstanceModificationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__InstanceModificationEvent, err error) {
	tmp, err := New__InstanceOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__InstanceModificationEvent{
		__InstanceOperationEvent: tmp,
	}
	return
}

// SetPreviousInstance sets the value of PreviousInstance for the instance
func (instance *__InstanceModificationEvent) SetPropertyPreviousInstance(value interface{}) (err error) {
	return instance.SetProperty("PreviousInstance", (value))
}

// GetPreviousInstance gets the value of PreviousInstance for the instance
func (instance *__InstanceModificationEvent) GetPropertyPreviousInstance() (value interface{}, err error) {
	retValue, err := instance.GetProperty("PreviousInstance")
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
