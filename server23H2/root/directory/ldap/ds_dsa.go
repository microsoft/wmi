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

// ds_dsa struct
type ds_dsa struct {
	*ads_dsa
}

func Newds_dsaEx1(instance *cim.WmiInstance) (newInstance *ds_dsa, err error) {
	tmp, err := Newads_dsaEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_dsa{
		ads_dsa: tmp,
	}
	return
}

func Newds_dsaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_dsa, err error) {
	tmp, err := Newads_dsaEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_dsa{
		ads_dsa: tmp,
	}
	return
}
