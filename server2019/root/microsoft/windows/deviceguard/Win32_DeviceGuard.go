// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DeviceGuard
//////////////////////////////////////////////
package deviceguard

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_DeviceGuard struct
type Win32_DeviceGuard struct {
	cim.WmiInstance

	//
	AvailableSecurityProperties []uint32

	//
	CodeIntegrityPolicyEnforcementStatus uint32

	//
	InstanceIdentifier string

	//
	RequiredSecurityProperties []uint32

	//
	SecurityServicesConfigured []uint32

	//
	SecurityServicesRunning []uint32

	//
	UsermodeCodeIntegrityPolicyEnforcementStatus uint32

	//
	Version string

	//
	VirtualizationBasedSecurityStatus uint32
}

// SetAvailableSecurityProperties sets the value of AvailableSecurityProperties for the instance
func (instance *Win32_DeviceGuard) SetPropertyAvailableSecurityProperties(value []uint32) (err error) {
	return instance.SetProperty("AvailableSecurityProperties", value)
}

// GetAvailableSecurityProperties gets the value of AvailableSecurityProperties for the instance
func (instance *Win32_DeviceGuard) GetPropertyAvailableSecurityProperties() (value []uint32, err error) {
	retValue, err := instance.GetProperty("AvailableSecurityProperties")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCodeIntegrityPolicyEnforcementStatus sets the value of CodeIntegrityPolicyEnforcementStatus for the instance
func (instance *Win32_DeviceGuard) SetPropertyCodeIntegrityPolicyEnforcementStatus(value uint32) (err error) {
	return instance.SetProperty("CodeIntegrityPolicyEnforcementStatus", value)
}

// GetCodeIntegrityPolicyEnforcementStatus gets the value of CodeIntegrityPolicyEnforcementStatus for the instance
func (instance *Win32_DeviceGuard) GetPropertyCodeIntegrityPolicyEnforcementStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("CodeIntegrityPolicyEnforcementStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceIdentifier sets the value of InstanceIdentifier for the instance
func (instance *Win32_DeviceGuard) SetPropertyInstanceIdentifier(value string) (err error) {
	return instance.SetProperty("InstanceIdentifier", value)
}

// GetInstanceIdentifier gets the value of InstanceIdentifier for the instance
func (instance *Win32_DeviceGuard) GetPropertyInstanceIdentifier() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceIdentifier")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequiredSecurityProperties sets the value of RequiredSecurityProperties for the instance
func (instance *Win32_DeviceGuard) SetPropertyRequiredSecurityProperties(value []uint32) (err error) {
	return instance.SetProperty("RequiredSecurityProperties", value)
}

// GetRequiredSecurityProperties gets the value of RequiredSecurityProperties for the instance
func (instance *Win32_DeviceGuard) GetPropertyRequiredSecurityProperties() (value []uint32, err error) {
	retValue, err := instance.GetProperty("RequiredSecurityProperties")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSecurityServicesConfigured sets the value of SecurityServicesConfigured for the instance
func (instance *Win32_DeviceGuard) SetPropertySecurityServicesConfigured(value []uint32) (err error) {
	return instance.SetProperty("SecurityServicesConfigured", value)
}

// GetSecurityServicesConfigured gets the value of SecurityServicesConfigured for the instance
func (instance *Win32_DeviceGuard) GetPropertySecurityServicesConfigured() (value []uint32, err error) {
	retValue, err := instance.GetProperty("SecurityServicesConfigured")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSecurityServicesRunning sets the value of SecurityServicesRunning for the instance
func (instance *Win32_DeviceGuard) SetPropertySecurityServicesRunning(value []uint32) (err error) {
	return instance.SetProperty("SecurityServicesRunning", value)
}

// GetSecurityServicesRunning gets the value of SecurityServicesRunning for the instance
func (instance *Win32_DeviceGuard) GetPropertySecurityServicesRunning() (value []uint32, err error) {
	retValue, err := instance.GetProperty("SecurityServicesRunning")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUsermodeCodeIntegrityPolicyEnforcementStatus sets the value of UsermodeCodeIntegrityPolicyEnforcementStatus for the instance
func (instance *Win32_DeviceGuard) SetPropertyUsermodeCodeIntegrityPolicyEnforcementStatus(value uint32) (err error) {
	return instance.SetProperty("UsermodeCodeIntegrityPolicyEnforcementStatus", value)
}

// GetUsermodeCodeIntegrityPolicyEnforcementStatus gets the value of UsermodeCodeIntegrityPolicyEnforcementStatus for the instance
func (instance *Win32_DeviceGuard) GetPropertyUsermodeCodeIntegrityPolicyEnforcementStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("UsermodeCodeIntegrityPolicyEnforcementStatus")
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
func (instance *Win32_DeviceGuard) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *Win32_DeviceGuard) GetPropertyVersion() (value string, err error) {
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

// SetVirtualizationBasedSecurityStatus sets the value of VirtualizationBasedSecurityStatus for the instance
func (instance *Win32_DeviceGuard) SetPropertyVirtualizationBasedSecurityStatus(value uint32) (err error) {
	return instance.SetProperty("VirtualizationBasedSecurityStatus", value)
}

// GetVirtualizationBasedSecurityStatus gets the value of VirtualizationBasedSecurityStatus for the instance
func (instance *Win32_DeviceGuard) GetPropertyVirtualizationBasedSecurityStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualizationBasedSecurityStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
