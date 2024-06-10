// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSiSCSI_InitiatorInstanceFailureEvent struct
type MSiSCSI_InitiatorInstanceFailureEvent struct {
	*__ExtrinsicEvent

	//
	Active bool

	// **typedef** Type of failure
	FailureType InitiatorInstanceFailureEvent_FailureType

	//
	InstanceName string

	// Name of target involved in failure
	RemoteNodeName string
}

func NewMSiSCSI_InitiatorInstanceFailureEventEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_InitiatorInstanceFailureEvent, err error) {
	tmp, err := New__ExtrinsicEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_InitiatorInstanceFailureEvent{
		__ExtrinsicEvent: tmp,
	}
	return
}

func NewMSiSCSI_InitiatorInstanceFailureEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_InitiatorInstanceFailureEvent, err error) {
	tmp, err := New__ExtrinsicEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_InitiatorInstanceFailureEvent{
		__ExtrinsicEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_InitiatorInstanceFailureEvent) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_InitiatorInstanceFailureEvent) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetFailureType sets the value of FailureType for the instance
func (instance *MSiSCSI_InitiatorInstanceFailureEvent) SetPropertyFailureType(value InitiatorInstanceFailureEvent_FailureType) (err error) {
	return instance.SetProperty("FailureType", (value))
}

// GetFailureType gets the value of FailureType for the instance
func (instance *MSiSCSI_InitiatorInstanceFailureEvent) GetPropertyFailureType() (value InitiatorInstanceFailureEvent_FailureType, err error) {
	retValue, err := instance.GetProperty("FailureType")
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

	value = InitiatorInstanceFailureEvent_FailureType(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_InitiatorInstanceFailureEvent) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_InitiatorInstanceFailureEvent) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetRemoteNodeName sets the value of RemoteNodeName for the instance
func (instance *MSiSCSI_InitiatorInstanceFailureEvent) SetPropertyRemoteNodeName(value string) (err error) {
	return instance.SetProperty("RemoteNodeName", (value))
}

// GetRemoteNodeName gets the value of RemoteNodeName for the instance
func (instance *MSiSCSI_InitiatorInstanceFailureEvent) GetPropertyRemoteNodeName() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteNodeName")
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
