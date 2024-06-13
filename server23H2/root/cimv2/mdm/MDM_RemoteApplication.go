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

// MDM_RemoteApplication struct
type MDM_RemoteApplication struct {
	*cim.WmiInstance

	//
	AppId string

	//
	FeedUrl string
}

func NewMDM_RemoteApplicationEx1(instance *cim.WmiInstance) (newInstance *MDM_RemoteApplication, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_RemoteApplication{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_RemoteApplicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_RemoteApplication, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_RemoteApplication{
		WmiInstance: tmp,
	}
	return
}

// SetAppId sets the value of AppId for the instance
func (instance *MDM_RemoteApplication) SetPropertyAppId(value string) (err error) {
	return instance.SetProperty("AppId", (value))
}

// GetAppId gets the value of AppId for the instance
func (instance *MDM_RemoteApplication) GetPropertyAppId() (value string, err error) {
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

// SetFeedUrl sets the value of FeedUrl for the instance
func (instance *MDM_RemoteApplication) SetPropertyFeedUrl(value string) (err error) {
	return instance.SetProperty("FeedUrl", (value))
}

// GetFeedUrl gets the value of FeedUrl for the instance
func (instance *MDM_RemoteApplication) GetPropertyFeedUrl() (value string, err error) {
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
