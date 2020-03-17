// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ShadowProvider struct
type Win32_ShadowProvider struct {
	CIM_LogicalElement

	//
	CLSID string

	//
	ID string

	//
	Type uint32

	//
	Version string

	//
	VersionID string
}

// SetCLSID sets the value of CLSID for the instance
func (instance *Win32_ShadowProvider) SetPropertyCLSID(value string) (err error) {
	return instance.SetProperty("CLSID", value)
}

// GetCLSID gets the value of CLSID for the instance
func (instance *Win32_ShadowProvider) GetPropertyCLSID() (value string, err error) {
	retValue, err := instance.GetProperty("CLSID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetID sets the value of ID for the instance
func (instance *Win32_ShadowProvider) SetPropertyID(value string) (err error) {
	return instance.SetProperty("ID", value)
}

// GetID gets the value of ID for the instance
func (instance *Win32_ShadowProvider) GetPropertyID() (value string, err error) {
	retValue, err := instance.GetProperty("ID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *Win32_ShadowProvider) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *Win32_ShadowProvider) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersion sets the value of Version for the instance
func (instance *Win32_ShadowProvider) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *Win32_ShadowProvider) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersionID sets the value of VersionID for the instance
func (instance *Win32_ShadowProvider) SetPropertyVersionID(value string) (err error) {
	return instance.SetProperty("VersionID", value)
}

// GetVersionID gets the value of VersionID for the instance
func (instance *Win32_ShadowProvider) GetPropertyVersionID() (value string, err error) {
	retValue, err := instance.GetProperty("VersionID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
