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

// MDM_EnrollmentStatusTracking_Tracking03_02 struct
type MDM_EnrollmentStatusTracking_Tracking03_02 struct {
	cim.WmiInstance

	//
	InstallationState int32

	//
	InstanceID string

	//
	ParentID string

	//
	RebootRequired int32

	//
	TrackingUri string
}

// SetInstallationState sets the value of InstallationState for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) SetPropertyInstallationState(value int32) (err error) {
	return instance.SetProperty("InstallationState", value)
}

// GetInstallationState gets the value of InstallationState for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) GetPropertyInstallationState() (value int32, err error) {
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
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) GetPropertyParentID() (value string, err error) {
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

// SetRebootRequired sets the value of RebootRequired for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) SetPropertyRebootRequired(value int32) (err error) {
	return instance.SetProperty("RebootRequired", value)
}

// GetRebootRequired gets the value of RebootRequired for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) GetPropertyRebootRequired() (value int32, err error) {
	retValue, err := instance.GetProperty("RebootRequired")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrackingUri sets the value of TrackingUri for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) SetPropertyTrackingUri(value string) (err error) {
	return instance.SetProperty("TrackingUri", value)
}

// GetTrackingUri gets the value of TrackingUri for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) GetPropertyTrackingUri() (value string, err error) {
	retValue, err := instance.GetProperty("TrackingUri")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
