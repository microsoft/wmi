// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

// Win32_TSStartMenuApplication struct
type Win32_TSStartMenuApplication struct {
	CIM_LogicalElement

	// Command line arguments
	CommandLineArguments string

	// Index of the icon
	IconIndex int32

	// Path to the application icon
	IconPath string

	// Path to the application
	Path string

	// Virtual Path to the application (includes Environment Variables)
	VPath string
}

// SetCommandLineArguments sets the value of CommandLineArguments for the instance
func (instance *Win32_TSStartMenuApplication) SetPropertyCommandLineArguments(value string) (err error) {
	return instance.SetProperty("CommandLineArguments", value)
}

// GetCommandLineArguments gets the value of CommandLineArguments for the instance
func (instance *Win32_TSStartMenuApplication) GetPropertyCommandLineArguments() (value string, err error) {
	retValue, err := instance.GetProperty("CommandLineArguments")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIconIndex sets the value of IconIndex for the instance
func (instance *Win32_TSStartMenuApplication) SetPropertyIconIndex(value int32) (err error) {
	return instance.SetProperty("IconIndex", value)
}

// GetIconIndex gets the value of IconIndex for the instance
func (instance *Win32_TSStartMenuApplication) GetPropertyIconIndex() (value int32, err error) {
	retValue, err := instance.GetProperty("IconIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIconPath sets the value of IconPath for the instance
func (instance *Win32_TSStartMenuApplication) SetPropertyIconPath(value string) (err error) {
	return instance.SetProperty("IconPath", value)
}

// GetIconPath gets the value of IconPath for the instance
func (instance *Win32_TSStartMenuApplication) GetPropertyIconPath() (value string, err error) {
	retValue, err := instance.GetProperty("IconPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPath sets the value of Path for the instance
func (instance *Win32_TSStartMenuApplication) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", value)
}

// GetPath gets the value of Path for the instance
func (instance *Win32_TSStartMenuApplication) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVPath sets the value of VPath for the instance
func (instance *Win32_TSStartMenuApplication) SetPropertyVPath(value string) (err error) {
	return instance.SetProperty("VPath", value)
}

// GetVPath gets the value of VPath for the instance
func (instance *Win32_TSStartMenuApplication) GetPropertyVPath() (value string, err error) {
	retValue, err := instance.GetProperty("VPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
