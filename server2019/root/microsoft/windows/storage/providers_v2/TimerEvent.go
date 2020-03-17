// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

// __TimerEvent struct
type __TimerEvent struct {
	__Event

	//
	NumFirings uint32

	//
	TimerId string
}

// SetNumFirings sets the value of NumFirings for the instance
func (instance *__TimerEvent) SetPropertyNumFirings(value uint32) (err error) {
	return instance.SetProperty("NumFirings", value)
}

// GetNumFirings gets the value of NumFirings for the instance
func (instance *__TimerEvent) GetPropertyNumFirings() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumFirings")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimerId sets the value of TimerId for the instance
func (instance *__TimerEvent) SetPropertyTimerId(value string) (err error) {
	return instance.SetProperty("TimerId", value)
}

// GetTimerId gets the value of TimerId for the instance
func (instance *__TimerEvent) GetPropertyTimerId() (value string, err error) {
	retValue, err := instance.GetProperty("TimerId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
