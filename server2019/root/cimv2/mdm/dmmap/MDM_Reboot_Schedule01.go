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

// MDM_Reboot_Schedule01 struct
type MDM_Reboot_Schedule01 struct {
	cim.WmiInstance

	//
	DailyRecurrent string

	//
	InstanceID string

	//
	ParentID string

	//
	Single string
}

// SetDailyRecurrent sets the value of DailyRecurrent for the instance
func (instance *MDM_Reboot_Schedule01) SetPropertyDailyRecurrent(value string) (err error) {
	return instance.SetProperty("DailyRecurrent", value)
}

// GetDailyRecurrent gets the value of DailyRecurrent for the instance
func (instance *MDM_Reboot_Schedule01) GetPropertyDailyRecurrent() (value string, err error) {
	retValue, err := instance.GetProperty("DailyRecurrent")
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
func (instance *MDM_Reboot_Schedule01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Reboot_Schedule01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Reboot_Schedule01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Reboot_Schedule01) GetPropertyParentID() (value string, err error) {
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

// SetSingle sets the value of Single for the instance
func (instance *MDM_Reboot_Schedule01) SetPropertySingle(value string) (err error) {
	return instance.SetProperty("Single", value)
}

// GetSingle gets the value of Single for the instance
func (instance *MDM_Reboot_Schedule01) GetPropertySingle() (value string, err error) {
	retValue, err := instance.GetProperty("Single")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
