// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_DiagnosticServiceRecord struct
type CIM_DiagnosticServiceRecord struct {
	CIM_DiagnosticRecord

	//
	ErrorCode []string

	//
	ErrorCount []uint32

	//
	LoopsFailed uint32

	//
	LoopsPassed uint32
}

// SetErrorCode sets the value of ErrorCode for the instance
func (instance *CIM_DiagnosticServiceRecord) SetPropertyErrorCode(value []string) (err error) {
	return instance.SetProperty("ErrorCode", value)
}

// GetErrorCode gets the value of ErrorCode for the instance
func (instance *CIM_DiagnosticServiceRecord) GetPropertyErrorCode() (value []string, err error) {
	retValue, err := instance.GetProperty("ErrorCode")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorCount sets the value of ErrorCount for the instance
func (instance *CIM_DiagnosticServiceRecord) SetPropertyErrorCount(value []uint32) (err error) {
	return instance.SetProperty("ErrorCount", value)
}

// GetErrorCount gets the value of ErrorCount for the instance
func (instance *CIM_DiagnosticServiceRecord) GetPropertyErrorCount() (value []uint32, err error) {
	retValue, err := instance.GetProperty("ErrorCount")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLoopsFailed sets the value of LoopsFailed for the instance
func (instance *CIM_DiagnosticServiceRecord) SetPropertyLoopsFailed(value uint32) (err error) {
	return instance.SetProperty("LoopsFailed", value)
}

// GetLoopsFailed gets the value of LoopsFailed for the instance
func (instance *CIM_DiagnosticServiceRecord) GetPropertyLoopsFailed() (value uint32, err error) {
	retValue, err := instance.GetProperty("LoopsFailed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLoopsPassed sets the value of LoopsPassed for the instance
func (instance *CIM_DiagnosticServiceRecord) SetPropertyLoopsPassed(value uint32) (err error) {
	return instance.SetProperty("LoopsPassed", value)
}

// GetLoopsPassed gets the value of LoopsPassed for the instance
func (instance *CIM_DiagnosticServiceRecord) GetPropertyLoopsPassed() (value uint32, err error) {
	retValue, err := instance.GetProperty("LoopsPassed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
