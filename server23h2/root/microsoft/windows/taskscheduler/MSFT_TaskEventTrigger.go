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

// MSFT_TaskEventTrigger struct
type MSFT_TaskEventTrigger struct {
	*MSFT_TaskTrigger

	//
	Delay string

	//
	Subscription string

	//
	ValueQueries []MSFT_TaskNamedValue
}

func NewMSFT_TaskEventTriggerEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskEventTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskEventTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}

func NewMSFT_TaskEventTriggerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskEventTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskEventTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}

// SetDelay sets the value of Delay for the instance
func (instance *MSFT_TaskEventTrigger) SetPropertyDelay(value string) (err error) {
	return instance.SetProperty("Delay", (value))
}

// GetDelay gets the value of Delay for the instance
func (instance *MSFT_TaskEventTrigger) GetPropertyDelay() (value string, err error) {
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

// SetSubscription sets the value of Subscription for the instance
func (instance *MSFT_TaskEventTrigger) SetPropertySubscription(value string) (err error) {
	return instance.SetProperty("Subscription", (value))
}

// GetSubscription gets the value of Subscription for the instance
func (instance *MSFT_TaskEventTrigger) GetPropertySubscription() (value string, err error) {
	retValue, err := instance.GetProperty("Subscription")
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

// SetValueQueries sets the value of ValueQueries for the instance
func (instance *MSFT_TaskEventTrigger) SetPropertyValueQueries(value []MSFT_TaskNamedValue) (err error) {
	return instance.SetProperty("ValueQueries", (value))
}

// GetValueQueries gets the value of ValueQueries for the instance
func (instance *MSFT_TaskEventTrigger) GetPropertyValueQueries() (value []MSFT_TaskNamedValue, err error) {
	retValue, err := instance.GetProperty("ValueQueries")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_TaskNamedValue)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_TaskNamedValue is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_TaskNamedValue(valuetmp))
	}

	return
}
