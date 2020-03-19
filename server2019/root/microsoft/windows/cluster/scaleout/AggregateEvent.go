// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Cluster.Scaleout
//////////////////////////////////////////////
package scaleout

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __AggregateEvent struct
type __AggregateEvent struct {
	*__IndicationRelated

	//
	NumberOfEvents uint32

	//
	Representative interface{}
}

func New__AggregateEventEx1(instance *cim.WmiInstance) (newInstance *__AggregateEvent, err error) {
	tmp, err := New__IndicationRelatedEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__AggregateEvent{
		__IndicationRelated: tmp,
	}
	return
}

func New__AggregateEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__AggregateEvent, err error) {
	tmp, err := New__IndicationRelatedEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__AggregateEvent{
		__IndicationRelated: tmp,
	}
	return
}

// SetNumberOfEvents sets the value of NumberOfEvents for the instance
func (instance *__AggregateEvent) SetPropertyNumberOfEvents(value uint32) (err error) {
	return instance.SetProperty("NumberOfEvents", value)
}

// GetNumberOfEvents gets the value of NumberOfEvents for the instance
func (instance *__AggregateEvent) GetPropertyNumberOfEvents() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfEvents")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRepresentative sets the value of Representative for the instance
func (instance *__AggregateEvent) SetPropertyRepresentative(value interface{}) (err error) {
	return instance.SetProperty("Representative", value)
}

// GetRepresentative gets the value of Representative for the instance
func (instance *__AggregateEvent) GetPropertyRepresentative() (value interface{}, err error) {
	retValue, err := instance.GetProperty("Representative")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}
