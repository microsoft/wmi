// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_TaskTimeTrigger struct
type MSFT_TaskTimeTrigger struct {
	*MSFT_TaskTrigger

	//
	RandomDelay string
}

func NewMSFT_TaskTimeTriggerEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskTimeTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskTimeTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}

func NewMSFT_TaskTimeTriggerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskTimeTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskTimeTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}

// SetRandomDelay sets the value of RandomDelay for the instance
func (instance *MSFT_TaskTimeTrigger) SetPropertyRandomDelay(value string) (err error) {
	return instance.SetProperty("RandomDelay", (value))
}

// GetRandomDelay gets the value of RandomDelay for the instance
func (instance *MSFT_TaskTimeTrigger) GetPropertyRandomDelay() (value string, err error) {
	retValue, err := instance.GetProperty("RandomDelay")
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
