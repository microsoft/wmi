// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

// MSFT_ServerError struct
type MSFT_ServerError struct {
	CIM_Error

	//
	ErrorCode uint32

	//
	ExtendedErrorCode uint32
}

// SetErrorCode sets the value of ErrorCode for the instance
func (instance *MSFT_ServerError) SetPropertyErrorCode(value uint32) (err error) {
	return instance.SetProperty("ErrorCode", value)
}

// GetErrorCode gets the value of ErrorCode for the instance
func (instance *MSFT_ServerError) GetPropertyErrorCode() (value uint32, err error) {
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

// SetExtendedErrorCode sets the value of ExtendedErrorCode for the instance
func (instance *MSFT_ServerError) SetPropertyExtendedErrorCode(value uint32) (err error) {
	return instance.SetProperty("ExtendedErrorCode", value)
}

// GetExtendedErrorCode gets the value of ExtendedErrorCode for the instance
func (instance *MSFT_ServerError) GetPropertyExtendedErrorCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExtendedErrorCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
