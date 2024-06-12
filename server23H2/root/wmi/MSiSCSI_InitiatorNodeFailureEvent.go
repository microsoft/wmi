// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSiSCSI_InitiatorNodeFailureEvent struct
type MSiSCSI_InitiatorNodeFailureEvent struct {
	*__ExtrinsicEvent

	//
	Active bool

	// Timestamp denoting time failure occured
	FailureTime uint64

	// **typedef** Types of initiator node failure
	FailureType InitiatorNodeFailureEvent_FailureType

	//
	InstanceName string

	// Network address of target involved in failure
	TargetFailureAddr ISCSI_IP_Address

	// Name of target involved in failure
	TargetFailureName string
}

func NewMSiSCSI_InitiatorNodeFailureEventEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_InitiatorNodeFailureEvent, err error) {
	tmp, err := New__ExtrinsicEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_InitiatorNodeFailureEvent{
		__ExtrinsicEvent: tmp,
	}
	return
}

func NewMSiSCSI_InitiatorNodeFailureEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_InitiatorNodeFailureEvent, err error) {
	tmp, err := New__ExtrinsicEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_InitiatorNodeFailureEvent{
		__ExtrinsicEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_InitiatorNodeFailureEvent) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_InitiatorNodeFailureEvent) GetPropertyActive() (value bool, err error) {
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

// SetFailureTime sets the value of FailureTime for the instance
func (instance *MSiSCSI_InitiatorNodeFailureEvent) SetPropertyFailureTime(value uint64) (err error) {
	return instance.SetProperty("FailureTime", (value))
}

// GetFailureTime gets the value of FailureTime for the instance
func (instance *MSiSCSI_InitiatorNodeFailureEvent) GetPropertyFailureTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("FailureTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFailureType sets the value of FailureType for the instance
func (instance *MSiSCSI_InitiatorNodeFailureEvent) SetPropertyFailureType(value InitiatorNodeFailureEvent_FailureType) (err error) {
	return instance.SetProperty("FailureType", (value))
}

// GetFailureType gets the value of FailureType for the instance
func (instance *MSiSCSI_InitiatorNodeFailureEvent) GetPropertyFailureType() (value InitiatorNodeFailureEvent_FailureType, err error) {
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

	value = InitiatorNodeFailureEvent_FailureType(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_InitiatorNodeFailureEvent) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_InitiatorNodeFailureEvent) GetPropertyInstanceName() (value string, err error) {
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

// SetTargetFailureAddr sets the value of TargetFailureAddr for the instance
func (instance *MSiSCSI_InitiatorNodeFailureEvent) SetPropertyTargetFailureAddr(value ISCSI_IP_Address) (err error) {
	return instance.SetProperty("TargetFailureAddr", (value))
}

// GetTargetFailureAddr gets the value of TargetFailureAddr for the instance
func (instance *MSiSCSI_InitiatorNodeFailureEvent) GetPropertyTargetFailureAddr() (value ISCSI_IP_Address, err error) {
	retValue, err := instance.GetProperty("TargetFailureAddr")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ISCSI_IP_Address)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ISCSI_IP_Address is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ISCSI_IP_Address(valuetmp)

	return
}

// SetTargetFailureName sets the value of TargetFailureName for the instance
func (instance *MSiSCSI_InitiatorNodeFailureEvent) SetPropertyTargetFailureName(value string) (err error) {
	return instance.SetProperty("TargetFailureName", (value))
}

// GetTargetFailureName gets the value of TargetFailureName for the instance
func (instance *MSiSCSI_InitiatorNodeFailureEvent) GetPropertyTargetFailureName() (value string, err error) {
	retValue, err := instance.GetProperty("TargetFailureName")
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
