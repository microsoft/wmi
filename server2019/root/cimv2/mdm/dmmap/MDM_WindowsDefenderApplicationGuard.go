// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_WindowsDefenderApplicationGuard struct
type MDM_WindowsDefenderApplicationGuard struct {
	cim.WmiInstance

	//
	InstallWindowsDefenderApplicationGuard string

	//
	InstanceID string

	//
	ParentID string

	//
	PlatformStatus int32

	//
	Status int32
}

// SetInstallWindowsDefenderApplicationGuard sets the value of InstallWindowsDefenderApplicationGuard for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) SetPropertyInstallWindowsDefenderApplicationGuard(value string) (err error) {
	return instance.SetProperty("InstallWindowsDefenderApplicationGuard", value)
}

// GetInstallWindowsDefenderApplicationGuard gets the value of InstallWindowsDefenderApplicationGuard for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) GetPropertyInstallWindowsDefenderApplicationGuard() (value string, err error) {
	retValue, err := instance.GetProperty("InstallWindowsDefenderApplicationGuard")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPlatformStatus sets the value of PlatformStatus for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) SetPropertyPlatformStatus(value int32) (err error) {
	return instance.SetProperty("PlatformStatus", value)
}

// GetPlatformStatus gets the value of PlatformStatus for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) GetPropertyPlatformStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("PlatformStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) SetPropertyStatus(value int32) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) GetPropertyStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="param" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_WindowsDefenderApplicationGuard) InstallWindowsDefenderApplicationGuardMethod( /* IN */ param string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("InstallWindowsDefenderApplicationGuardMethod", param)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
