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

// ds_linktrackomtentry struct
type ds_linktrackomtentry struct {
	*ads_linktrackomtentry
}

func Newds_linktrackomtentryEx1(instance *cim.WmiInstance) (newInstance *ds_linktrackomtentry, err error) {
	tmp, err := Newads_linktrackomtentryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_linktrackomtentry{
		ads_linktrackomtentry: tmp,
	}
	return
}

func Newds_linktrackomtentryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_linktrackomtentry, err error) {
	tmp, err := Newads_linktrackomtentryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_linktrackomtentry{
		ads_linktrackomtentry: tmp,
	}
	return
}
