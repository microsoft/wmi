// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Policy
//////////////////////////////////////////////
package policy

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __TimerEvent struct
type __TimerEvent struct {
	*__Event

	//
	NumFirings uint32

	//
	TimerId string
}

func New__TimerEventEx1(instance *cim.WmiInstance) (newInstance *__TimerEvent, err error) {
	tmp, err := New__EventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__TimerEvent{
		__Event: tmp,
	}
	return
}

func New__TimerEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__TimerEvent, err error) {
	tmp, err := New__EventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__TimerEvent{
		__Event: tmp,
	}
	return
}

// SetNumFirings sets the value of NumFirings for the instance
func (instance *__TimerEvent) SetPropertyNumFirings(value uint32) (err error) {
	return instance.SetProperty("NumFirings", value)
}

// GetNumFirings gets the value of NumFirings for the instance
func (instance *__TimerEvent) GetPropertyNumFirings() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumFirings")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimerId sets the value of TimerId for the instance
func (instance *__TimerEvent) SetPropertyTimerId(value string) (err error) {
	return instance.SetProperty("TimerId", value)
}

// GetTimerId gets the value of TimerId for the instance
func (instance *__TimerEvent) GetPropertyTimerId() (value string, err error) {
	retValue, err := instance.GetProperty("TimerId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
