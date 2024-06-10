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

// DAAppServerConnection struct
type DAAppServerConnection struct {
	*cim.WmiInstance

	//
	ConnectionType string

	//
	TrafficProtection string
}

func NewDAAppServerConnectionEx1(instance *cim.WmiInstance) (newInstance *DAAppServerConnection, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DAAppServerConnection{
		WmiInstance: tmp,
	}
	return
}

func NewDAAppServerConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DAAppServerConnection, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DAAppServerConnection{
		WmiInstance: tmp,
	}
	return
}

// SetConnectionType sets the value of ConnectionType for the instance
func (instance *DAAppServerConnection) SetPropertyConnectionType(value string) (err error) {
	return instance.SetProperty("ConnectionType", (value))
}

// GetConnectionType gets the value of ConnectionType for the instance
func (instance *DAAppServerConnection) GetPropertyConnectionType() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionType")
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

// SetTrafficProtection sets the value of TrafficProtection for the instance
func (instance *DAAppServerConnection) SetPropertyTrafficProtection(value string) (err error) {
	return instance.SetProperty("TrafficProtection", (value))
}

// GetTrafficProtection gets the value of TrafficProtection for the instance
func (instance *DAAppServerConnection) GetPropertyTrafficProtection() (value string, err error) {
	retValue, err := instance.GetProperty("TrafficProtection")
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
