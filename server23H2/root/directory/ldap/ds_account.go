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
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_account struct
type ds_account struct {
	*ads_account
}

func Newds_accountEx1(instance *cim.WmiInstance) (newInstance *ds_account, err error) {
	tmp, err := Newads_accountEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_account{
		ads_account: tmp,
	}
	return
}

func Newds_accountEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_account, err error) {
	tmp, err := Newads_accountEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_account{
		ads_account: tmp,
	}
	return
}
