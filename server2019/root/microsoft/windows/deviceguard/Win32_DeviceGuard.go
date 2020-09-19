// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.DeviceGuard
//////////////////////////////////////////////
package deviceguard

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_DeviceGuard struct
type Win32_DeviceGuard struct {
	*cim.WmiInstance

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

func NewWin32_DeviceGuardEx1(instance *cim.WmiInstance) (newInstance *Win32_DeviceGuard, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Win32_DeviceGuard{
		WmiInstance: tmp,
	}
	return
}

func NewWin32_DeviceGuardEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_DeviceGuard, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_DeviceGuard{
		WmiInstance: tmp,
	}
	return
}

// SetAvailableSecurityProperties sets the value of AvailableSecurityProperties for the instance
func (instance *Win32_DeviceGuard) SetPropertyAvailableSecurityProperties(value []uint32) (err error) {
	return instance.SetProperty("AvailableSecurityProperties", (value))
}

// GetAvailableSecurityProperties gets the value of AvailableSecurityProperties for the instance
func (instance *Win32_DeviceGuard) GetPropertyAvailableSecurityProperties() (value []uint32, err error) {
	retValue, err := instance.GetProperty("AvailableSecurityProperties")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetCodeIntegrityPolicyEnforcementStatus sets the value of CodeIntegrityPolicyEnforcementStatus for the instance
func (instance *Win32_DeviceGuard) SetPropertyCodeIntegrityPolicyEnforcementStatus(value uint32) (err error) {
	return instance.SetProperty("CodeIntegrityPolicyEnforcementStatus", (value))
}

// GetCodeIntegrityPolicyEnforcementStatus gets the value of CodeIntegrityPolicyEnforcementStatus for the instance
func (instance *Win32_DeviceGuard) GetPropertyCodeIntegrityPolicyEnforcementStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("CodeIntegrityPolicyEnforcementStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetInstanceIdentifier sets the value of InstanceIdentifier for the instance
func (instance *Win32_DeviceGuard) SetPropertyInstanceIdentifier(value string) (err error) {
	return instance.SetProperty("InstanceIdentifier", (value))
}

// GetInstanceIdentifier gets the value of InstanceIdentifier for the instance
func (instance *Win32_DeviceGuard) GetPropertyInstanceIdentifier() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceIdentifier")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetRequiredSecurityProperties sets the value of RequiredSecurityProperties for the instance
func (instance *Win32_DeviceGuard) SetPropertyRequiredSecurityProperties(value []uint32) (err error) {
	return instance.SetProperty("RequiredSecurityProperties", (value))
}

// GetRequiredSecurityProperties gets the value of RequiredSecurityProperties for the instance
func (instance *Win32_DeviceGuard) GetPropertyRequiredSecurityProperties() (value []uint32, err error) {
	retValue, err := instance.GetProperty("RequiredSecurityProperties")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetSecurityServicesConfigured sets the value of SecurityServicesConfigured for the instance
func (instance *Win32_DeviceGuard) SetPropertySecurityServicesConfigured(value []uint32) (err error) {
	return instance.SetProperty("SecurityServicesConfigured", (value))
}

// GetSecurityServicesConfigured gets the value of SecurityServicesConfigured for the instance
func (instance *Win32_DeviceGuard) GetPropertySecurityServicesConfigured() (value []uint32, err error) {
	retValue, err := instance.GetProperty("SecurityServicesConfigured")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetSecurityServicesRunning sets the value of SecurityServicesRunning for the instance
func (instance *Win32_DeviceGuard) SetPropertySecurityServicesRunning(value []uint32) (err error) {
	return instance.SetProperty("SecurityServicesRunning", (value))
}

// GetSecurityServicesRunning gets the value of SecurityServicesRunning for the instance
func (instance *Win32_DeviceGuard) GetPropertySecurityServicesRunning() (value []uint32, err error) {
	retValue, err := instance.GetProperty("SecurityServicesRunning")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetUsermodeCodeIntegrityPolicyEnforcementStatus sets the value of UsermodeCodeIntegrityPolicyEnforcementStatus for the instance
func (instance *Win32_DeviceGuard) SetPropertyUsermodeCodeIntegrityPolicyEnforcementStatus(value uint32) (err error) {
	return instance.SetProperty("UsermodeCodeIntegrityPolicyEnforcementStatus", (value))
}

// GetUsermodeCodeIntegrityPolicyEnforcementStatus gets the value of UsermodeCodeIntegrityPolicyEnforcementStatus for the instance
func (instance *Win32_DeviceGuard) GetPropertyUsermodeCodeIntegrityPolicyEnforcementStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("UsermodeCodeIntegrityPolicyEnforcementStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetVersion sets the value of Version for the instance
func (instance *Win32_DeviceGuard) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *Win32_DeviceGuard) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetVirtualizationBasedSecurityStatus sets the value of VirtualizationBasedSecurityStatus for the instance
func (instance *Win32_DeviceGuard) SetPropertyVirtualizationBasedSecurityStatus(value uint32) (err error) {
	return instance.SetProperty("VirtualizationBasedSecurityStatus", (value))
}

// GetVirtualizationBasedSecurityStatus gets the value of VirtualizationBasedSecurityStatus for the instance
func (instance *Win32_DeviceGuard) GetPropertyVirtualizationBasedSecurityStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualizationBasedSecurityStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
