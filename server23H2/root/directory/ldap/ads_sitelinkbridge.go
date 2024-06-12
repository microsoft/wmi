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

// ads_sitelinkbridge struct
type ads_sitelinkbridge struct {
	*ds_top

	//
	DS_siteLinkList []string
}

func Newads_sitelinkbridgeEx1(instance *cim.WmiInstance) (newInstance *ads_sitelinkbridge, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_sitelinkbridge{
		ds_top: tmp,
	}
	return
}

func Newads_sitelinkbridgeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_sitelinkbridge, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_sitelinkbridge{
		ds_top: tmp,
	}
	return
}

// SetDS_siteLinkList sets the value of DS_siteLinkList for the instance
func (instance *ads_sitelinkbridge) SetPropertyDS_siteLinkList(value []string) (err error) {
	return instance.SetProperty("DS_siteLinkList", (value))
}

// GetDS_siteLinkList gets the value of DS_siteLinkList for the instance
func (instance *ads_sitelinkbridge) GetPropertyDS_siteLinkList() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_siteLinkList")
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
