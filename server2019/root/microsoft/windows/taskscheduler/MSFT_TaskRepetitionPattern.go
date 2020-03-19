// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_TaskRepetitionPattern struct
type MSFT_TaskRepetitionPattern struct {
	*cim.WmiInstance

	//
	Duration string

	//
	Interval string

	//
	StopAtDurationEnd bool
}

func NewMSFT_TaskRepetitionPatternEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskRepetitionPattern, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskRepetitionPattern{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_TaskRepetitionPatternEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskRepetitionPattern, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskRepetitionPattern{
		WmiInstance: tmp,
	}
	return
}

// SetDuration sets the value of Duration for the instance
func (instance *MSFT_TaskRepetitionPattern) SetPropertyDuration(value string) (err error) {
	return instance.SetProperty("Duration", value)
}

// GetDuration gets the value of Duration for the instance
func (instance *MSFT_TaskRepetitionPattern) GetPropertyDuration() (value string, err error) {
	retValue, err := instance.GetProperty("Duration")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterval sets the value of Interval for the instance
func (instance *MSFT_TaskRepetitionPattern) SetPropertyInterval(value string) (err error) {
	return instance.SetProperty("Interval", value)
}

// GetInterval gets the value of Interval for the instance
func (instance *MSFT_TaskRepetitionPattern) GetPropertyInterval() (value string, err error) {
	retValue, err := instance.GetProperty("Interval")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStopAtDurationEnd sets the value of StopAtDurationEnd for the instance
func (instance *MSFT_TaskRepetitionPattern) SetPropertyStopAtDurationEnd(value bool) (err error) {
	return instance.SetProperty("StopAtDurationEnd", value)
}

// GetStopAtDurationEnd gets the value of StopAtDurationEnd for the instance
func (instance *MSFT_TaskRepetitionPattern) GetPropertyStopAtDurationEnd() (value bool, err error) {
	retValue, err := instance.GetProperty("StopAtDurationEnd")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
