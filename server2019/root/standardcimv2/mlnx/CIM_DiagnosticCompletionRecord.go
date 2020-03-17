// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_DiagnosticCompletionRecord struct
type CIM_DiagnosticCompletionRecord struct {
	CIM_DiagnosticServiceRecord

	//
	CompletionState DiagnosticCompletionRecord_CompletionState

	//
	OtherCompletionStateDescription string
}

// SetCompletionState sets the value of CompletionState for the instance
func (instance *CIM_DiagnosticCompletionRecord) SetPropertyCompletionState(value DiagnosticCompletionRecord_CompletionState) (err error) {
	return instance.SetProperty("CompletionState", value)
}

// GetCompletionState gets the value of CompletionState for the instance
func (instance *CIM_DiagnosticCompletionRecord) GetPropertyCompletionState() (value DiagnosticCompletionRecord_CompletionState, err error) {
	retValue, err := instance.GetProperty("CompletionState")
	if err != nil {
		return
	}
	value, ok := retValue.(DiagnosticCompletionRecord_CompletionState)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherCompletionStateDescription sets the value of OtherCompletionStateDescription for the instance
func (instance *CIM_DiagnosticCompletionRecord) SetPropertyOtherCompletionStateDescription(value string) (err error) {
	return instance.SetProperty("OtherCompletionStateDescription", value)
}

// GetOtherCompletionStateDescription gets the value of OtherCompletionStateDescription for the instance
func (instance *CIM_DiagnosticCompletionRecord) GetPropertyOtherCompletionStateDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherCompletionStateDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
