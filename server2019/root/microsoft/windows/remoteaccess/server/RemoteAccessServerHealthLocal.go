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
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessServerHealthLocal struct
type RemoteAccessServerHealthLocal struct {
	*cim.WmiInstance

	//
	Data []uint8

	//
	HealthStatus uint32

	//
	NumOfHealthMonitors uint32
}

func NewRemoteAccessServerHealthLocalEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessServerHealthLocal, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RemoteAccessServerHealthLocal{
		WmiInstance: tmp,
	}
	return
}

func NewRemoteAccessServerHealthLocalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessServerHealthLocal, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessServerHealthLocal{
		WmiInstance: tmp,
	}
	return
}

// SetData sets the value of Data for the instance
func (instance *RemoteAccessServerHealthLocal) SetPropertyData(value []uint8) (err error) {
	return instance.SetProperty("Data", (value))
}

// GetData gets the value of Data for the instance
func (instance *RemoteAccessServerHealthLocal) GetPropertyData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *RemoteAccessServerHealthLocal) SetPropertyHealthStatus(value uint32) (err error) {
	return instance.SetProperty("HealthStatus", (value))
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *RemoteAccessServerHealthLocal) GetPropertyHealthStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
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

// SetNumOfHealthMonitors sets the value of NumOfHealthMonitors for the instance
func (instance *RemoteAccessServerHealthLocal) SetPropertyNumOfHealthMonitors(value uint32) (err error) {
	return instance.SetProperty("NumOfHealthMonitors", (value))
}

// GetNumOfHealthMonitors gets the value of NumOfHealthMonitors for the instance
func (instance *RemoteAccessServerHealthLocal) GetPropertyNumOfHealthMonitors() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumOfHealthMonitors")
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
