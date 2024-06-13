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

// ads_rpcserver struct
type ads_rpcserver struct {
	*ds_rpcentry

	//
	DS_rpcNsCodeset []string

	//
	DS_rpcNsEntryFlags int32

	//
	DS_rpcNsObjectID []string
}

func Newads_rpcserverEx1(instance *cim.WmiInstance) (newInstance *ads_rpcserver, err error) {
	tmp, err := Newds_rpcentryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_rpcserver{
		ds_rpcentry: tmp,
	}
	return
}

func Newads_rpcserverEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_rpcserver, err error) {
	tmp, err := Newds_rpcentryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_rpcserver{
		ds_rpcentry: tmp,
	}
	return
}

// SetDS_rpcNsCodeset sets the value of DS_rpcNsCodeset for the instance
func (instance *ads_rpcserver) SetPropertyDS_rpcNsCodeset(value []string) (err error) {
	return instance.SetProperty("DS_rpcNsCodeset", (value))
}

// GetDS_rpcNsCodeset gets the value of DS_rpcNsCodeset for the instance
func (instance *ads_rpcserver) GetPropertyDS_rpcNsCodeset() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_rpcNsCodeset")
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

// SetDS_rpcNsEntryFlags sets the value of DS_rpcNsEntryFlags for the instance
func (instance *ads_rpcserver) SetPropertyDS_rpcNsEntryFlags(value int32) (err error) {
	return instance.SetProperty("DS_rpcNsEntryFlags", (value))
}

// GetDS_rpcNsEntryFlags gets the value of DS_rpcNsEntryFlags for the instance
func (instance *ads_rpcserver) GetPropertyDS_rpcNsEntryFlags() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_rpcNsEntryFlags")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetDS_rpcNsObjectID sets the value of DS_rpcNsObjectID for the instance
func (instance *ads_rpcserver) SetPropertyDS_rpcNsObjectID(value []string) (err error) {
	return instance.SetProperty("DS_rpcNsObjectID", (value))
}

// GetDS_rpcNsObjectID gets the value of DS_rpcNsObjectID for the instance
func (instance *ads_rpcserver) GetPropertyDS_rpcNsObjectID() (value []string, err error) {
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
