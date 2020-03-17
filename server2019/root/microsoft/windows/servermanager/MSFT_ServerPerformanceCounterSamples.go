// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ServerPerformanceCounterSamples struct
type MSFT_ServerPerformanceCounterSamples struct {
	cim.WmiInstance

	//
	CounterPaths []string

	//
	Timestamps []string

	//
	Values []string
}

// SetCounterPaths sets the value of CounterPaths for the instance
func (instance *MSFT_ServerPerformanceCounterSamples) SetPropertyCounterPaths(value []string) (err error) {
	return instance.SetProperty("CounterPaths", value)
}

// GetCounterPaths gets the value of CounterPaths for the instance
func (instance *MSFT_ServerPerformanceCounterSamples) GetPropertyCounterPaths() (value []string, err error) {
	retValue, err := instance.GetProperty("CounterPaths")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimestamps sets the value of Timestamps for the instance
func (instance *MSFT_ServerPerformanceCounterSamples) SetPropertyTimestamps(value []string) (err error) {
	return instance.SetProperty("Timestamps", value)
}

// GetTimestamps gets the value of Timestamps for the instance
func (instance *MSFT_ServerPerformanceCounterSamples) GetPropertyTimestamps() (value []string, err error) {
	retValue, err := instance.GetProperty("Timestamps")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValues sets the value of Values for the instance
func (instance *MSFT_ServerPerformanceCounterSamples) SetPropertyValues(value []string) (err error) {
	return instance.SetProperty("Values", value)
}

// GetValues gets the value of Values for the instance
func (instance *MSFT_ServerPerformanceCounterSamples) GetPropertyValues() (value []string, err error) {
	retValue, err := instance.GetProperty("Values")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
