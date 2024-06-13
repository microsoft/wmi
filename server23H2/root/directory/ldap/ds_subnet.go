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

// ds_subnet struct
type ds_subnet struct {
	*ads_subnet
}

func Newds_subnetEx1(instance *cim.WmiInstance) (newInstance *ds_subnet, err error) {
	tmp, err := Newads_subnetEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_subnet{
		ads_subnet: tmp,
	}
	return
}

func Newds_subnetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_subnet, err error) {
	tmp, err := Newads_subnetEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_subnet{
		ads_subnet: tmp,
	}
	return
}
