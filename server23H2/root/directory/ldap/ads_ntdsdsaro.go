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
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ads_ntdsdsaro struct
type ads_ntdsdsaro struct {
	*ads_ntdsdsa
}

func Newads_ntdsdsaroEx1(instance *cim.WmiInstance) (newInstance *ads_ntdsdsaro, err error) {
	tmp, err := Newads_ntdsdsaEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ntdsdsaro{
		ads_ntdsdsa: tmp,
	}
	return
}

func Newads_ntdsdsaroEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ntdsdsaro, err error) {
	tmp, err := Newads_ntdsdsaEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ntdsdsaro{
		ads_ntdsdsa: tmp,
	}
	return
}
