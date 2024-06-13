// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ds_ieee802device struct
type ds_ieee802device struct {
	*ds_top

	//
	DS_macAddress []string
}

func Newds_ieee802deviceEx1(instance *cim.WmiInstance) (newInstance *ds_ieee802device, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ieee802device{
		ds_top: tmp,
	}
	return
}

func Newds_ieee802deviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ieee802device, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ieee802device{
		ds_top: tmp,
	}
	return
}

// SetDS_macAddress sets the value of DS_macAddress for the instance
func (instance *ds_ieee802device) SetPropertyDS_macAddress(value []string) (err error) {
	return instance.SetProperty("DS_macAddress", (value))
}

// GetDS_macAddress gets the value of DS_macAddress for the instance
func (instance *ds_ieee802device) GetPropertyDS_macAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_macAddress")
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
