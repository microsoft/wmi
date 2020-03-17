// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

// CIM_Log struct
type CIM_Log struct {
	CIM_EnabledLogicalElement

	//
	CurrentNumberOfRecords uint64

	//
	MaxNumberOfRecords uint64
}

// SetCurrentNumberOfRecords sets the value of CurrentNumberOfRecords for the instance
func (instance *CIM_Log) SetPropertyCurrentNumberOfRecords(value uint64) (err error) {
	return instance.SetProperty("CurrentNumberOfRecords", value)
}

// GetCurrentNumberOfRecords gets the value of CurrentNumberOfRecords for the instance
func (instance *CIM_Log) GetPropertyCurrentNumberOfRecords() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentNumberOfRecords")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxNumberOfRecords sets the value of MaxNumberOfRecords for the instance
func (instance *CIM_Log) SetPropertyMaxNumberOfRecords(value uint64) (err error) {
	return instance.SetProperty("MaxNumberOfRecords", value)
}

// GetMaxNumberOfRecords gets the value of MaxNumberOfRecords for the instance
func (instance *CIM_Log) GetPropertyMaxNumberOfRecords() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxNumberOfRecords")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Log) ClearLog() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ClearLog")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
