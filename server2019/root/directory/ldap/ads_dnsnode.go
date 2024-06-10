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

// ads_dnsnode struct
type ads_dnsnode struct {
	*ds_top

	//
	DS_dc string

	//
	DS_dNSProperty []Uint8Array

	//
	DS_dnsRecord []Uint8Array

	//
	DS_dNSTombstoned bool
}

func Newads_dnsnodeEx1(instance *cim.WmiInstance) (newInstance *ads_dnsnode, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_dnsnode{
		ds_top: tmp,
	}
	return
}

func Newads_dnsnodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_dnsnode, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_dnsnode{
		ds_top: tmp,
	}
	return
}

// SetDS_dc sets the value of DS_dc for the instance
func (instance *ads_dnsnode) SetPropertyDS_dc(value string) (err error) {
	return instance.SetProperty("DS_dc", (value))
}

// GetDS_dc gets the value of DS_dc for the instance
func (instance *ads_dnsnode) GetPropertyDS_dc() (value string, err error) {
	retValue, err := instance.GetProperty("DS_dc")
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

// SetDS_dNSProperty sets the value of DS_dNSProperty for the instance
func (instance *ads_dnsnode) SetPropertyDS_dNSProperty(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_dNSProperty", (value))
}

// GetDS_dNSProperty gets the value of DS_dNSProperty for the instance
func (instance *ads_dnsnode) GetPropertyDS_dNSProperty() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_dNSProperty")
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

// SetDS_dnsRecord sets the value of DS_dnsRecord for the instance
func (instance *ads_dnsnode) SetPropertyDS_dnsRecord(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_dnsRecord", (value))
}

// GetDS_dnsRecord gets the value of DS_dnsRecord for the instance
func (instance *ads_dnsnode) GetPropertyDS_dnsRecord() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_dnsRecord")
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

// SetDS_dNSTombstoned sets the value of DS_dNSTombstoned for the instance
func (instance *ads_dnsnode) SetPropertyDS_dNSTombstoned(value bool) (err error) {
	return instance.SetProperty("DS_dNSTombstoned", (value))
}

// GetDS_dNSTombstoned gets the value of DS_dNSTombstoned for the instance
func (instance *ads_dnsnode) GetPropertyDS_dNSTombstoned() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_dNSTombstoned")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
