// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

// MSFT_TaskEventTrigger struct
type MSFT_TaskEventTrigger struct {
	MSFT_TaskTrigger

	//
	Delay string

	//
	Subscription string

	//
	ValueQueries []MSFT_TaskNamedValue
}

// SetDelay sets the value of Delay for the instance
func (instance *MSFT_TaskEventTrigger) SetPropertyDelay(value string) (err error) {
	return instance.SetProperty("Delay", value)
}

// GetDelay gets the value of Delay for the instance
func (instance *MSFT_TaskEventTrigger) GetPropertyDelay() (value string, err error) {
	retValue, err := instance.GetProperty("Delay")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSubscription sets the value of Subscription for the instance
func (instance *MSFT_TaskEventTrigger) SetPropertySubscription(value string) (err error) {
	return instance.SetProperty("Subscription", value)
}

// GetSubscription gets the value of Subscription for the instance
func (instance *MSFT_TaskEventTrigger) GetPropertySubscription() (value string, err error) {
	retValue, err := instance.GetProperty("Subscription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValueQueries sets the value of ValueQueries for the instance
func (instance *MSFT_TaskEventTrigger) SetPropertyValueQueries(value []MSFT_TaskNamedValue) (err error) {
	return instance.SetProperty("ValueQueries", value)
}

// GetValueQueries gets the value of ValueQueries for the instance
func (instance *MSFT_TaskEventTrigger) GetPropertyValueQueries() (value []MSFT_TaskNamedValue, err error) {
	retValue, err := instance.GetProperty("ValueQueries")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_TaskNamedValue)
	if !ok {
		// TODO: Set an error
	}
	return
}
