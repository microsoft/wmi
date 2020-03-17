// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_IP4PersistedRouteTable struct
type Win32_IP4PersistedRouteTable struct {
	CIM_LogicalElement

	//
	Destination string

	//
	Mask string

	//
	Metric1 int32

	//
	NextHop string
}

// SetDestination sets the value of Destination for the instance
func (instance *Win32_IP4PersistedRouteTable) SetPropertyDestination(value string) (err error) {
	return instance.SetProperty("Destination", value)
}

// GetDestination gets the value of Destination for the instance
func (instance *Win32_IP4PersistedRouteTable) GetPropertyDestination() (value string, err error) {
	retValue, err := instance.GetProperty("Destination")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMask sets the value of Mask for the instance
func (instance *Win32_IP4PersistedRouteTable) SetPropertyMask(value string) (err error) {
	return instance.SetProperty("Mask", value)
}

// GetMask gets the value of Mask for the instance
func (instance *Win32_IP4PersistedRouteTable) GetPropertyMask() (value string, err error) {
	retValue, err := instance.GetProperty("Mask")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMetric1 sets the value of Metric1 for the instance
func (instance *Win32_IP4PersistedRouteTable) SetPropertyMetric1(value int32) (err error) {
	return instance.SetProperty("Metric1", value)
}

// GetMetric1 gets the value of Metric1 for the instance
func (instance *Win32_IP4PersistedRouteTable) GetPropertyMetric1() (value int32, err error) {
	retValue, err := instance.GetProperty("Metric1")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNextHop sets the value of NextHop for the instance
func (instance *Win32_IP4PersistedRouteTable) SetPropertyNextHop(value string) (err error) {
	return instance.SetProperty("NextHop", value)
}

// GetNextHop gets the value of NextHop for the instance
func (instance *Win32_IP4PersistedRouteTable) GetPropertyNextHop() (value string, err error) {
	retValue, err := instance.GetProperty("NextHop")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
