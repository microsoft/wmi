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

// MDM_Win32CompatibilityAppraiser_CompatibilityAppraiser01 struct
type MDM_Win32CompatibilityAppraiser_CompatibilityAppraiser01 struct {
	cim.WmiInstance

	//
	AppraiserRunResultReport string

	//
	InstanceID string

	//
	ParentID string
}

// SetAppraiserRunResultReport sets the value of AppraiserRunResultReport for the instance
func (instance *MDM_Win32CompatibilityAppraiser_CompatibilityAppraiser01) SetPropertyAppraiserRunResultReport(value string) (err error) {
	return instance.SetProperty("AppraiserRunResultReport", value)
}

// GetAppraiserRunResultReport gets the value of AppraiserRunResultReport for the instance
func (instance *MDM_Win32CompatibilityAppraiser_CompatibilityAppraiser01) GetPropertyAppraiserRunResultReport() (value string, err error) {
	retValue, err := instance.GetProperty("AppraiserRunResultReport")
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
func (instance *MDM_Win32CompatibilityAppraiser_CompatibilityAppraiser01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_CompatibilityAppraiser01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Win32CompatibilityAppraiser_CompatibilityAppraiser01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_CompatibilityAppraiser01) GetPropertyParentID() (value string, err error) {
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
