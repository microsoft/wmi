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

// MDM_Policy_User_Config01_Settings02 struct
type MDM_Policy_User_Config01_Settings02 struct {
	cim.WmiInstance

	//
	ConfigureTaskbarCalendar int32

	//
	InstanceID string

	//
	PageVisibilityList string

	//
	ParentID string
}

// SetConfigureTaskbarCalendar sets the value of ConfigureTaskbarCalendar for the instance
func (instance *MDM_Policy_User_Config01_Settings02) SetPropertyConfigureTaskbarCalendar(value int32) (err error) {
	return instance.SetProperty("ConfigureTaskbarCalendar", value)
}

// GetConfigureTaskbarCalendar gets the value of ConfigureTaskbarCalendar for the instance
func (instance *MDM_Policy_User_Config01_Settings02) GetPropertyConfigureTaskbarCalendar() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureTaskbarCalendar")
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
func (instance *MDM_Policy_User_Config01_Settings02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Config01_Settings02) GetPropertyInstanceID() (value string, err error) {
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

// SetPageVisibilityList sets the value of PageVisibilityList for the instance
func (instance *MDM_Policy_User_Config01_Settings02) SetPropertyPageVisibilityList(value string) (err error) {
	return instance.SetProperty("PageVisibilityList", value)
}

// GetPageVisibilityList gets the value of PageVisibilityList for the instance
func (instance *MDM_Policy_User_Config01_Settings02) GetPropertyPageVisibilityList() (value string, err error) {
	retValue, err := instance.GetProperty("PageVisibilityList")
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
func (instance *MDM_Policy_User_Config01_Settings02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_User_Config01_Settings02) GetPropertyParentID() (value string, err error) {
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
