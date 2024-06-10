// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_physicallocation struct
type ads_physicallocation struct {
	*ads_locality

	//
	DS_managedBy string
}

func Newads_physicallocationEx1(instance *cim.WmiInstance) (newInstance *ads_physicallocation, err error) {
	tmp, err := Newads_localityEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_physicallocation{
		ads_locality: tmp,
	}
	return
}

func Newads_physicallocationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_physicallocation, err error) {
	tmp, err := Newads_localityEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_physicallocation{
		ads_locality: tmp,
	}
	return
}

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ads_physicallocation) SetPropertyDS_managedBy(value string) (err error) {
	return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_physicallocation) GetPropertyDS_managedBy() (value string, err error) {
	retValue, err := instance.GetProperty("DS_managedBy")
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
