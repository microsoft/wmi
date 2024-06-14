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

// VpnAuth struct
type VpnAuth struct {
	*cim.WmiInstance

	//
	RadiusServerList []string

	//
	Type string
}

func NewVpnAuthEx1(instance *cim.WmiInstance) (newInstance *VpnAuth, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &VpnAuth{
		WmiInstance: tmp,
	}
	return
}

func NewVpnAuthEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnAuth, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnAuth{
		WmiInstance: tmp,
	}
	return
}

// SetRadiusServerList sets the value of RadiusServerList for the instance
func (instance *VpnAuth) SetPropertyRadiusServerList(value []string) (err error) {
	return instance.SetProperty("RadiusServerList", (value))
}

// GetRadiusServerList gets the value of RadiusServerList for the instance
func (instance *VpnAuth) GetPropertyRadiusServerList() (value []string, err error) {
	retValue, err := instance.GetProperty("RadiusServerList")
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

// SetType sets the value of Type for the instance
func (instance *VpnAuth) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *VpnAuth) GetPropertyType() (value string, err error) {
	retValue, err := instance.GetProperty("Type")
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
