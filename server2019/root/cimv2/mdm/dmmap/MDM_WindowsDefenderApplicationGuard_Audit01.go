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

// MDM_WindowsDefenderApplicationGuard_Audit01 struct
type MDM_WindowsDefenderApplicationGuard_Audit01 struct {
	cim.WmiInstance

	//
	AuditApplicationGuard int32

	//
	InstanceID string

	//
	ParentID string
}

// SetAuditApplicationGuard sets the value of AuditApplicationGuard for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Audit01) SetPropertyAuditApplicationGuard(value int32) (err error) {
	return instance.SetProperty("AuditApplicationGuard", value)
}

// GetAuditApplicationGuard gets the value of AuditApplicationGuard for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Audit01) GetPropertyAuditApplicationGuard() (value int32, err error) {
	retValue, err := instance.GetProperty("AuditApplicationGuard")
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
func (instance *MDM_WindowsDefenderApplicationGuard_Audit01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Audit01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_WindowsDefenderApplicationGuard_Audit01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Audit01) GetPropertyParentID() (value string, err error) {
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
