// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess.Server
//
// ////////////////////////////////////////////
package server

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessMonitoringConnectionLocal struct
type RemoteAccessMonitoringConnectionLocal struct {
	*RemoteAccessConnectionLocal

	//
	Bandwidth uint32

	//
	UserActivityState string

	//
	UserName []string
}

func NewRemoteAccessMonitoringConnectionLocalEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessMonitoringConnectionLocal, err error) {
	tmp, err := NewRemoteAccessConnectionLocalEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessMonitoringConnectionLocal{
		RemoteAccessConnectionLocal: tmp,
	}
	return
}

func NewRemoteAccessMonitoringConnectionLocalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessMonitoringConnectionLocal, err error) {
	tmp, err := NewRemoteAccessConnectionLocalEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessMonitoringConnectionLocal{
		RemoteAccessConnectionLocal: tmp,
	}
	return
}

// SetBandwidth sets the value of Bandwidth for the instance
func (instance *RemoteAccessMonitoringConnectionLocal) SetPropertyBandwidth(value uint32) (err error) {
	return instance.SetProperty("Bandwidth", (value))
}

// GetBandwidth gets the value of Bandwidth for the instance
func (instance *RemoteAccessMonitoringConnectionLocal) GetPropertyBandwidth() (value uint32, err error) {
	retValue, err := instance.GetProperty("Bandwidth")
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

// SetUserActivityState sets the value of UserActivityState for the instance
func (instance *RemoteAccessMonitoringConnectionLocal) SetPropertyUserActivityState(value string) (err error) {
	return instance.SetProperty("UserActivityState", (value))
}

// GetUserActivityState gets the value of UserActivityState for the instance
func (instance *RemoteAccessMonitoringConnectionLocal) GetPropertyUserActivityState() (value string, err error) {
	retValue, err := instance.GetProperty("UserActivityState")
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

// SetUserName sets the value of UserName for the instance
func (instance *RemoteAccessMonitoringConnectionLocal) SetPropertyUserName(value []string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *RemoteAccessMonitoringConnectionLocal) GetPropertyUserName() (value []string, err error) {
	retValue, err := instance.GetProperty("UserName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
