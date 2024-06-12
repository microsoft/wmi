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

// ds_user struct
type ds_user struct {
	*ads_user
}

func Newds_userEx1(instance *cim.WmiInstance) (newInstance *ds_user, err error) {
	tmp, err := Newads_userEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_user{
		ads_user: tmp,
	}
	return
}

func Newds_userEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_user, err error) {
	tmp, err := Newads_userEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_user{
		ads_user: tmp,
	}
	return
}
