// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_LogonSession struct
type Win32_LogonSession struct {
	Win32_Session

	//
	AuthenticationPackage string

	//
	LogonId string

	//
	LogonType uint32
}

// SetAuthenticationPackage sets the value of AuthenticationPackage for the instance
func (instance *Win32_LogonSession) SetPropertyAuthenticationPackage(value string) (err error) {
	return instance.SetProperty("AuthenticationPackage", value)
}

// GetAuthenticationPackage gets the value of AuthenticationPackage for the instance
func (instance *Win32_LogonSession) GetPropertyAuthenticationPackage() (value string, err error) {
	retValue, err := instance.GetProperty("AuthenticationPackage")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogonId sets the value of LogonId for the instance
func (instance *Win32_LogonSession) SetPropertyLogonId(value string) (err error) {
	return instance.SetProperty("LogonId", value)
}

// GetLogonId gets the value of LogonId for the instance
func (instance *Win32_LogonSession) GetPropertyLogonId() (value string, err error) {
	retValue, err := instance.GetProperty("LogonId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogonType sets the value of LogonType for the instance
func (instance *Win32_LogonSession) SetPropertyLogonType(value uint32) (err error) {
	return instance.SetProperty("LogonType", value)
}

// GetLogonType gets the value of LogonType for the instance
func (instance *Win32_LogonSession) GetPropertyLogonType() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogonType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
