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

// MDM_Win32CompatibilityAppraiser_WindowsErrorReporting01 struct
type MDM_Win32CompatibilityAppraiser_WindowsErrorReporting01 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	ParentID string

	//
	WerConnectionReport string
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_WindowsErrorReporting01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_WindowsErrorReporting01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Win32CompatibilityAppraiser_WindowsErrorReporting01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_WindowsErrorReporting01) GetPropertyParentID() (value string, err error) {
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

// SetWerConnectionReport sets the value of WerConnectionReport for the instance
func (instance *MDM_Win32CompatibilityAppraiser_WindowsErrorReporting01) SetPropertyWerConnectionReport(value string) (err error) {
	return instance.SetProperty("WerConnectionReport", value)
}

// GetWerConnectionReport gets the value of WerConnectionReport for the instance
func (instance *MDM_Win32CompatibilityAppraiser_WindowsErrorReporting01) GetPropertyWerConnectionReport() (value string, err error) {
	retValue, err := instance.GetProperty("WerConnectionReport")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
