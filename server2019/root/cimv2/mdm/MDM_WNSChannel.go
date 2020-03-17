// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_WNSChannel struct
type MDM_WNSChannel struct {
	cim.WmiInstance

	//
	AppId string

	//
	Channel string

	//
	ChannelStatus uint32

	//
	ExpirationTime string

	//
	LastError uint32

	//
	UserSID string
}

// SetAppId sets the value of AppId for the instance
func (instance *MDM_WNSChannel) SetPropertyAppId(value string) (err error) {
	return instance.SetProperty("AppId", value)
}

// GetAppId gets the value of AppId for the instance
func (instance *MDM_WNSChannel) GetPropertyAppId() (value string, err error) {
	retValue, err := instance.GetProperty("AppId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetChannel sets the value of Channel for the instance
func (instance *MDM_WNSChannel) SetPropertyChannel(value string) (err error) {
	return instance.SetProperty("Channel", value)
}

// GetChannel gets the value of Channel for the instance
func (instance *MDM_WNSChannel) GetPropertyChannel() (value string, err error) {
	retValue, err := instance.GetProperty("Channel")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetChannelStatus sets the value of ChannelStatus for the instance
func (instance *MDM_WNSChannel) SetPropertyChannelStatus(value uint32) (err error) {
	return instance.SetProperty("ChannelStatus", value)
}

// GetChannelStatus gets the value of ChannelStatus for the instance
func (instance *MDM_WNSChannel) GetPropertyChannelStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("ChannelStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExpirationTime sets the value of ExpirationTime for the instance
func (instance *MDM_WNSChannel) SetPropertyExpirationTime(value string) (err error) {
	return instance.SetProperty("ExpirationTime", value)
}

// GetExpirationTime gets the value of ExpirationTime for the instance
func (instance *MDM_WNSChannel) GetPropertyExpirationTime() (value string, err error) {
	retValue, err := instance.GetProperty("ExpirationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastError sets the value of LastError for the instance
func (instance *MDM_WNSChannel) SetPropertyLastError(value uint32) (err error) {
	return instance.SetProperty("LastError", value)
}

// GetLastError gets the value of LastError for the instance
func (instance *MDM_WNSChannel) GetPropertyLastError() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastError")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserSID sets the value of UserSID for the instance
func (instance *MDM_WNSChannel) SetPropertyUserSID(value string) (err error) {
	return instance.SetProperty("UserSID", value)
}

// GetUserSID gets the value of UserSID for the instance
func (instance *MDM_WNSChannel) GetPropertyUserSID() (value string, err error) {
	retValue, err := instance.GetProperty("UserSID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
