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

// MDM_ActiveSync_User_Options03 struct
type MDM_ActiveSync_User_Options03 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	Logging string

	//
	MailAgeFilter string

	//
	ParentID string

	//
	Schedule string

	//
	UseSSL string
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_ActiveSync_User_Options03) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_ActiveSync_User_Options03) GetPropertyInstanceID() (value string, err error) {
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

// SetLogging sets the value of Logging for the instance
func (instance *MDM_ActiveSync_User_Options03) SetPropertyLogging(value string) (err error) {
	return instance.SetProperty("Logging", value)
}

// GetLogging gets the value of Logging for the instance
func (instance *MDM_ActiveSync_User_Options03) GetPropertyLogging() (value string, err error) {
	retValue, err := instance.GetProperty("Logging")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMailAgeFilter sets the value of MailAgeFilter for the instance
func (instance *MDM_ActiveSync_User_Options03) SetPropertyMailAgeFilter(value string) (err error) {
	return instance.SetProperty("MailAgeFilter", value)
}

// GetMailAgeFilter gets the value of MailAgeFilter for the instance
func (instance *MDM_ActiveSync_User_Options03) GetPropertyMailAgeFilter() (value string, err error) {
	retValue, err := instance.GetProperty("MailAgeFilter")
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
func (instance *MDM_ActiveSync_User_Options03) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_ActiveSync_User_Options03) GetPropertyParentID() (value string, err error) {
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

// SetSchedule sets the value of Schedule for the instance
func (instance *MDM_ActiveSync_User_Options03) SetPropertySchedule(value string) (err error) {
	return instance.SetProperty("Schedule", value)
}

// GetSchedule gets the value of Schedule for the instance
func (instance *MDM_ActiveSync_User_Options03) GetPropertySchedule() (value string, err error) {
	retValue, err := instance.GetProperty("Schedule")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUseSSL sets the value of UseSSL for the instance
func (instance *MDM_ActiveSync_User_Options03) SetPropertyUseSSL(value string) (err error) {
	return instance.SetProperty("UseSSL", value)
}

// GetUseSSL gets the value of UseSSL for the instance
func (instance *MDM_ActiveSync_User_Options03) GetPropertyUseSSL() (value string, err error) {
	retValue, err := instance.GetProperty("UseSSL")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
