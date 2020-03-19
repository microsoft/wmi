// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Policy_Config01_DeviceGuard02 struct
type MDM_Policy_Config01_DeviceGuard02 struct {
	*cim.WmiInstance

	//
	ConfigureSystemGuardLaunch int32

	//
	EnableVirtualizationBasedSecurity int32

	//
	InstanceID string

	//
	LsaCfgFlags int32

	//
	ParentID string

	//
	RequirePlatformSecurityFeatures int32
}

func NewMDM_Policy_Config01_DeviceGuard02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_DeviceGuard02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_DeviceGuard02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_DeviceGuard02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_DeviceGuard02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_DeviceGuard02{
		WmiInstance: tmp,
	}
	return
}

// SetConfigureSystemGuardLaunch sets the value of ConfigureSystemGuardLaunch for the instance
func (instance *MDM_Policy_Config01_DeviceGuard02) SetPropertyConfigureSystemGuardLaunch(value int32) (err error) {
	return instance.SetProperty("ConfigureSystemGuardLaunch", value)
}

// GetConfigureSystemGuardLaunch gets the value of ConfigureSystemGuardLaunch for the instance
func (instance *MDM_Policy_Config01_DeviceGuard02) GetPropertyConfigureSystemGuardLaunch() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureSystemGuardLaunch")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableVirtualizationBasedSecurity sets the value of EnableVirtualizationBasedSecurity for the instance
func (instance *MDM_Policy_Config01_DeviceGuard02) SetPropertyEnableVirtualizationBasedSecurity(value int32) (err error) {
	return instance.SetProperty("EnableVirtualizationBasedSecurity", value)
}

// GetEnableVirtualizationBasedSecurity gets the value of EnableVirtualizationBasedSecurity for the instance
func (instance *MDM_Policy_Config01_DeviceGuard02) GetPropertyEnableVirtualizationBasedSecurity() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableVirtualizationBasedSecurity")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_DeviceGuard02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_DeviceGuard02) GetPropertyInstanceID() (value string, err error) {
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

// SetLsaCfgFlags sets the value of LsaCfgFlags for the instance
func (instance *MDM_Policy_Config01_DeviceGuard02) SetPropertyLsaCfgFlags(value int32) (err error) {
	return instance.SetProperty("LsaCfgFlags", value)
}

// GetLsaCfgFlags gets the value of LsaCfgFlags for the instance
func (instance *MDM_Policy_Config01_DeviceGuard02) GetPropertyLsaCfgFlags() (value int32, err error) {
	retValue, err := instance.GetProperty("LsaCfgFlags")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_DeviceGuard02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_DeviceGuard02) GetPropertyParentID() (value string, err error) {
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

// SetRequirePlatformSecurityFeatures sets the value of RequirePlatformSecurityFeatures for the instance
func (instance *MDM_Policy_Config01_DeviceGuard02) SetPropertyRequirePlatformSecurityFeatures(value int32) (err error) {
	return instance.SetProperty("RequirePlatformSecurityFeatures", value)
}

// GetRequirePlatformSecurityFeatures gets the value of RequirePlatformSecurityFeatures for the instance
func (instance *MDM_Policy_Config01_DeviceGuard02) GetPropertyRequirePlatformSecurityFeatures() (value int32, err error) {
	retValue, err := instance.GetProperty("RequirePlatformSecurityFeatures")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
