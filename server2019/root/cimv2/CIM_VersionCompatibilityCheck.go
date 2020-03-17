// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_VersionCompatibilityCheck struct
type CIM_VersionCompatibilityCheck struct {
	CIM_Check

	//
	AllowDownVersion bool

	//
	AllowMultipleVersions bool

	//
	Reinstall bool
}

// SetAllowDownVersion sets the value of AllowDownVersion for the instance
func (instance *CIM_VersionCompatibilityCheck) SetPropertyAllowDownVersion(value bool) (err error) {
	return instance.SetProperty("AllowDownVersion", value)
}

// GetAllowDownVersion gets the value of AllowDownVersion for the instance
func (instance *CIM_VersionCompatibilityCheck) GetPropertyAllowDownVersion() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowDownVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowMultipleVersions sets the value of AllowMultipleVersions for the instance
func (instance *CIM_VersionCompatibilityCheck) SetPropertyAllowMultipleVersions(value bool) (err error) {
	return instance.SetProperty("AllowMultipleVersions", value)
}

// GetAllowMultipleVersions gets the value of AllowMultipleVersions for the instance
func (instance *CIM_VersionCompatibilityCheck) GetPropertyAllowMultipleVersions() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowMultipleVersions")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReinstall sets the value of Reinstall for the instance
func (instance *CIM_VersionCompatibilityCheck) SetPropertyReinstall(value bool) (err error) {
	return instance.SetProperty("Reinstall", value)
}

// GetReinstall gets the value of Reinstall for the instance
func (instance *CIM_VersionCompatibilityCheck) GetPropertyReinstall() (value bool, err error) {
	retValue, err := instance.GetProperty("Reinstall")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
