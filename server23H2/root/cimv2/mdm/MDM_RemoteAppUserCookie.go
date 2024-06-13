// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_RemoteAppUserCookie struct
type MDM_RemoteAppUserCookie struct {
	*cim.WmiInstance

	//
	Cookie string

	//
	CookieHash string

	//
	FeedUrl string
}

func NewMDM_RemoteAppUserCookieEx1(instance *cim.WmiInstance) (newInstance *MDM_RemoteAppUserCookie, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_RemoteAppUserCookie{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_RemoteAppUserCookieEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_RemoteAppUserCookie, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_RemoteAppUserCookie{
		WmiInstance: tmp,
	}
	return
}

// SetCookie sets the value of Cookie for the instance
func (instance *MDM_RemoteAppUserCookie) SetPropertyCookie(value string) (err error) {
	return instance.SetProperty("Cookie", (value))
}

// GetCookie gets the value of Cookie for the instance
func (instance *MDM_RemoteAppUserCookie) GetPropertyCookie() (value string, err error) {
	retValue, err := instance.GetProperty("Cookie")
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

// SetCookieHash sets the value of CookieHash for the instance
func (instance *MDM_RemoteAppUserCookie) SetPropertyCookieHash(value string) (err error) {
	return instance.SetProperty("CookieHash", (value))
}

// GetCookieHash gets the value of CookieHash for the instance
func (instance *MDM_RemoteAppUserCookie) GetPropertyCookieHash() (value string, err error) {
	retValue, err := instance.GetProperty("CookieHash")
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

// SetFeedUrl sets the value of FeedUrl for the instance
func (instance *MDM_RemoteAppUserCookie) SetPropertyFeedUrl(value string) (err error) {
	return instance.SetProperty("FeedUrl", (value))
}

// GetFeedUrl gets the value of FeedUrl for the instance
func (instance *MDM_RemoteAppUserCookie) GetPropertyFeedUrl() (value string, err error) {
	retValue, err := instance.GetProperty("FeedUrl")
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
