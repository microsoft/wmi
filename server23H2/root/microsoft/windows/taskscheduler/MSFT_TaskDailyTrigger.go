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

// MSFT_TaskDailyTrigger struct
type MSFT_TaskDailyTrigger struct {
	*MSFT_TaskTrigger

	//
	DaysInterval int16

	//
	RandomDelay string
}

func NewMSFT_TaskDailyTriggerEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskDailyTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskDailyTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}

func NewMSFT_TaskDailyTriggerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskDailyTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskDailyTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}

// SetDaysInterval sets the value of DaysInterval for the instance
func (instance *MSFT_TaskDailyTrigger) SetPropertyDaysInterval(value int16) (err error) {
	return instance.SetProperty("DaysInterval", (value))
}

// GetDaysInterval gets the value of DaysInterval for the instance
func (instance *MSFT_TaskDailyTrigger) GetPropertyDaysInterval() (value int16, err error) {
	retValue, err := instance.GetProperty("DaysInterval")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int16(valuetmp)

	return
}

// SetRandomDelay sets the value of RandomDelay for the instance
func (instance *MSFT_TaskDailyTrigger) SetPropertyRandomDelay(value string) (err error) {
	return instance.SetProperty("RandomDelay", (value))
}

// GetRandomDelay gets the value of RandomDelay for the instance
func (instance *MSFT_TaskDailyTrigger) GetPropertyRandomDelay() (value string, err error) {
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
