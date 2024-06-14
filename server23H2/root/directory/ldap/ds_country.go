// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_country struct
type ds_country struct {
	*ads_country
}

func Newds_countryEx1(instance *cim.WmiInstance) (newInstance *ds_country, err error) {
	tmp, err := Newads_countryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_country{
		ads_country: tmp,
	}
	return
}

func Newds_countryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_country, err error) {
	tmp, err := Newads_countryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_country{
		ads_country: tmp,
	}
	return
}
