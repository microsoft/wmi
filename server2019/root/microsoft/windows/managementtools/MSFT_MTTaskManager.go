// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

// MSFT_MTTaskManager struct
type MSFT_MTTaskManager struct {
	CIM_ManagedElement

	//
	CurrentIndex uint16

	//
	IntervalSeconds uint16

	//
	Name string
}

// SetCurrentIndex sets the value of CurrentIndex for the instance
func (instance *MSFT_MTTaskManager) SetPropertyCurrentIndex(value uint16) (err error) {
	return instance.SetProperty("CurrentIndex", value)
}

// GetCurrentIndex gets the value of CurrentIndex for the instance
func (instance *MSFT_MTTaskManager) GetPropertyCurrentIndex() (value uint16, err error) {
	retValue, err := instance.GetProperty("CurrentIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntervalSeconds sets the value of IntervalSeconds for the instance
func (instance *MSFT_MTTaskManager) SetPropertyIntervalSeconds(value uint16) (err error) {
	return instance.SetProperty("IntervalSeconds", value)
}

// GetIntervalSeconds gets the value of IntervalSeconds for the instance
func (instance *MSFT_MTTaskManager) GetPropertyIntervalSeconds() (value uint16, err error) {
	retValue, err := instance.GetProperty("IntervalSeconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_MTTaskManager) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_MTTaskManager) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Seconds" type="uint16 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTTaskManager) SetInterval( /* IN */ Seconds uint16) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetInterval", Seconds)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTTaskManager) ForceRefresh() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ForceRefresh")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
