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

// ds_ntdsdsaro struct
type ds_ntdsdsaro struct {
	*ads_ntdsdsaro
}

func Newds_ntdsdsaroEx1(instance *cim.WmiInstance) (newInstance *ds_ntdsdsaro, err error) {
	tmp, err := Newads_ntdsdsaroEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ntdsdsaro{
		ads_ntdsdsaro: tmp,
	}
	return
}

func Newds_ntdsdsaroEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ntdsdsaro, err error) {
	tmp, err := Newads_ntdsdsaroEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ntdsdsaro{
		ads_ntdsdsaro: tmp,
	}
	return
}
