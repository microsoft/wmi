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

// ads_rpcgroup struct
type ads_rpcgroup struct {
	*ds_rpcentry

	//
	DS_rpcNsGroup []string

	//
	DS_rpcNsObjectID []string
}

func Newads_rpcgroupEx1(instance *cim.WmiInstance) (newInstance *ads_rpcgroup, err error) {
	tmp, err := Newds_rpcentryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_rpcgroup{
		ds_rpcentry: tmp,
	}
	return
}

func Newads_rpcgroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_rpcgroup, err error) {
	tmp, err := Newds_rpcentryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_rpcgroup{
		ds_rpcentry: tmp,
	}
	return
}

// SetDS_rpcNsGroup sets the value of DS_rpcNsGroup for the instance
func (instance *ads_rpcgroup) SetPropertyDS_rpcNsGroup(value []string) (err error) {
	return instance.SetProperty("DS_rpcNsGroup", (value))
}

// GetDS_rpcNsGroup gets the value of DS_rpcNsGroup for the instance
func (instance *ads_rpcgroup) GetPropertyDS_rpcNsGroup() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_rpcNsGroup")
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

// SetDS_rpcNsObjectID sets the value of DS_rpcNsObjectID for the instance
func (instance *ads_rpcgroup) SetPropertyDS_rpcNsObjectID(value []string) (err error) {
	return instance.SetProperty("DS_rpcNsObjectID", (value))
}

// GetDS_rpcNsObjectID gets the value of DS_rpcNsObjectID for the instance
func (instance *ads_rpcgroup) GetPropertyDS_rpcNsObjectID() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_rpcNsObjectID")
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
