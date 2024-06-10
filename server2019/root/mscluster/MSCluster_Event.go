// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.MSCluster
//
// ////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_Event struct
type MSCluster_Event struct {
	*__ExtrinsicEvent

	//
	EventObjectName string

	//
	EventObjectPath string

	//
	EventObjectType uint32

	//
	EventTypeMajor uint32

	//
	EventTypeMinor uint32
}

func NewMSCluster_EventEx1(instance *cim.WmiInstance) (newInstance *MSCluster_Event, err error) {
	tmp, err := New__ExtrinsicEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_Event{
		__ExtrinsicEvent: tmp,
	}
	return
}

func NewMSCluster_EventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_Event, err error) {
	tmp, err := New__ExtrinsicEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_Event{
		__ExtrinsicEvent: tmp,
	}
	return
}

// SetEventObjectName sets the value of EventObjectName for the instance
func (instance *MSCluster_Event) SetPropertyEventObjectName(value string) (err error) {
	return instance.SetProperty("EventObjectName", (value))
}

// GetEventObjectName gets the value of EventObjectName for the instance
func (instance *MSCluster_Event) GetPropertyEventObjectName() (value string, err error) {
	retValue, err := instance.GetProperty("EventObjectName")
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

// SetEventObjectPath sets the value of EventObjectPath for the instance
func (instance *MSCluster_Event) SetPropertyEventObjectPath(value string) (err error) {
	return instance.SetProperty("EventObjectPath", (value))
}

// GetEventObjectPath gets the value of EventObjectPath for the instance
func (instance *MSCluster_Event) GetPropertyEventObjectPath() (value string, err error) {
	retValue, err := instance.GetProperty("EventObjectPath")
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

// SetEventObjectType sets the value of EventObjectType for the instance
func (instance *MSCluster_Event) SetPropertyEventObjectType(value uint32) (err error) {
	return instance.SetProperty("EventObjectType", (value))
}

// GetEventObjectType gets the value of EventObjectType for the instance
func (instance *MSCluster_Event) GetPropertyEventObjectType() (value uint32, err error) {
	retValue, err := instance.GetProperty("EventObjectType")
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

// SetEventTypeMajor sets the value of EventTypeMajor for the instance
func (instance *MSCluster_Event) SetPropertyEventTypeMajor(value uint32) (err error) {
	return instance.SetProperty("EventTypeMajor", (value))
}

// GetEventTypeMajor gets the value of EventTypeMajor for the instance
func (instance *MSCluster_Event) GetPropertyEventTypeMajor() (value uint32, err error) {
	retValue, err := instance.GetProperty("EventTypeMajor")
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

// SetEventTypeMinor sets the value of EventTypeMinor for the instance
func (instance *MSCluster_Event) SetPropertyEventTypeMinor(value uint32) (err error) {
	return instance.SetProperty("EventTypeMinor", (value))
}

// GetEventTypeMinor gets the value of EventTypeMinor for the instance
func (instance *MSCluster_Event) GetPropertyEventTypeMinor() (value uint32, err error) {
	retValue, err := instance.GetProperty("EventTypeMinor")
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
