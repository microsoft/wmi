// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessMonitoringConnection struct
type RemoteAccessMonitoringConnection struct {
	*RemoteAccessConnection

	//
	Bandwidth uint32

	//
	RoutingDomain string

	//
	UserActivityState string

	//
	UserName []string
}

func NewRemoteAccessMonitoringConnectionEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessMonitoringConnection, err error) {
	tmp, err := NewRemoteAccessConnectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessMonitoringConnection{
		RemoteAccessConnection: tmp,
	}
	return
}

func NewRemoteAccessMonitoringConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessMonitoringConnection, err error) {
	tmp, err := NewRemoteAccessConnectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessMonitoringConnection{
		RemoteAccessConnection: tmp,
	}
	return
}

// SetBandwidth sets the value of Bandwidth for the instance
func (instance *RemoteAccessMonitoringConnection) SetPropertyBandwidth(value uint32) (err error) {
	return instance.SetProperty("Bandwidth", (value))
}

// GetBandwidth gets the value of Bandwidth for the instance
func (instance *RemoteAccessMonitoringConnection) GetPropertyBandwidth() (value uint32, err error) {
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

// SetRoutingDomain sets the value of RoutingDomain for the instance
func (instance *RemoteAccessMonitoringConnection) SetPropertyRoutingDomain(value string) (err error) {
	return instance.SetProperty("RoutingDomain", (value))
}

// GetRoutingDomain gets the value of RoutingDomain for the instance
func (instance *RemoteAccessMonitoringConnection) GetPropertyRoutingDomain() (value string, err error) {
	retValue, err := instance.GetProperty("RoutingDomain")
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

// SetUserActivityState sets the value of UserActivityState for the instance
func (instance *RemoteAccessMonitoringConnection) SetPropertyUserActivityState(value string) (err error) {
	return instance.SetProperty("UserActivityState", (value))
}

// GetUserActivityState gets the value of UserActivityState for the instance
func (instance *RemoteAccessMonitoringConnection) GetPropertyUserActivityState() (value string, err error) {
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
func (instance *RemoteAccessMonitoringConnection) SetPropertyUserName(value []string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *RemoteAccessMonitoringConnection) GetPropertyUserName() (value []string, err error) {
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
