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

// ads_friendlycountry struct
type ads_friendlycountry struct {
	*ads_country
}

func Newads_friendlycountryEx1(instance *cim.WmiInstance) (newInstance *ads_friendlycountry, err error) {
	tmp, err := Newads_countryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_friendlycountry{
		ads_country: tmp,
	}
	return
}

func Newads_friendlycountryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_friendlycountry, err error) {
	tmp, err := Newads_countryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_friendlycountry{
		ads_country: tmp,
	}
	return
}
