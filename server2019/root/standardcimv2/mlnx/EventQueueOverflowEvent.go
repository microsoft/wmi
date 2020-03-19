// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
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
	return instance.SetProperty("CurrentQueueSize", value)
}

// GetCurrentQueueSize gets the value of CurrentQueueSize for the instance
func (instance *__EventQueueOverflowEvent) GetPropertyCurrentQueueSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentQueueSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
