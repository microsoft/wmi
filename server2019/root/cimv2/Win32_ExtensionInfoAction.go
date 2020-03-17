// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ExtensionInfoAction struct
type Win32_ExtensionInfoAction struct {
	CIM_Action

	//
	Argument string

	//
	Command string

	//
	Extension string

	//
	MIME string

	//
	ProgID string

	//
	ShellNew string

	//
	ShellNewValue string

	//
	Verb string
}

// SetArgument sets the value of Argument for the instance
func (instance *Win32_ExtensionInfoAction) SetPropertyArgument(value string) (err error) {
	return instance.SetProperty("Argument", value)
}

// GetArgument gets the value of Argument for the instance
func (instance *Win32_ExtensionInfoAction) GetPropertyArgument() (value string, err error) {
	retValue, err := instance.GetProperty("Argument")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCommand sets the value of Command for the instance
func (instance *Win32_ExtensionInfoAction) SetPropertyCommand(value string) (err error) {
	return instance.SetProperty("Command", value)
}

// GetCommand gets the value of Command for the instance
func (instance *Win32_ExtensionInfoAction) GetPropertyCommand() (value string, err error) {
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

// SetExtension sets the value of Extension for the instance
func (instance *Win32_ExtensionInfoAction) SetPropertyExtension(value string) (err error) {
	return instance.SetProperty("Extension", value)
}

// GetExtension gets the value of Extension for the instance
func (instance *Win32_ExtensionInfoAction) GetPropertyExtension() (value string, err error) {
	retValue, err := instance.GetProperty("Extension")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMIME sets the value of MIME for the instance
func (instance *Win32_ExtensionInfoAction) SetPropertyMIME(value string) (err error) {
	return instance.SetProperty("MIME", value)
}

// GetMIME gets the value of MIME for the instance
func (instance *Win32_ExtensionInfoAction) GetPropertyMIME() (value string, err error) {
	retValue, err := instance.GetProperty("MIME")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProgID sets the value of ProgID for the instance
func (instance *Win32_ExtensionInfoAction) SetPropertyProgID(value string) (err error) {
	return instance.SetProperty("ProgID", value)
}

// GetProgID gets the value of ProgID for the instance
func (instance *Win32_ExtensionInfoAction) GetPropertyProgID() (value string, err error) {
	retValue, err := instance.GetProperty("ProgID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShellNew sets the value of ShellNew for the instance
func (instance *Win32_ExtensionInfoAction) SetPropertyShellNew(value string) (err error) {
	return instance.SetProperty("ShellNew", value)
}

// GetShellNew gets the value of ShellNew for the instance
func (instance *Win32_ExtensionInfoAction) GetPropertyShellNew() (value string, err error) {
	retValue, err := instance.GetProperty("ShellNew")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShellNewValue sets the value of ShellNewValue for the instance
func (instance *Win32_ExtensionInfoAction) SetPropertyShellNewValue(value string) (err error) {
	return instance.SetProperty("ShellNewValue", value)
}

// GetShellNewValue gets the value of ShellNewValue for the instance
func (instance *Win32_ExtensionInfoAction) GetPropertyShellNewValue() (value string, err error) {
	retValue, err := instance.GetProperty("ShellNewValue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVerb sets the value of Verb for the instance
func (instance *Win32_ExtensionInfoAction) SetPropertyVerb(value string) (err error) {
	return instance.SetProperty("Verb", value)
}

// GetVerb gets the value of Verb for the instance
func (instance *Win32_ExtensionInfoAction) GetPropertyVerb() (value string, err error) {
	retValue, err := instance.GetProperty("Verb")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
