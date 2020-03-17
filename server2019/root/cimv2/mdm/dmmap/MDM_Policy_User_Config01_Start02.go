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

// MDM_Policy_User_Config01_Start02 struct
type MDM_Policy_User_Config01_Start02 struct {
	cim.WmiInstance

	//
	DisableContextMenus int32

	//
	ForceStartSize int32

	//
	HideAppList int32

	//
	HideFrequentlyUsedApps int32

	//
	HidePeopleBar int32

	//
	HideRecentJumplists int32

	//
	HideRecentlyAddedApps int32

	//
	InstanceID string

	//
	ParentID string

	//
	StartLayout string
}

// SetDisableContextMenus sets the value of DisableContextMenus for the instance
func (instance *MDM_Policy_User_Config01_Start02) SetPropertyDisableContextMenus(value int32) (err error) {
	return instance.SetProperty("DisableContextMenus", value)
}

// GetDisableContextMenus gets the value of DisableContextMenus for the instance
func (instance *MDM_Policy_User_Config01_Start02) GetPropertyDisableContextMenus() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableContextMenus")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetForceStartSize sets the value of ForceStartSize for the instance
func (instance *MDM_Policy_User_Config01_Start02) SetPropertyForceStartSize(value int32) (err error) {
	return instance.SetProperty("ForceStartSize", value)
}

// GetForceStartSize gets the value of ForceStartSize for the instance
func (instance *MDM_Policy_User_Config01_Start02) GetPropertyForceStartSize() (value int32, err error) {
	retValue, err := instance.GetProperty("ForceStartSize")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHideAppList sets the value of HideAppList for the instance
func (instance *MDM_Policy_User_Config01_Start02) SetPropertyHideAppList(value int32) (err error) {
	return instance.SetProperty("HideAppList", value)
}

// GetHideAppList gets the value of HideAppList for the instance
func (instance *MDM_Policy_User_Config01_Start02) GetPropertyHideAppList() (value int32, err error) {
	retValue, err := instance.GetProperty("HideAppList")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHideFrequentlyUsedApps sets the value of HideFrequentlyUsedApps for the instance
func (instance *MDM_Policy_User_Config01_Start02) SetPropertyHideFrequentlyUsedApps(value int32) (err error) {
	return instance.SetProperty("HideFrequentlyUsedApps", value)
}

// GetHideFrequentlyUsedApps gets the value of HideFrequentlyUsedApps for the instance
func (instance *MDM_Policy_User_Config01_Start02) GetPropertyHideFrequentlyUsedApps() (value int32, err error) {
	retValue, err := instance.GetProperty("HideFrequentlyUsedApps")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHidePeopleBar sets the value of HidePeopleBar for the instance
func (instance *MDM_Policy_User_Config01_Start02) SetPropertyHidePeopleBar(value int32) (err error) {
	return instance.SetProperty("HidePeopleBar", value)
}

// GetHidePeopleBar gets the value of HidePeopleBar for the instance
func (instance *MDM_Policy_User_Config01_Start02) GetPropertyHidePeopleBar() (value int32, err error) {
	retValue, err := instance.GetProperty("HidePeopleBar")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHideRecentJumplists sets the value of HideRecentJumplists for the instance
func (instance *MDM_Policy_User_Config01_Start02) SetPropertyHideRecentJumplists(value int32) (err error) {
	return instance.SetProperty("HideRecentJumplists", value)
}

// GetHideRecentJumplists gets the value of HideRecentJumplists for the instance
func (instance *MDM_Policy_User_Config01_Start02) GetPropertyHideRecentJumplists() (value int32, err error) {
	retValue, err := instance.GetProperty("HideRecentJumplists")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHideRecentlyAddedApps sets the value of HideRecentlyAddedApps for the instance
func (instance *MDM_Policy_User_Config01_Start02) SetPropertyHideRecentlyAddedApps(value int32) (err error) {
	return instance.SetProperty("HideRecentlyAddedApps", value)
}

// GetHideRecentlyAddedApps gets the value of HideRecentlyAddedApps for the instance
func (instance *MDM_Policy_User_Config01_Start02) GetPropertyHideRecentlyAddedApps() (value int32, err error) {
	retValue, err := instance.GetProperty("HideRecentlyAddedApps")
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
func (instance *MDM_Policy_User_Config01_Start02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Config01_Start02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_User_Config01_Start02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_User_Config01_Start02) GetPropertyParentID() (value string, err error) {
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

// SetStartLayout sets the value of StartLayout for the instance
func (instance *MDM_Policy_User_Config01_Start02) SetPropertyStartLayout(value string) (err error) {
	return instance.SetProperty("StartLayout", value)
}

// GetStartLayout gets the value of StartLayout for the instance
func (instance *MDM_Policy_User_Config01_Start02) GetPropertyStartLayout() (value string, err error) {
	retValue, err := instance.GetProperty("StartLayout")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
