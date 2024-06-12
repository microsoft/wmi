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

// ads_msds_resourcepropertylist struct
type ads_msds_resourcepropertylist struct {
	*ds_top

	//
	DS_msDS_MembersOfResourcePropertyList []string
}

func Newads_msds_resourcepropertylistEx1(instance *cim.WmiInstance) (newInstance *ads_msds_resourcepropertylist, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_resourcepropertylist{
		ds_top: tmp,
	}
	return
}

func Newads_msds_resourcepropertylistEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_resourcepropertylist, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_resourcepropertylist{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_MembersOfResourcePropertyList sets the value of DS_msDS_MembersOfResourcePropertyList for the instance
func (instance *ads_msds_resourcepropertylist) SetPropertyDS_msDS_MembersOfResourcePropertyList(value []string) (err error) {
	return instance.SetProperty("DS_msDS_MembersOfResourcePropertyList", (value))
}

// GetDS_msDS_MembersOfResourcePropertyList gets the value of DS_msDS_MembersOfResourcePropertyList for the instance
func (instance *ads_msds_resourcepropertylist) GetPropertyDS_msDS_MembersOfResourcePropertyList() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_MembersOfResourcePropertyList")
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
