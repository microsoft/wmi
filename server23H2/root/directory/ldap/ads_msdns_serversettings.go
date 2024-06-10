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

// ads_msdns_serversettings struct
type ads_msdns_serversettings struct {
	*ds_top

	//
	DS_msDNS_KeymasterZones []string
}

func Newads_msdns_serversettingsEx1(instance *cim.WmiInstance) (newInstance *ads_msdns_serversettings, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msdns_serversettings{
		ds_top: tmp,
	}
	return
}

func Newads_msdns_serversettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msdns_serversettings, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msdns_serversettings{
		ds_top: tmp,
	}
	return
}

// SetDS_msDNS_KeymasterZones sets the value of DS_msDNS_KeymasterZones for the instance
func (instance *ads_msdns_serversettings) SetPropertyDS_msDNS_KeymasterZones(value []string) (err error) {
	return instance.SetProperty("DS_msDNS_KeymasterZones", (value))
}

// GetDS_msDNS_KeymasterZones gets the value of DS_msDNS_KeymasterZones for the instance
func (instance *ads_msdns_serversettings) GetPropertyDS_msDNS_KeymasterZones() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_KeymasterZones")
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
