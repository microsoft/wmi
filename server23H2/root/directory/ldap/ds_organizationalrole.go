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

// ds_organizationalrole struct
type ds_organizationalrole struct {
	*ads_organizationalrole
}

func Newds_organizationalroleEx1(instance *cim.WmiInstance) (newInstance *ds_organizationalrole, err error) {
	tmp, err := Newads_organizationalroleEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_organizationalrole{
		ads_organizationalrole: tmp,
	}
	return
}

func Newds_organizationalroleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_organizationalrole, err error) {
	tmp, err := Newads_organizationalroleEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_organizationalrole{
		ads_organizationalrole: tmp,
	}
	return
}
