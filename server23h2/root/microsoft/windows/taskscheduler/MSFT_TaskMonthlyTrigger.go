// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_TaskMonthlyTrigger struct
type MSFT_TaskMonthlyTrigger struct {
	*MSFT_TaskTrigger

	//
	DaysOfMonth uint16

	//
	MonthOfYear uint16

	//
	RandomDelay string

	//
	RunOnLastDayOfMonth bool
}

func NewMSFT_TaskMonthlyTriggerEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskMonthlyTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskMonthlyTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}

func NewMSFT_TaskMonthlyTriggerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskMonthlyTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskMonthlyTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}

// SetDaysOfMonth sets the value of DaysOfMonth for the instance
func (instance *MSFT_TaskMonthlyTrigger) SetPropertyDaysOfMonth(value uint16) (err error) {
	return instance.SetProperty("DaysOfMonth", (value))
}

// GetDaysOfMonth gets the value of DaysOfMonth for the instance
func (instance *MSFT_TaskMonthlyTrigger) GetPropertyDaysOfMonth() (value uint16, err error) {
	retValue, err := instance.GetProperty("DaysOfMonth")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetMonthOfYear sets the value of MonthOfYear for the instance
func (instance *MSFT_TaskMonthlyTrigger) SetPropertyMonthOfYear(value uint16) (err error) {
	return instance.SetProperty("MonthOfYear", (value))
}

// GetMonthOfYear gets the value of MonthOfYear for the instance
func (instance *MSFT_TaskMonthlyTrigger) GetPropertyMonthOfYear() (value uint16, err error) {
	retValue, err := instance.GetProperty("MonthOfYear")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetRandomDelay sets the value of RandomDelay for the instance
func (instance *MSFT_TaskMonthlyTrigger) SetPropertyRandomDelay(value string) (err error) {
	return instance.SetProperty("RandomDelay", (value))
}

// GetRandomDelay gets the value of RandomDelay for the instance
func (instance *MSFT_TaskMonthlyTrigger) GetPropertyRandomDelay() (value string, err error) {
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

// SetRunOnLastDayOfMonth sets the value of RunOnLastDayOfMonth for the instance
func (instance *MSFT_TaskMonthlyTrigger) SetPropertyRunOnLastDayOfMonth(value bool) (err error) {
	return instance.SetProperty("RunOnLastDayOfMonth", (value))
}

// GetRunOnLastDayOfMonth gets the value of RunOnLastDayOfMonth for the instance
func (instance *MSFT_TaskMonthlyTrigger) GetPropertyRunOnLastDayOfMonth() (value bool, err error) {
	retValue, err := instance.GetProperty("RunOnLastDayOfMonth")
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
