// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_NetServiceDifferentPIDConnected struct
type MSFT_NetServiceDifferentPIDConnected struct {
	MSFT_SCMEventLogEvent

	//
	ActualPID uint32

	//
	ExpectedPID uint32

	//
	Service string
}

// SetActualPID sets the value of ActualPID for the instance
func (instance *MSFT_NetServiceDifferentPIDConnected) SetPropertyActualPID(value uint32) (err error) {
	return instance.SetProperty("ActualPID", value)
}

// GetActualPID gets the value of ActualPID for the instance
func (instance *MSFT_NetServiceDifferentPIDConnected) GetPropertyActualPID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActualPID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExpectedPID sets the value of ExpectedPID for the instance
func (instance *MSFT_NetServiceDifferentPIDConnected) SetPropertyExpectedPID(value uint32) (err error) {
	return instance.SetProperty("ExpectedPID", value)
}

// GetExpectedPID gets the value of ExpectedPID for the instance
func (instance *MSFT_NetServiceDifferentPIDConnected) GetPropertyExpectedPID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExpectedPID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetService sets the value of Service for the instance
func (instance *MSFT_NetServiceDifferentPIDConnected) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *MSFT_NetServiceDifferentPIDConnected) GetPropertyService() (value string, err error) {
	retValue, err := instance.GetProperty("Service")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
