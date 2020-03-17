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

// MDM_EnrollmentStatusTracking_Setup01 struct
type MDM_EnrollmentStatusTracking_Setup01 struct {
	cim.WmiInstance

	//
	HasProvisioningCompleted bool

	//
	InstanceID string

	//
	ParentID string
}

// SetHasProvisioningCompleted sets the value of HasProvisioningCompleted for the instance
func (instance *MDM_EnrollmentStatusTracking_Setup01) SetPropertyHasProvisioningCompleted(value bool) (err error) {
	return instance.SetProperty("HasProvisioningCompleted", value)
}

// GetHasProvisioningCompleted gets the value of HasProvisioningCompleted for the instance
func (instance *MDM_EnrollmentStatusTracking_Setup01) GetPropertyHasProvisioningCompleted() (value bool, err error) {
	retValue, err := instance.GetProperty("HasProvisioningCompleted")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_EnrollmentStatusTracking_Setup01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnrollmentStatusTracking_Setup01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_EnrollmentStatusTracking_Setup01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnrollmentStatusTracking_Setup01) GetPropertyParentID() (value string, err error) {
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
