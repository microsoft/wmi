// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.Security
//////////////////////////////////////////////
package security

// __QOSFailureEvent struct
type __QOSFailureEvent struct {
	__EventDroppedEvent

	//
	ErrorCode uint32

	//
	ErrorDescription string
}

// SetErrorCode sets the value of ErrorCode for the instance
func (instance *__QOSFailureEvent) SetPropertyErrorCode(value uint32) (err error) {
	return instance.SetProperty("ErrorCode", value)
}

// GetErrorCode gets the value of ErrorCode for the instance
func (instance *__QOSFailureEvent) GetPropertyErrorCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("ErrorCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorDescription sets the value of ErrorDescription for the instance
func (instance *__QOSFailureEvent) SetPropertyErrorDescription(value string) (err error) {
	return instance.SetProperty("ErrorDescription", value)
}

// GetErrorDescription gets the value of ErrorDescription for the instance
func (instance *__QOSFailureEvent) GetPropertyErrorDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
