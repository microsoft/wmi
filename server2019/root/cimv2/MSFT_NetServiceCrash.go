// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_NetServiceCrash struct
type MSFT_NetServiceCrash struct {
	MSFT_SCMEventLogEvent

	//
	Action string

	//
	ActionDelay uint32

	//
	ActionType uint32

	//
	Service string

	//
	TimesFailed uint32
}

// SetAction sets the value of Action for the instance
func (instance *MSFT_NetServiceCrash) SetPropertyAction(value string) (err error) {
	return instance.SetProperty("Action", value)
}

// GetAction gets the value of Action for the instance
func (instance *MSFT_NetServiceCrash) GetPropertyAction() (value string, err error) {
	retValue, err := instance.GetProperty("Action")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetActionDelay sets the value of ActionDelay for the instance
func (instance *MSFT_NetServiceCrash) SetPropertyActionDelay(value uint32) (err error) {
	return instance.SetProperty("ActionDelay", value)
}

// GetActionDelay gets the value of ActionDelay for the instance
func (instance *MSFT_NetServiceCrash) GetPropertyActionDelay() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActionDelay")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetActionType sets the value of ActionType for the instance
func (instance *MSFT_NetServiceCrash) SetPropertyActionType(value uint32) (err error) {
	return instance.SetProperty("ActionType", value)
}

// GetActionType gets the value of ActionType for the instance
func (instance *MSFT_NetServiceCrash) GetPropertyActionType() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActionType")
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
func (instance *MSFT_NetServiceCrash) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *MSFT_NetServiceCrash) GetPropertyService() (value string, err error) {
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

// SetTimesFailed sets the value of TimesFailed for the instance
func (instance *MSFT_NetServiceCrash) SetPropertyTimesFailed(value uint32) (err error) {
	return instance.SetProperty("TimesFailed", value)
}

// GetTimesFailed gets the value of TimesFailed for the instance
func (instance *MSFT_NetServiceCrash) GetPropertyTimesFailed() (value uint32, err error) {
	retValue, err := instance.GetProperty("TimesFailed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
