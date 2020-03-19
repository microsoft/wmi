// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_UserStateConfigurationControls struct
type Win32_UserStateConfigurationControls struct {
	*cim.WmiInstance

	// Controls whether the computer's folder redirection feature settings are configured by using UST Manageability WMI classes or by using Group Policy.
	FolderRedirection UserStateConfigurationControls_FolderRedirection

	// Controls whether the computer's Offline files feature settings are configured by using UST Manageability WMI classes or by using Group Policy.
	OfflineFiles UserStateConfigurationControls_OfflineFiles

	// Controls whether the computer's roaming user profile feature settings are configured by using UST Manageability WMI classes or by using Group Policy.
	RoamingUserProfile UserStateConfigurationControls_RoamingUserProfile
}

func NewWin32_UserStateConfigurationControlsEx1(instance *cim.WmiInstance) (newInstance *Win32_UserStateConfigurationControls, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Win32_UserStateConfigurationControls{
		WmiInstance: tmp,
	}
	return
}

func NewWin32_UserStateConfigurationControlsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_UserStateConfigurationControls, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_UserStateConfigurationControls{
		WmiInstance: tmp,
	}
	return
}

// SetFolderRedirection sets the value of FolderRedirection for the instance
func (instance *Win32_UserStateConfigurationControls) SetPropertyFolderRedirection(value UserStateConfigurationControls_FolderRedirection) (err error) {
	return instance.SetProperty("FolderRedirection", value)
}

// GetFolderRedirection gets the value of FolderRedirection for the instance
func (instance *Win32_UserStateConfigurationControls) GetPropertyFolderRedirection() (value UserStateConfigurationControls_FolderRedirection, err error) {
	retValue, err := instance.GetProperty("FolderRedirection")
	if err != nil {
		return
	}
	value, ok := retValue.(UserStateConfigurationControls_FolderRedirection)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOfflineFiles sets the value of OfflineFiles for the instance
func (instance *Win32_UserStateConfigurationControls) SetPropertyOfflineFiles(value UserStateConfigurationControls_OfflineFiles) (err error) {
	return instance.SetProperty("OfflineFiles", value)
}

// GetOfflineFiles gets the value of OfflineFiles for the instance
func (instance *Win32_UserStateConfigurationControls) GetPropertyOfflineFiles() (value UserStateConfigurationControls_OfflineFiles, err error) {
	retValue, err := instance.GetProperty("OfflineFiles")
	if err != nil {
		return
	}
	value, ok := retValue.(UserStateConfigurationControls_OfflineFiles)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoamingUserProfile sets the value of RoamingUserProfile for the instance
func (instance *Win32_UserStateConfigurationControls) SetPropertyRoamingUserProfile(value UserStateConfigurationControls_RoamingUserProfile) (err error) {
	return instance.SetProperty("RoamingUserProfile", value)
}

// GetRoamingUserProfile gets the value of RoamingUserProfile for the instance
func (instance *Win32_UserStateConfigurationControls) GetPropertyRoamingUserProfile() (value UserStateConfigurationControls_RoamingUserProfile, err error) {
	retValue, err := instance.GetProperty("RoamingUserProfile")
	if err != nil {
		return
	}
	value, ok := retValue.(UserStateConfigurationControls_RoamingUserProfile)
	if !ok {
		// TODO: Set an error
	}
	return
}
