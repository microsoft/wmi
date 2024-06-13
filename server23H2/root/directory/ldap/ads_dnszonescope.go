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

// ads_dnszonescope struct
type ads_dnszonescope struct {
	*ds_top

	//
	DS_dc string

	//
	DS_dNSProperty []Uint8Array

	//
	DS_managedBy string
}

func Newads_dnszonescopeEx1(instance *cim.WmiInstance) (newInstance *ads_dnszonescope, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_dnszonescope{
		ds_top: tmp,
	}
	return
}

func Newads_dnszonescopeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_dnszonescope, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_dnszonescope{
		ds_top: tmp,
	}
	return
}

// SetDS_dc sets the value of DS_dc for the instance
func (instance *ads_dnszonescope) SetPropertyDS_dc(value string) (err error) {
	return instance.SetProperty("DS_dc", (value))
}

// GetDS_dc gets the value of DS_dc for the instance
func (instance *ads_dnszonescope) GetPropertyDS_dc() (value string, err error) {
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
func (instance *ads_dnszonescope) SetPropertyDS_dNSProperty(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_dNSProperty", (value))
}

// GetDS_dNSProperty gets the value of DS_dNSProperty for the instance
func (instance *ads_dnszonescope) GetPropertyDS_dNSProperty() (value []Uint8Array, err error) {
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

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ads_dnszonescope) SetPropertyDS_managedBy(value string) (err error) {
	return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_dnszonescope) GetPropertyDS_managedBy() (value string, err error) {
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
