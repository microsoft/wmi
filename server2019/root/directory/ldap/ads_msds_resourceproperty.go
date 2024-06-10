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

// ads_msds_resourceproperty struct
type ads_msds_resourceproperty struct {
	*ds_msds_claimtypepropertybase

	//
	DS_msDS_AppliesToResourceTypes []string

	//
	DS_msDS_IsUsedAsResourceSecurityAttribute bool

	//
	DS_msDS_ValueTypeReference string
}

func Newads_msds_resourcepropertyEx1(instance *cim.WmiInstance) (newInstance *ads_msds_resourceproperty, err error) {
	tmp, err := Newds_msds_claimtypepropertybaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_resourceproperty{
		ds_msds_claimtypepropertybase: tmp,
	}
	return
}

func Newads_msds_resourcepropertyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_resourceproperty, err error) {
	tmp, err := Newds_msds_claimtypepropertybaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_resourceproperty{
		ds_msds_claimtypepropertybase: tmp,
	}
	return
}

// SetDS_msDS_AppliesToResourceTypes sets the value of DS_msDS_AppliesToResourceTypes for the instance
func (instance *ads_msds_resourceproperty) SetPropertyDS_msDS_AppliesToResourceTypes(value []string) (err error) {
	return instance.SetProperty("DS_msDS_AppliesToResourceTypes", (value))
}

// GetDS_msDS_AppliesToResourceTypes gets the value of DS_msDS_AppliesToResourceTypes for the instance
func (instance *ads_msds_resourceproperty) GetPropertyDS_msDS_AppliesToResourceTypes() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AppliesToResourceTypes")
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

// SetDS_msDS_IsUsedAsResourceSecurityAttribute sets the value of DS_msDS_IsUsedAsResourceSecurityAttribute for the instance
func (instance *ads_msds_resourceproperty) SetPropertyDS_msDS_IsUsedAsResourceSecurityAttribute(value bool) (err error) {
	return instance.SetProperty("DS_msDS_IsUsedAsResourceSecurityAttribute", (value))
}

// GetDS_msDS_IsUsedAsResourceSecurityAttribute gets the value of DS_msDS_IsUsedAsResourceSecurityAttribute for the instance
func (instance *ads_msds_resourceproperty) GetPropertyDS_msDS_IsUsedAsResourceSecurityAttribute() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDS_IsUsedAsResourceSecurityAttribute")
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

// SetDS_msDS_ValueTypeReference sets the value of DS_msDS_ValueTypeReference for the instance
func (instance *ads_msds_resourceproperty) SetPropertyDS_msDS_ValueTypeReference(value string) (err error) {
	return instance.SetProperty("DS_msDS_ValueTypeReference", (value))
}

// GetDS_msDS_ValueTypeReference gets the value of DS_msDS_ValueTypeReference for the instance
func (instance *ads_msds_resourceproperty) GetPropertyDS_msDS_ValueTypeReference() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ValueTypeReference")
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
