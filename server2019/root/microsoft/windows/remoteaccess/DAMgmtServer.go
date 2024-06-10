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

// DAMgmtServer struct
type DAMgmtServer struct {
	*cim.WmiInstance

	//
	Servers []DAMgmtChangedHost

	//
	Summary string
}

func NewDAMgmtServerEx1(instance *cim.WmiInstance) (newInstance *DAMgmtServer, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DAMgmtServer{
		WmiInstance: tmp,
	}
	return
}

func NewDAMgmtServerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DAMgmtServer, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DAMgmtServer{
		WmiInstance: tmp,
	}
	return
}

// SetServers sets the value of Servers for the instance
func (instance *DAMgmtServer) SetPropertyServers(value []DAMgmtChangedHost) (err error) {
	return instance.SetProperty("Servers", (value))
}

// GetServers gets the value of Servers for the instance
func (instance *DAMgmtServer) GetPropertyServers() (value []DAMgmtChangedHost, err error) {
	retValue, err := instance.GetProperty("Servers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(DAMgmtChangedHost)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " DAMgmtChangedHost is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DAMgmtChangedHost(valuetmp))
	}

	return
}

// SetSummary sets the value of Summary for the instance
func (instance *DAMgmtServer) SetPropertySummary(value string) (err error) {
	return instance.SetProperty("Summary", (value))
}

// GetSummary gets the value of Summary for the instance
func (instance *DAMgmtServer) GetPropertySummary() (value string, err error) {
	retValue, err := instance.GetProperty("Summary")
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
