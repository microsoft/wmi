// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_NetServiceStartFailedNone struct
type MSFT_NetServiceStartFailedNone struct {
	MSFT_SCMEventLogEvent

	//
	NonExistingService string

	//
	Service string
}

// SetNonExistingService sets the value of NonExistingService for the instance
func (instance *MSFT_NetServiceStartFailedNone) SetPropertyNonExistingService(value string) (err error) {
	return instance.SetProperty("NonExistingService", value)
}

// GetNonExistingService gets the value of NonExistingService for the instance
func (instance *MSFT_NetServiceStartFailedNone) GetPropertyNonExistingService() (value string, err error) {
	retValue, err := instance.GetProperty("NonExistingService")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetService sets the value of Service for the instance
func (instance *MSFT_NetServiceStartFailedNone) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *MSFT_NetServiceStartFailedNone) GetPropertyService() (value string, err error) {
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
