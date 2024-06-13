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

// ds_rpcgroup struct
type ds_rpcgroup struct {
	*ads_rpcgroup
}

func Newds_rpcgroupEx1(instance *cim.WmiInstance) (newInstance *ds_rpcgroup, err error) {
	tmp, err := Newads_rpcgroupEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_rpcgroup{
		ads_rpcgroup: tmp,
	}
	return
}

func Newds_rpcgroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_rpcgroup, err error) {
	tmp, err := Newads_rpcgroupEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_rpcgroup{
		ads_rpcgroup: tmp,
	}
	return
}
