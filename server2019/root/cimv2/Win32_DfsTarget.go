// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_DfsTarget struct
type Win32_DfsTarget struct {
	CIM_LogicalElement

	//
	LinkName string

	//
	ServerName string

	//
	ShareName string

	//
	State uint32
}

// SetLinkName sets the value of LinkName for the instance
func (instance *Win32_DfsTarget) SetPropertyLinkName(value string) (err error) {
	return instance.SetProperty("LinkName", value)
}

// GetLinkName gets the value of LinkName for the instance
func (instance *Win32_DfsTarget) GetPropertyLinkName() (value string, err error) {
	retValue, err := instance.GetProperty("LinkName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerName sets the value of ServerName for the instance
func (instance *Win32_DfsTarget) SetPropertyServerName(value string) (err error) {
	return instance.SetProperty("ServerName", value)
}

// GetServerName gets the value of ServerName for the instance
func (instance *Win32_DfsTarget) GetPropertyServerName() (value string, err error) {
	retValue, err := instance.GetProperty("ServerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShareName sets the value of ShareName for the instance
func (instance *Win32_DfsTarget) SetPropertyShareName(value string) (err error) {
	return instance.SetProperty("ShareName", value)
}

// GetShareName gets the value of ShareName for the instance
func (instance *Win32_DfsTarget) GetPropertyShareName() (value string, err error) {
	retValue, err := instance.GetProperty("ShareName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *Win32_DfsTarget) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *Win32_DfsTarget) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
