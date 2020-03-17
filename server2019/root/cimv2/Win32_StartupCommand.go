// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_StartupCommand struct
type Win32_StartupCommand struct {
	CIM_Setting

	//
	Command string

	//
	Location string

	//
	Name string

	//
	User string

	//
	UserSID string
}

// SetCommand sets the value of Command for the instance
func (instance *Win32_StartupCommand) SetPropertyCommand(value string) (err error) {
	return instance.SetProperty("Command", value)
}

// GetCommand gets the value of Command for the instance
func (instance *Win32_StartupCommand) GetPropertyCommand() (value string, err error) {
	retValue, err := instance.GetProperty("Command")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocation sets the value of Location for the instance
func (instance *Win32_StartupCommand) SetPropertyLocation(value string) (err error) {
	return instance.SetProperty("Location", value)
}

// GetLocation gets the value of Location for the instance
func (instance *Win32_StartupCommand) GetPropertyLocation() (value string, err error) {
	retValue, err := instance.GetProperty("Location")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *Win32_StartupCommand) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *Win32_StartupCommand) GetPropertyName() (value string, err error) {
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

// SetUser sets the value of User for the instance
func (instance *Win32_StartupCommand) SetPropertyUser(value string) (err error) {
	return instance.SetProperty("User", value)
}

// GetUser gets the value of User for the instance
func (instance *Win32_StartupCommand) GetPropertyUser() (value string, err error) {
	retValue, err := instance.GetProperty("User")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserSID sets the value of UserSID for the instance
func (instance *Win32_StartupCommand) SetPropertyUserSID(value string) (err error) {
	return instance.SetProperty("UserSID", value)
}

// GetUserSID gets the value of UserSID for the instance
func (instance *Win32_StartupCommand) GetPropertyUserSID() (value string, err error) {
	retValue, err := instance.GetProperty("UserSID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
