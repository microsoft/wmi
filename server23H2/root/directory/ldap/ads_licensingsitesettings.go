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

// ads_licensingsitesettings struct
type ads_licensingsitesettings struct {
	*ds_applicationsitesettings

	//
	DS_siteServer []string
}

func Newads_licensingsitesettingsEx1(instance *cim.WmiInstance) (newInstance *ads_licensingsitesettings, err error) {
	tmp, err := Newds_applicationsitesettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_licensingsitesettings{
		ds_applicationsitesettings: tmp,
	}
	return
}

func Newads_licensingsitesettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_licensingsitesettings, err error) {
	tmp, err := Newds_applicationsitesettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_licensingsitesettings{
		ds_applicationsitesettings: tmp,
	}
	return
}

// SetDS_siteServer sets the value of DS_siteServer for the instance
func (instance *ads_licensingsitesettings) SetPropertyDS_siteServer(value []string) (err error) {
	return instance.SetProperty("DS_siteServer", (value))
}

// GetDS_siteServer gets the value of DS_siteServer for the instance
func (instance *ads_licensingsitesettings) GetPropertyDS_siteServer() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_siteServer")
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
