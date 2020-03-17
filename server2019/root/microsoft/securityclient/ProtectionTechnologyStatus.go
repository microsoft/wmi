// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.SecurityClient
//////////////////////////////////////////////
package securityclient

// ProtectionTechnologyStatus struct
type ProtectionTechnologyStatus struct {
	SerializableToXml

	// Is protection technology enabled
	Enabled bool

	// Protection technology name
	Name string

	// Protection technology version (major, minor, build, revision)
	Version string
}

// SetEnabled sets the value of Enabled for the instance
func (instance *ProtectionTechnologyStatus) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *ProtectionTechnologyStatus) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *ProtectionTechnologyStatus) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *ProtectionTechnologyStatus) GetPropertyName() (value string, err error) {
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

// SetVersion sets the value of Version for the instance
func (instance *ProtectionTechnologyStatus) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *ProtectionTechnologyStatus) GetPropertyVersion() (value string, err error) {
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
