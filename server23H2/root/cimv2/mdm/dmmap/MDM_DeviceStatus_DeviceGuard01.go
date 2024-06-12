// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_DeviceStatus_DeviceGuard01 struct
type MDM_DeviceStatus_DeviceGuard01 struct {
	*cim.WmiInstance

	//
	HypervisorEnforcedCodeIntegrityStatus int32

	//
	InstanceID string

	//
	LsaCfgCredGuardStatus int32

	//
	ParentID string

	//
	SystemGuardStatus int32

	//
	VirtualizationBasedSecurityHwReq int32

	//
	VirtualizationBasedSecurityStatus int32
}

func NewMDM_DeviceStatus_DeviceGuard01Ex1(instance *cim.WmiInstance) (newInstance *MDM_DeviceStatus_DeviceGuard01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_DeviceStatus_DeviceGuard01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_DeviceStatus_DeviceGuard01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_DeviceStatus_DeviceGuard01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_DeviceStatus_DeviceGuard01{
		WmiInstance: tmp,
	}
	return
}

// SetHypervisorEnforcedCodeIntegrityStatus sets the value of HypervisorEnforcedCodeIntegrityStatus for the instance
func (instance *MDM_DeviceStatus_DeviceGuard01) SetPropertyHypervisorEnforcedCodeIntegrityStatus(value int32) (err error) {
	return instance.SetProperty("HypervisorEnforcedCodeIntegrityStatus", (value))
}

// GetHypervisorEnforcedCodeIntegrityStatus gets the value of HypervisorEnforcedCodeIntegrityStatus for the instance
func (instance *MDM_DeviceStatus_DeviceGuard01) GetPropertyHypervisorEnforcedCodeIntegrityStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("HypervisorEnforcedCodeIntegrityStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_DeviceStatus_DeviceGuard01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DeviceStatus_DeviceGuard01) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetLsaCfgCredGuardStatus sets the value of LsaCfgCredGuardStatus for the instance
func (instance *MDM_DeviceStatus_DeviceGuard01) SetPropertyLsaCfgCredGuardStatus(value int32) (err error) {
	return instance.SetProperty("LsaCfgCredGuardStatus", (value))
}

// GetLsaCfgCredGuardStatus gets the value of LsaCfgCredGuardStatus for the instance
func (instance *MDM_DeviceStatus_DeviceGuard01) GetPropertyLsaCfgCredGuardStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("LsaCfgCredGuardStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_DeviceStatus_DeviceGuard01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DeviceStatus_DeviceGuard01) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
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

// SetSystemGuardStatus sets the value of SystemGuardStatus for the instance
func (instance *MDM_DeviceStatus_DeviceGuard01) SetPropertySystemGuardStatus(value int32) (err error) {
	return instance.SetProperty("SystemGuardStatus", (value))
}

// GetSystemGuardStatus gets the value of SystemGuardStatus for the instance
func (instance *MDM_DeviceStatus_DeviceGuard01) GetPropertySystemGuardStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("SystemGuardStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetVirtualizationBasedSecurityHwReq sets the value of VirtualizationBasedSecurityHwReq for the instance
func (instance *MDM_DeviceStatus_DeviceGuard01) SetPropertyVirtualizationBasedSecurityHwReq(value int32) (err error) {
	return instance.SetProperty("VirtualizationBasedSecurityHwReq", (value))
}

// GetVirtualizationBasedSecurityHwReq gets the value of VirtualizationBasedSecurityHwReq for the instance
func (instance *MDM_DeviceStatus_DeviceGuard01) GetPropertyVirtualizationBasedSecurityHwReq() (value int32, err error) {
	retValue, err := instance.GetProperty("VirtualizationBasedSecurityHwReq")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetVirtualizationBasedSecurityStatus sets the value of VirtualizationBasedSecurityStatus for the instance
func (instance *MDM_DeviceStatus_DeviceGuard01) SetPropertyVirtualizationBasedSecurityStatus(value int32) (err error) {
	return instance.SetProperty("VirtualizationBasedSecurityStatus", (value))
}

// GetVirtualizationBasedSecurityStatus gets the value of VirtualizationBasedSecurityStatus for the instance
func (instance *MDM_DeviceStatus_DeviceGuard01) GetPropertyVirtualizationBasedSecurityStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("VirtualizationBasedSecurityStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}
