// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

// MSFT_TaskComHandlerAction struct
type MSFT_TaskComHandlerAction struct {
	MSFT_TaskAction

	//
	ClassId string

	//
	Data string
}

// SetClassId sets the value of ClassId for the instance
func (instance *MSFT_TaskComHandlerAction) SetPropertyClassId(value string) (err error) {
	return instance.SetProperty("ClassId", value)
}

// GetClassId gets the value of ClassId for the instance
func (instance *MSFT_TaskComHandlerAction) GetPropertyClassId() (value string, err error) {
	retValue, err := instance.GetProperty("ClassId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetData sets the value of Data for the instance
func (instance *MSFT_TaskComHandlerAction) SetPropertyData(value string) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *MSFT_TaskComHandlerAction) GetPropertyData() (value string, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
