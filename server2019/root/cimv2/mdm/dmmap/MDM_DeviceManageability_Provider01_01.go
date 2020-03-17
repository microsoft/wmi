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

// MDM_DeviceManageability_Provider01_01 struct
type MDM_DeviceManageability_Provider01_01 struct {
	cim.WmiInstance

	//
	ConfigInfo string

	//
	EnrollmentInfo string

	//
	InstanceID string

	//
	ParentID string
}

// SetConfigInfo sets the value of ConfigInfo for the instance
func (instance *MDM_DeviceManageability_Provider01_01) SetPropertyConfigInfo(value string) (err error) {
	return instance.SetProperty("ConfigInfo", value)
}

// GetConfigInfo gets the value of ConfigInfo for the instance
func (instance *MDM_DeviceManageability_Provider01_01) GetPropertyConfigInfo() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnrollmentInfo sets the value of EnrollmentInfo for the instance
func (instance *MDM_DeviceManageability_Provider01_01) SetPropertyEnrollmentInfo(value string) (err error) {
	return instance.SetProperty("EnrollmentInfo", value)
}

// GetEnrollmentInfo gets the value of EnrollmentInfo for the instance
func (instance *MDM_DeviceManageability_Provider01_01) GetPropertyEnrollmentInfo() (value string, err error) {
	retValue, err := instance.GetProperty("EnrollmentInfo")
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
func (instance *MDM_DeviceManageability_Provider01_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DeviceManageability_Provider01_01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_DeviceManageability_Provider01_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DeviceManageability_Provider01_01) GetPropertyParentID() (value string, err error) {
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
