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

// ds_organizationalunit struct
type ds_organizationalunit struct {
	*ads_organizationalunit
}

func Newds_organizationalunitEx1(instance *cim.WmiInstance) (newInstance *ds_organizationalunit, err error) {
	tmp, err := Newads_organizationalunitEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_organizationalunit{
		ads_organizationalunit: tmp,
	}
	return
}

func Newds_organizationalunitEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_organizationalunit, err error) {
	tmp, err := Newads_organizationalunitEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_organizationalunit{
		ads_organizationalunit: tmp,
	}
	return
}
