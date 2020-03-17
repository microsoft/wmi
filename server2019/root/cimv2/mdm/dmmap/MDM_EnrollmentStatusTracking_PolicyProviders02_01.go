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

// MDM_EnrollmentStatusTracking_PolicyProviders02_01 struct
type MDM_EnrollmentStatusTracking_PolicyProviders02_01 struct {
	cim.WmiInstance

	//
	InstallationState int32

	//
	InstanceID string

	//
	LastError int32

	//
	ParentID string

	//
	Timeout int32
}

// SetInstallationState sets the value of InstallationState for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) SetPropertyInstallationState(value int32) (err error) {
	return instance.SetProperty("InstallationState", value)
}

// GetInstallationState gets the value of InstallationState for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) GetPropertyInstallationState() (value int32, err error) {
	retValue, err := instance.GetProperty("InstallationState")
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
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) GetPropertyInstanceID() (value string, err error) {
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

// SetLastError sets the value of LastError for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) SetPropertyLastError(value int32) (err error) {
	return instance.SetProperty("LastError", value)
}

// GetLastError gets the value of LastError for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) GetPropertyLastError() (value int32, err error) {
	retValue, err := instance.GetProperty("LastError")
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
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) GetPropertyParentID() (value string, err error) {
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

// SetTimeout sets the value of Timeout for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) SetPropertyTimeout(value int32) (err error) {
	return instance.SetProperty("Timeout", value)
}

// GetTimeout gets the value of Timeout for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) GetPropertyTimeout() (value int32, err error) {
	retValue, err := instance.GetProperty("Timeout")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
