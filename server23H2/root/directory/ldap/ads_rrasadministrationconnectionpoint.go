// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_rrasadministrationconnectionpoint struct
type ads_rrasadministrationconnectionpoint struct {
	*ads_serviceadministrationpoint

	//
	DS_msRRASAttribute []string
}

func Newads_rrasadministrationconnectionpointEx1(instance *cim.WmiInstance) (newInstance *ads_rrasadministrationconnectionpoint, err error) {
	tmp, err := Newads_serviceadministrationpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_rrasadministrationconnectionpoint{
		ads_serviceadministrationpoint: tmp,
	}
	return
}

func Newads_rrasadministrationconnectionpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_rrasadministrationconnectionpoint, err error) {
	tmp, err := Newads_serviceadministrationpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_rrasadministrationconnectionpoint{
		ads_serviceadministrationpoint: tmp,
	}
	return
}

// SetDS_msRRASAttribute sets the value of DS_msRRASAttribute for the instance
func (instance *ads_rrasadministrationconnectionpoint) SetPropertyDS_msRRASAttribute(value []string) (err error) {
	return instance.SetProperty("DS_msRRASAttribute", (value))
}

// GetDS_msRRASAttribute gets the value of DS_msRRASAttribute for the instance
func (instance *ads_rrasadministrationconnectionpoint) GetPropertyDS_msRRASAttribute() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msRRASAttribute")
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
