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
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_TaskTrigger struct
type MSFT_TaskTrigger struct {
	*cim.WmiInstance

	//
	Enabled bool

	//
	EndBoundary string

	//
	ExecutionTimeLimit string

	//
	Id string

	//
	Repetition MSFT_TaskRepetitionPattern

	//
	StartBoundary string
}

func NewMSFT_TaskTriggerEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskTrigger, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskTrigger{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_TaskTriggerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskTrigger, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskTrigger{
		WmiInstance: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MSFT_TaskTrigger) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSFT_TaskTrigger) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetEndBoundary sets the value of EndBoundary for the instance
func (instance *MSFT_TaskTrigger) SetPropertyEndBoundary(value string) (err error) {
	return instance.SetProperty("EndBoundary", (value))
}

// GetEndBoundary gets the value of EndBoundary for the instance
func (instance *MSFT_TaskTrigger) GetPropertyEndBoundary() (value string, err error) {
	retValue, err := instance.GetProperty("EndBoundary")
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

// SetExecutionTimeLimit sets the value of ExecutionTimeLimit for the instance
func (instance *MSFT_TaskTrigger) SetPropertyExecutionTimeLimit(value string) (err error) {
	return instance.SetProperty("ExecutionTimeLimit", (value))
}

// GetExecutionTimeLimit gets the value of ExecutionTimeLimit for the instance
func (instance *MSFT_TaskTrigger) GetPropertyExecutionTimeLimit() (value string, err error) {
	retValue, err := instance.GetProperty("ExecutionTimeLimit")
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

// SetId sets the value of Id for the instance
func (instance *MSFT_TaskTrigger) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *MSFT_TaskTrigger) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetRepetition sets the value of Repetition for the instance
func (instance *MSFT_TaskTrigger) SetPropertyRepetition(value MSFT_TaskRepetitionPattern) (err error) {
	return instance.SetProperty("Repetition", (value))
}

// GetRepetition gets the value of Repetition for the instance
func (instance *MSFT_TaskTrigger) GetPropertyRepetition() (value MSFT_TaskRepetitionPattern, err error) {
	retValue, err := instance.GetProperty("Repetition")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_TaskRepetitionPattern)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_TaskRepetitionPattern is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_TaskRepetitionPattern(valuetmp)

	return
}

// SetStartBoundary sets the value of StartBoundary for the instance
func (instance *MSFT_TaskTrigger) SetPropertyStartBoundary(value string) (err error) {
	return instance.SetProperty("StartBoundary", (value))
}

// GetStartBoundary gets the value of StartBoundary for the instance
func (instance *MSFT_TaskTrigger) GetPropertyStartBoundary() (value string, err error) {
	retValue, err := instance.GetProperty("StartBoundary")
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
