// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.Security
//////////////////////////////////////////////
package security

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __ConsumerFailureEvent struct
type __ConsumerFailureEvent struct {
	*__EventDroppedEvent

	//
	ErrorCode uint32

	//
	ErrorDescription string

	//
	ErrorObject __ExtendedStatus
}

func New__ConsumerFailureEventEx1(instance *cim.WmiInstance) (newInstance *__ConsumerFailureEvent, err error) {
	tmp, err := New__EventDroppedEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__ConsumerFailureEvent{
		__EventDroppedEvent: tmp,
	}
	return
}

func New__ConsumerFailureEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__ConsumerFailureEvent, err error) {
	tmp, err := New__EventDroppedEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__ConsumerFailureEvent{
		__EventDroppedEvent: tmp,
	}
	return
}

// SetErrorCode sets the value of ErrorCode for the instance
func (instance *__ConsumerFailureEvent) SetPropertyErrorCode(value uint32) (err error) {
	return instance.SetProperty("ErrorCode", value)
}

// GetErrorCode gets the value of ErrorCode for the instance
func (instance *__ConsumerFailureEvent) GetPropertyErrorCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("ErrorCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorDescription sets the value of ErrorDescription for the instance
func (instance *__ConsumerFailureEvent) SetPropertyErrorDescription(value string) (err error) {
	return instance.SetProperty("ErrorDescription", value)
}

// GetErrorDescription gets the value of ErrorDescription for the instance
func (instance *__ConsumerFailureEvent) GetPropertyErrorDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorObject sets the value of ErrorObject for the instance
func (instance *__ConsumerFailureEvent) SetPropertyErrorObject(value __ExtendedStatus) (err error) {
	return instance.SetProperty("ErrorObject", value)
}

// GetErrorObject gets the value of ErrorObject for the instance
func (instance *__ConsumerFailureEvent) GetPropertyErrorObject() (value __ExtendedStatus, err error) {
	retValue, err := instance.GetProperty("ErrorObject")
	if err != nil {
		return
	}
	value, ok := retValue.(__ExtendedStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}
