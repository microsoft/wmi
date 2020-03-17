// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// MLNX_SoftwareIdentity struct
type MLNX_SoftwareIdentity struct {
	CIM_SoftwareIdentity

	//
	InstallLocation string
}

// SetInstallLocation sets the value of InstallLocation for the instance
func (instance *MLNX_SoftwareIdentity) SetPropertyInstallLocation(value string) (err error) {
	return instance.SetProperty("InstallLocation", value)
}

// GetInstallLocation gets the value of InstallLocation for the instance
func (instance *MLNX_SoftwareIdentity) GetPropertyInstallLocation() (value string, err error) {
	retValue, err := instance.GetProperty("InstallLocation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
