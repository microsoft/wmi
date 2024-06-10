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

// ds_friendlycountry struct
type ds_friendlycountry struct {
	*ads_friendlycountry
}

func Newds_friendlycountryEx1(instance *cim.WmiInstance) (newInstance *ds_friendlycountry, err error) {
	tmp, err := Newads_friendlycountryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_friendlycountry{
		ads_friendlycountry: tmp,
	}
	return
}

func Newds_friendlycountryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_friendlycountry, err error) {
	tmp, err := Newads_friendlycountryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_friendlycountry{
		ads_friendlycountry: tmp,
	}
	return
}
