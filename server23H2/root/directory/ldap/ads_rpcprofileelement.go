// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_rpcprofileelement struct
type ads_rpcprofileelement struct {
	*ds_rpcentry

	//
	DS_rpcNsAnnotation string

	//
	DS_rpcNsInterfaceID string

	//
	DS_rpcNsPriority []int32

	//
	DS_rpcNsProfileEntry string
}

func Newads_rpcprofileelementEx1(instance *cim.WmiInstance) (newInstance *ads_rpcprofileelement, err error) {
	tmp, err := Newds_rpcentryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_rpcprofileelement{
		ds_rpcentry: tmp,
	}
	return
}

func Newads_rpcprofileelementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_rpcprofileelement, err error) {
	tmp, err := Newds_rpcentryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_rpcprofileelement{
		ds_rpcentry: tmp,
	}
	return
}

// SetDS_rpcNsAnnotation sets the value of DS_rpcNsAnnotation for the instance
func (instance *ads_rpcprofileelement) SetPropertyDS_rpcNsAnnotation(value string) (err error) {
	return instance.SetProperty("DS_rpcNsAnnotation", (value))
}

// GetDS_rpcNsAnnotation gets the value of DS_rpcNsAnnotation for the instance
func (instance *ads_rpcprofileelement) GetPropertyDS_rpcNsAnnotation() (value string, err error) {
	retValue, err := instance.GetProperty("DS_rpcNsAnnotation")
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

// SetDS_rpcNsInterfaceID sets the value of DS_rpcNsInterfaceID for the instance
func (instance *ads_rpcprofileelement) SetPropertyDS_rpcNsInterfaceID(value string) (err error) {
	return instance.SetProperty("DS_rpcNsInterfaceID", (value))
}

// GetDS_rpcNsInterfaceID gets the value of DS_rpcNsInterfaceID for the instance
func (instance *ads_rpcprofileelement) GetPropertyDS_rpcNsInterfaceID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_rpcNsInterfaceID")
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

// SetDS_rpcNsPriority sets the value of DS_rpcNsPriority for the instance
func (instance *ads_rpcprofileelement) SetPropertyDS_rpcNsPriority(value []int32) (err error) {
	return instance.SetProperty("DS_rpcNsPriority", (value))
}

// GetDS_rpcNsPriority gets the value of DS_rpcNsPriority for the instance
func (instance *ads_rpcprofileelement) GetPropertyDS_rpcNsPriority() (value []int32, err error) {
	retValue, err := instance.GetProperty("DS_rpcNsPriority")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, int32(valuetmp))
	}

	return
}

// SetDS_rpcNsProfileEntry sets the value of DS_rpcNsProfileEntry for the instance
func (instance *ads_rpcprofileelement) SetPropertyDS_rpcNsProfileEntry(value string) (err error) {
	return instance.SetProperty("DS_rpcNsProfileEntry", (value))
}

// GetDS_rpcNsProfileEntry gets the value of DS_rpcNsProfileEntry for the instance
func (instance *ads_rpcprofileelement) GetPropertyDS_rpcNsProfileEntry() (value string, err error) {
	retValue, err := instance.GetProperty("DS_rpcNsProfileEntry")
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
