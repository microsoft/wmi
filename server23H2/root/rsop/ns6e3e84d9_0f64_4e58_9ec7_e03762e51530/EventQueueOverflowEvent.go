// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS6E3E84D9_0F64_4E58_9EC7_E03762E51530
//////////////////////////////////////////////
package ns6e3e84d9_0f64_4e58_9ec7_e03762e51530

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// __EventQueueOverflowEvent struct
type __EventQueueOverflowEvent struct {
	*__EventDroppedEvent

	//
	CurrentQueueSize uint32
}

func New__EventQueueOverflowEventEx1(instance *cim.WmiInstance) (newInstance *__EventQueueOverflowEvent, err error) {
	tmp, err := New__EventDroppedEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__EventQueueOverflowEvent{
		__EventDroppedEvent: tmp,
	}
	return
}

func New__EventQueueOverflowEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__EventQueueOverflowEvent, err error) {
	tmp, err := New__EventDroppedEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__EventQueueOverflowEvent{
		__EventDroppedEvent: tmp,
	}
	return
}

// SetCurrentQueueSize sets the value of CurrentQueueSize for the instance
func (instance *__EventQueueOverflowEvent) SetPropertyCurrentQueueSize(value uint32) (err error) {
	return instance.SetProperty("CurrentQueueSize", (value))
}

// GetCurrentQueueSize gets the value of CurrentQueueSize for the instance
func (instance *__EventQueueOverflowEvent) GetPropertyCurrentQueueSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentQueueSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
