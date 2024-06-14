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

// ds_foreignsecurityprincipal struct
type ds_foreignsecurityprincipal struct {
	*ads_foreignsecurityprincipal
}

func Newds_foreignsecurityprincipalEx1(instance *cim.WmiInstance) (newInstance *ds_foreignsecurityprincipal, err error) {
	tmp, err := Newads_foreignsecurityprincipalEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_foreignsecurityprincipal{
		ads_foreignsecurityprincipal: tmp,
	}
	return
}

func Newds_foreignsecurityprincipalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_foreignsecurityprincipal, err error) {
	tmp, err := Newads_foreignsecurityprincipalEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_foreignsecurityprincipal{
		ads_foreignsecurityprincipal: tmp,
	}
	return
}
