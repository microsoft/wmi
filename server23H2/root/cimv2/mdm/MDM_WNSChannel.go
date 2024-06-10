// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.mdm
//
// ////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_WNSChannel struct
type MDM_WNSChannel struct {
	*cim.WmiInstance

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

func NewMDM_WNSChannelEx1(instance *cim.WmiInstance) (newInstance *MDM_WNSChannel, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_WNSChannel{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_WNSChannelEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_WNSChannel, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_WNSChannel{
		WmiInstance: tmp,
	}
	return
}

// SetAppId sets the value of AppId for the instance
func (instance *MDM_WNSChannel) SetPropertyAppId(value string) (err error) {
	return instance.SetProperty("AppId", (value))
}

// GetAppId gets the value of AppId for the instance
func (instance *MDM_WNSChannel) GetPropertyAppId() (value string, err error) {
	retValue, err := instance.GetProperty("AppId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetChannel sets the value of Channel for the instance
func (instance *MDM_WNSChannel) SetPropertyChannel(value string) (err error) {
	return instance.SetProperty("Channel", (value))
}

// GetChannel gets the value of Channel for the instance
func (instance *MDM_WNSChannel) GetPropertyChannel() (value string, err error) {
	retValue, err := instance.GetProperty("Channel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetChannelStatus sets the value of ChannelStatus for the instance
func (instance *MDM_WNSChannel) SetPropertyChannelStatus(value uint32) (err error) {
	return instance.SetProperty("ChannelStatus", (value))
}

// GetChannelStatus gets the value of ChannelStatus for the instance
func (instance *MDM_WNSChannel) GetPropertyChannelStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("ChannelStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetExpirationTime sets the value of ExpirationTime for the instance
func (instance *MDM_WNSChannel) SetPropertyExpirationTime(value string) (err error) {
	return instance.SetProperty("ExpirationTime", (value))
}

// GetExpirationTime gets the value of ExpirationTime for the instance
func (instance *MDM_WNSChannel) GetPropertyExpirationTime() (value string, err error) {
	retValue, err := instance.GetProperty("ExpirationTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetLastError sets the value of LastError for the instance
func (instance *MDM_WNSChannel) SetPropertyLastError(value uint32) (err error) {
	return instance.SetProperty("LastError", (value))
}

// GetLastError gets the value of LastError for the instance
func (instance *MDM_WNSChannel) GetPropertyLastError() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastError")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetUserSID sets the value of UserSID for the instance
func (instance *MDM_WNSChannel) SetPropertyUserSID(value string) (err error) {
	return instance.SetProperty("UserSID", (value))
}

// GetUserSID gets the value of UserSID for the instance
func (instance *MDM_WNSChannel) GetPropertyUserSID() (value string, err error) {
	retValue, err := instance.GetProperty("UserSID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
