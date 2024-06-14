// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DAMgmtHost struct
type DAMgmtHost struct {
	*cim.WmiInstance

	//
	Host string

	//
	IPAddresses []string
}

func NewDAMgmtHostEx1(instance *cim.WmiInstance) (newInstance *DAMgmtHost, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DAMgmtHost{
		WmiInstance: tmp,
	}
	return
}

func NewDAMgmtHostEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DAMgmtHost, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DAMgmtHost{
		WmiInstance: tmp,
	}
	return
}

// SetHost sets the value of Host for the instance
func (instance *DAMgmtHost) SetPropertyHost(value string) (err error) {
	return instance.SetProperty("Host", (value))
}

// GetHost gets the value of Host for the instance
func (instance *DAMgmtHost) GetPropertyHost() (value string, err error) {
	retValue, err := instance.GetProperty("Host")
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

// SetIPAddresses sets the value of IPAddresses for the instance
func (instance *DAMgmtHost) SetPropertyIPAddresses(value []string) (err error) {
	return instance.SetProperty("IPAddresses", (value))
}

// GetIPAddresses gets the value of IPAddresses for the instance
func (instance *DAMgmtHost) GetPropertyIPAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("IPAddresses")
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
