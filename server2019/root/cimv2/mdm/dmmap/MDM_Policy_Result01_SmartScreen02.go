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

// MDM_Policy_Result01_SmartScreen02 struct
type MDM_Policy_Result01_SmartScreen02 struct {
	cim.WmiInstance

	//
	EnableAppInstallControl int32

	//
	EnableSmartScreenInShell int32

	//
	InstanceID string

	//
	ParentID string

	//
	PreventOverrideForFilesInShell int32
}

// SetEnableAppInstallControl sets the value of EnableAppInstallControl for the instance
func (instance *MDM_Policy_Result01_SmartScreen02) SetPropertyEnableAppInstallControl(value int32) (err error) {
	return instance.SetProperty("EnableAppInstallControl", value)
}

// GetEnableAppInstallControl gets the value of EnableAppInstallControl for the instance
func (instance *MDM_Policy_Result01_SmartScreen02) GetPropertyEnableAppInstallControl() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableAppInstallControl")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableSmartScreenInShell sets the value of EnableSmartScreenInShell for the instance
func (instance *MDM_Policy_Result01_SmartScreen02) SetPropertyEnableSmartScreenInShell(value int32) (err error) {
	return instance.SetProperty("EnableSmartScreenInShell", value)
}

// GetEnableSmartScreenInShell gets the value of EnableSmartScreenInShell for the instance
func (instance *MDM_Policy_Result01_SmartScreen02) GetPropertyEnableSmartScreenInShell() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableSmartScreenInShell")
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
func (instance *MDM_Policy_Result01_SmartScreen02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_SmartScreen02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_SmartScreen02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_SmartScreen02) GetPropertyParentID() (value string, err error) {
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

// SetPreventOverrideForFilesInShell sets the value of PreventOverrideForFilesInShell for the instance
func (instance *MDM_Policy_Result01_SmartScreen02) SetPropertyPreventOverrideForFilesInShell(value int32) (err error) {
	return instance.SetProperty("PreventOverrideForFilesInShell", value)
}

// GetPreventOverrideForFilesInShell gets the value of PreventOverrideForFilesInShell for the instance
func (instance *MDM_Policy_Result01_SmartScreen02) GetPropertyPreventOverrideForFilesInShell() (value int32, err error) {
	retValue, err := instance.GetProperty("PreventOverrideForFilesInShell")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
