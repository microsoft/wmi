// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessRadiusAccounting struct
type RemoteAccessRadiusAccounting struct {
	*cim.WmiInstance

	//
	RadiusAccountingStatus string

	//
	RemoteServerList []RemoteAccessRadiusServer
}

func NewRemoteAccessRadiusAccountingEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessRadiusAccounting, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RemoteAccessRadiusAccounting{
		WmiInstance: tmp,
	}
	return
}

func NewRemoteAccessRadiusAccountingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessRadiusAccounting, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessRadiusAccounting{
		WmiInstance: tmp,
	}
	return
}

// SetRadiusAccountingStatus sets the value of RadiusAccountingStatus for the instance
func (instance *RemoteAccessRadiusAccounting) SetPropertyRadiusAccountingStatus(value string) (err error) {
	return instance.SetProperty("RadiusAccountingStatus", (value))
}

// GetRadiusAccountingStatus gets the value of RadiusAccountingStatus for the instance
func (instance *RemoteAccessRadiusAccounting) GetPropertyRadiusAccountingStatus() (value string, err error) {
	retValue, err := instance.GetProperty("RadiusAccountingStatus")
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

// SetRemoteServerList sets the value of RemoteServerList for the instance
func (instance *RemoteAccessRadiusAccounting) SetPropertyRemoteServerList(value []RemoteAccessRadiusServer) (err error) {
	return instance.SetProperty("RemoteServerList", (value))
}

// GetRemoteServerList gets the value of RemoteServerList for the instance
func (instance *RemoteAccessRadiusAccounting) GetPropertyRemoteServerList() (value []RemoteAccessRadiusServer, err error) {
	retValue, err := instance.GetProperty("RemoteServerList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(RemoteAccessRadiusServer)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " RemoteAccessRadiusServer is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, RemoteAccessRadiusServer(valuetmp))
	}

	return
}
