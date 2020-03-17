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

// MDM_Policy_Result01_DeviceInstallation02 struct
type MDM_Policy_Result01_DeviceInstallation02 struct {
	cim.WmiInstance

	//
	AllowInstallationOfMatchingDeviceIDs string

	//
	AllowInstallationOfMatchingDeviceInstanceIDs string

	//
	AllowInstallationOfMatchingDeviceSetupClasses string

	//
	InstanceID string

	//
	ParentID string

	//
	PreventDeviceMetadataFromNetwork string

	//
	PreventInstallationOfDevicesNotDescribedByOtherPolicySettings string

	//
	PreventInstallationOfMatchingDeviceIDs string

	//
	PreventInstallationOfMatchingDeviceInstanceIDs string

	//
	PreventInstallationOfMatchingDeviceSetupClasses string
}

// SetAllowInstallationOfMatchingDeviceIDs sets the value of AllowInstallationOfMatchingDeviceIDs for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) SetPropertyAllowInstallationOfMatchingDeviceIDs(value string) (err error) {
	return instance.SetProperty("AllowInstallationOfMatchingDeviceIDs", value)
}

// GetAllowInstallationOfMatchingDeviceIDs gets the value of AllowInstallationOfMatchingDeviceIDs for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) GetPropertyAllowInstallationOfMatchingDeviceIDs() (value string, err error) {
	retValue, err := instance.GetProperty("AllowInstallationOfMatchingDeviceIDs")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowInstallationOfMatchingDeviceInstanceIDs sets the value of AllowInstallationOfMatchingDeviceInstanceIDs for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) SetPropertyAllowInstallationOfMatchingDeviceInstanceIDs(value string) (err error) {
	return instance.SetProperty("AllowInstallationOfMatchingDeviceInstanceIDs", value)
}

// GetAllowInstallationOfMatchingDeviceInstanceIDs gets the value of AllowInstallationOfMatchingDeviceInstanceIDs for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) GetPropertyAllowInstallationOfMatchingDeviceInstanceIDs() (value string, err error) {
	retValue, err := instance.GetProperty("AllowInstallationOfMatchingDeviceInstanceIDs")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowInstallationOfMatchingDeviceSetupClasses sets the value of AllowInstallationOfMatchingDeviceSetupClasses for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) SetPropertyAllowInstallationOfMatchingDeviceSetupClasses(value string) (err error) {
	return instance.SetProperty("AllowInstallationOfMatchingDeviceSetupClasses", value)
}

// GetAllowInstallationOfMatchingDeviceSetupClasses gets the value of AllowInstallationOfMatchingDeviceSetupClasses for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) GetPropertyAllowInstallationOfMatchingDeviceSetupClasses() (value string, err error) {
	retValue, err := instance.GetProperty("AllowInstallationOfMatchingDeviceSetupClasses")
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
func (instance *MDM_Policy_Result01_DeviceInstallation02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_DeviceInstallation02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) GetPropertyParentID() (value string, err error) {
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

// SetPreventDeviceMetadataFromNetwork sets the value of PreventDeviceMetadataFromNetwork for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) SetPropertyPreventDeviceMetadataFromNetwork(value string) (err error) {
	return instance.SetProperty("PreventDeviceMetadataFromNetwork", value)
}

// GetPreventDeviceMetadataFromNetwork gets the value of PreventDeviceMetadataFromNetwork for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) GetPropertyPreventDeviceMetadataFromNetwork() (value string, err error) {
	retValue, err := instance.GetProperty("PreventDeviceMetadataFromNetwork")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPreventInstallationOfDevicesNotDescribedByOtherPolicySettings sets the value of PreventInstallationOfDevicesNotDescribedByOtherPolicySettings for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) SetPropertyPreventInstallationOfDevicesNotDescribedByOtherPolicySettings(value string) (err error) {
	return instance.SetProperty("PreventInstallationOfDevicesNotDescribedByOtherPolicySettings", value)
}

// GetPreventInstallationOfDevicesNotDescribedByOtherPolicySettings gets the value of PreventInstallationOfDevicesNotDescribedByOtherPolicySettings for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) GetPropertyPreventInstallationOfDevicesNotDescribedByOtherPolicySettings() (value string, err error) {
	retValue, err := instance.GetProperty("PreventInstallationOfDevicesNotDescribedByOtherPolicySettings")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPreventInstallationOfMatchingDeviceIDs sets the value of PreventInstallationOfMatchingDeviceIDs for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) SetPropertyPreventInstallationOfMatchingDeviceIDs(value string) (err error) {
	return instance.SetProperty("PreventInstallationOfMatchingDeviceIDs", value)
}

// GetPreventInstallationOfMatchingDeviceIDs gets the value of PreventInstallationOfMatchingDeviceIDs for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) GetPropertyPreventInstallationOfMatchingDeviceIDs() (value string, err error) {
	retValue, err := instance.GetProperty("PreventInstallationOfMatchingDeviceIDs")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPreventInstallationOfMatchingDeviceInstanceIDs sets the value of PreventInstallationOfMatchingDeviceInstanceIDs for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) SetPropertyPreventInstallationOfMatchingDeviceInstanceIDs(value string) (err error) {
	return instance.SetProperty("PreventInstallationOfMatchingDeviceInstanceIDs", value)
}

// GetPreventInstallationOfMatchingDeviceInstanceIDs gets the value of PreventInstallationOfMatchingDeviceInstanceIDs for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) GetPropertyPreventInstallationOfMatchingDeviceInstanceIDs() (value string, err error) {
	retValue, err := instance.GetProperty("PreventInstallationOfMatchingDeviceInstanceIDs")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPreventInstallationOfMatchingDeviceSetupClasses sets the value of PreventInstallationOfMatchingDeviceSetupClasses for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) SetPropertyPreventInstallationOfMatchingDeviceSetupClasses(value string) (err error) {
	return instance.SetProperty("PreventInstallationOfMatchingDeviceSetupClasses", value)
}

// GetPreventInstallationOfMatchingDeviceSetupClasses gets the value of PreventInstallationOfMatchingDeviceSetupClasses for the instance
func (instance *MDM_Policy_Result01_DeviceInstallation02) GetPropertyPreventInstallationOfMatchingDeviceSetupClasses() (value string, err error) {
	retValue, err := instance.GetProperty("PreventInstallationOfMatchingDeviceSetupClasses")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
