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
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_rrasadministrationdictionary struct
type ds_rrasadministrationdictionary struct {
	*ads_rrasadministrationdictionary
}

func Newds_rrasadministrationdictionaryEx1(instance *cim.WmiInstance) (newInstance *ds_rrasadministrationdictionary, err error) {
	tmp, err := Newads_rrasadministrationdictionaryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_rrasadministrationdictionary{
		ads_rrasadministrationdictionary: tmp,
	}
	return
}

func Newds_rrasadministrationdictionaryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_rrasadministrationdictionary, err error) {
	tmp, err := Newads_rrasadministrationdictionaryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_rrasadministrationdictionary{
		ads_rrasadministrationdictionary: tmp,
	}
	return
}
