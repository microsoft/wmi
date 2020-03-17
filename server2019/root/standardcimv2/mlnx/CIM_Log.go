// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_Log struct
type CIM_Log struct {
	CIM_EnabledLogicalElement

	//
	CurrentNumberOfRecords uint64

	//
	LogState Log_LogState

	//
	MaxNumberOfRecords uint64

	//
	OverwritePolicy Log_OverwritePolicy
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

// SetLogState sets the value of LogState for the instance
func (instance *CIM_Log) SetPropertyLogState(value Log_LogState) (err error) {
	return instance.SetProperty("LogState", value)
}

// GetLogState gets the value of LogState for the instance
func (instance *CIM_Log) GetPropertyLogState() (value Log_LogState, err error) {
	retValue, err := instance.GetProperty("LogState")
	if err != nil {
		return
	}
	value, ok := retValue.(Log_LogState)
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

// SetOverwritePolicy sets the value of OverwritePolicy for the instance
func (instance *CIM_Log) SetPropertyOverwritePolicy(value Log_OverwritePolicy) (err error) {
	return instance.SetProperty("OverwritePolicy", value)
}

// GetOverwritePolicy gets the value of OverwritePolicy for the instance
func (instance *CIM_Log) GetPropertyOverwritePolicy() (value Log_OverwritePolicy, err error) {
	retValue, err := instance.GetProperty("OverwritePolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(Log_OverwritePolicy)
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
