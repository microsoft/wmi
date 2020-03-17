// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_QuickFixEngineering struct
type Win32_QuickFixEngineering struct {
	CIM_LogicalElement

	//
	CSName string

	//
	FixComments string

	//
	HotFixID string

	//
	InstalledBy string

	//
	InstalledOn string

	//
	ServicePackInEffect string
}

// SetCSName sets the value of CSName for the instance
func (instance *Win32_QuickFixEngineering) SetPropertyCSName(value string) (err error) {
	return instance.SetProperty("CSName", value)
}

// GetCSName gets the value of CSName for the instance
func (instance *Win32_QuickFixEngineering) GetPropertyCSName() (value string, err error) {
	retValue, err := instance.GetProperty("CSName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFixComments sets the value of FixComments for the instance
func (instance *Win32_QuickFixEngineering) SetPropertyFixComments(value string) (err error) {
	return instance.SetProperty("FixComments", value)
}

// GetFixComments gets the value of FixComments for the instance
func (instance *Win32_QuickFixEngineering) GetPropertyFixComments() (value string, err error) {
	retValue, err := instance.GetProperty("FixComments")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHotFixID sets the value of HotFixID for the instance
func (instance *Win32_QuickFixEngineering) SetPropertyHotFixID(value string) (err error) {
	return instance.SetProperty("HotFixID", value)
}

// GetHotFixID gets the value of HotFixID for the instance
func (instance *Win32_QuickFixEngineering) GetPropertyHotFixID() (value string, err error) {
	retValue, err := instance.GetProperty("HotFixID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstalledBy sets the value of InstalledBy for the instance
func (instance *Win32_QuickFixEngineering) SetPropertyInstalledBy(value string) (err error) {
	return instance.SetProperty("InstalledBy", value)
}

// GetInstalledBy gets the value of InstalledBy for the instance
func (instance *Win32_QuickFixEngineering) GetPropertyInstalledBy() (value string, err error) {
	retValue, err := instance.GetProperty("InstalledBy")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstalledOn sets the value of InstalledOn for the instance
func (instance *Win32_QuickFixEngineering) SetPropertyInstalledOn(value string) (err error) {
	return instance.SetProperty("InstalledOn", value)
}

// GetInstalledOn gets the value of InstalledOn for the instance
func (instance *Win32_QuickFixEngineering) GetPropertyInstalledOn() (value string, err error) {
	retValue, err := instance.GetProperty("InstalledOn")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServicePackInEffect sets the value of ServicePackInEffect for the instance
func (instance *Win32_QuickFixEngineering) SetPropertyServicePackInEffect(value string) (err error) {
	return instance.SetProperty("ServicePackInEffect", value)
}

// GetServicePackInEffect gets the value of ServicePackInEffect for the instance
func (instance *Win32_QuickFixEngineering) GetPropertyServicePackInEffect() (value string, err error) {
	retValue, err := instance.GetProperty("ServicePackInEffect")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
