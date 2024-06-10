// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.TaskScheduler
//
// ////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_TaskRegistrationTrigger struct
type MSFT_TaskRegistrationTrigger struct {
	*MSFT_TaskTrigger

	//
	Delay string
}

func NewMSFT_TaskRegistrationTriggerEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskRegistrationTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskRegistrationTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}

func NewMSFT_TaskRegistrationTriggerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskRegistrationTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskRegistrationTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}

// SetDelay sets the value of Delay for the instance
func (instance *MSFT_TaskRegistrationTrigger) SetPropertyDelay(value string) (err error) {
	return instance.SetProperty("Delay", (value))
}

// GetDelay gets the value of Delay for the instance
func (instance *MSFT_TaskRegistrationTrigger) GetPropertyDelay() (value string, err error) {
	retValue, err := instance.GetProperty("Delay")
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
