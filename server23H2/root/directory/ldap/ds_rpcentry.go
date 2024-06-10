// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_rpcentry struct
type ds_rpcentry struct {
	*ds_connectionpoint
}

func Newds_rpcentryEx1(instance *cim.WmiInstance) (newInstance *ds_rpcentry, err error) {
	tmp, err := Newds_connectionpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_rpcentry{
		ds_connectionpoint: tmp,
	}
	return
}

func Newds_rpcentryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_rpcentry, err error) {
	tmp, err := Newds_connectionpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_rpcentry{
		ds_connectionpoint: tmp,
	}
	return
}
