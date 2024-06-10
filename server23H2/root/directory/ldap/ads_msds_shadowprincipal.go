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

// ads_msds_shadowprincipal struct
type ads_msds_shadowprincipal struct {
	*ds_top

	//
	DS_member []string

	//
	DS_msDS_ShadowPrincipalSid Uint8Array
}

func Newads_msds_shadowprincipalEx1(instance *cim.WmiInstance) (newInstance *ads_msds_shadowprincipal, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_shadowprincipal{
		ds_top: tmp,
	}
	return
}

func Newads_msds_shadowprincipalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_shadowprincipal, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_shadowprincipal{
		ds_top: tmp,
	}
	return
}

// SetDS_member sets the value of DS_member for the instance
func (instance *ads_msds_shadowprincipal) SetPropertyDS_member(value []string) (err error) {
	return instance.SetProperty("DS_member", (value))
}

// GetDS_member gets the value of DS_member for the instance
func (instance *ads_msds_shadowprincipal) GetPropertyDS_member() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_member")
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

// SetDS_msDS_ShadowPrincipalSid sets the value of DS_msDS_ShadowPrincipalSid for the instance
func (instance *ads_msds_shadowprincipal) SetPropertyDS_msDS_ShadowPrincipalSid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_ShadowPrincipalSid", (value))
}

// GetDS_msDS_ShadowPrincipalSid gets the value of DS_msDS_ShadowPrincipalSid for the instance
func (instance *ads_msds_shadowprincipal) GetPropertyDS_msDS_ShadowPrincipalSid() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ShadowPrincipalSid")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}
