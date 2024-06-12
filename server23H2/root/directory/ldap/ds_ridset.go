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

// ds_ridset struct
type ds_ridset struct {
	*ads_ridset
}

func Newds_ridsetEx1(instance *cim.WmiInstance) (newInstance *ds_ridset, err error) {
	tmp, err := Newads_ridsetEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ridset{
		ads_ridset: tmp,
	}
	return
}

func Newds_ridsetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ridset, err error) {
	tmp, err := Newads_ridsetEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ridset{
		ads_ridset: tmp,
	}
	return
}
