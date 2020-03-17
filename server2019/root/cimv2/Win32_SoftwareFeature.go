// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_SoftwareFeature struct
type Win32_SoftwareFeature struct {
	CIM_SoftwareFeature

	//
	Accesses uint16

	//
	Attributes uint16

	//
	InstallState int16

	//
	LastUse string
}

// SetAccesses sets the value of Accesses for the instance
func (instance *Win32_SoftwareFeature) SetPropertyAccesses(value uint16) (err error) {
	return instance.SetProperty("Accesses", value)
}

// GetAccesses gets the value of Accesses for the instance
func (instance *Win32_SoftwareFeature) GetPropertyAccesses() (value uint16, err error) {
	retValue, err := instance.GetProperty("Accesses")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAttributes sets the value of Attributes for the instance
func (instance *Win32_SoftwareFeature) SetPropertyAttributes(value uint16) (err error) {
	return instance.SetProperty("Attributes", value)
}

// GetAttributes gets the value of Attributes for the instance
func (instance *Win32_SoftwareFeature) GetPropertyAttributes() (value uint16, err error) {
	retValue, err := instance.GetProperty("Attributes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstallState sets the value of InstallState for the instance
func (instance *Win32_SoftwareFeature) SetPropertyInstallState(value int16) (err error) {
	return instance.SetProperty("InstallState", value)
}

// GetInstallState gets the value of InstallState for the instance
func (instance *Win32_SoftwareFeature) GetPropertyInstallState() (value int16, err error) {
	retValue, err := instance.GetProperty("InstallState")
	if err != nil {
		return
	}
	value, ok := retValue.(int16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastUse sets the value of LastUse for the instance
func (instance *Win32_SoftwareFeature) SetPropertyLastUse(value string) (err error) {
	return instance.SetProperty("LastUse", value)
}

// GetLastUse gets the value of LastUse for the instance
func (instance *Win32_SoftwareFeature) GetPropertyLastUse() (value string, err error) {
	retValue, err := instance.GetProperty("LastUse")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ReinstallMode" type="uint16 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_SoftwareFeature) Reinstall( /* IN */ ReinstallMode uint16) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Reinstall", ReinstallMode)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="InstallState" type="uint16 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_SoftwareFeature) Configure( /* IN */ InstallState uint16) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Configure", InstallState)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
