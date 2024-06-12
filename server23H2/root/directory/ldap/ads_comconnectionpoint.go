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

// ads_comconnectionpoint struct
type ads_comconnectionpoint struct {
	*ds_connectionpoint

	//
	DS_marshalledInterface []Uint8Array

	//
	DS_moniker []Uint8Array

	//
	DS_monikerDisplayName []string
}

func Newads_comconnectionpointEx1(instance *cim.WmiInstance) (newInstance *ads_comconnectionpoint, err error) {
	tmp, err := Newds_connectionpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_comconnectionpoint{
		ds_connectionpoint: tmp,
	}
	return
}

func Newads_comconnectionpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_comconnectionpoint, err error) {
	tmp, err := Newds_connectionpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_comconnectionpoint{
		ds_connectionpoint: tmp,
	}
	return
}

// SetDS_marshalledInterface sets the value of DS_marshalledInterface for the instance
func (instance *ads_comconnectionpoint) SetPropertyDS_marshalledInterface(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_marshalledInterface", (value))
}

// GetDS_marshalledInterface gets the value of DS_marshalledInterface for the instance
func (instance *ads_comconnectionpoint) GetPropertyDS_marshalledInterface() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_marshalledInterface")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Uint8Array)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Uint8Array(valuetmp))
	}

	return
}

// SetDS_moniker sets the value of DS_moniker for the instance
func (instance *ads_comconnectionpoint) SetPropertyDS_moniker(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_moniker", (value))
}

// GetDS_moniker gets the value of DS_moniker for the instance
func (instance *ads_comconnectionpoint) GetPropertyDS_moniker() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_moniker")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Uint8Array)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Uint8Array(valuetmp))
	}

	return
}

// SetDS_monikerDisplayName sets the value of DS_monikerDisplayName for the instance
func (instance *ads_comconnectionpoint) SetPropertyDS_monikerDisplayName(value []string) (err error) {
	return instance.SetProperty("DS_monikerDisplayName", (value))
}

// GetDS_monikerDisplayName gets the value of DS_monikerDisplayName for the instance
func (instance *ads_comconnectionpoint) GetPropertyDS_monikerDisplayName() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_monikerDisplayName")
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
