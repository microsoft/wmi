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

// ds_physicallocation struct
type ds_physicallocation struct {
	*ads_physicallocation
}

func Newds_physicallocationEx1(instance *cim.WmiInstance) (newInstance *ds_physicallocation, err error) {
	tmp, err := Newads_physicallocationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_physicallocation{
		ads_physicallocation: tmp,
	}
	return
}

func Newds_physicallocationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_physicallocation, err error) {
	tmp, err := Newads_physicallocationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_physicallocation{
		ads_physicallocation: tmp,
	}
	return
}
