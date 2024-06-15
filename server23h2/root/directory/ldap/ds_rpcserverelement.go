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

// ds_rpcserverelement struct
type ds_rpcserverelement struct {
	*ads_rpcserverelement
}

func Newds_rpcserverelementEx1(instance *cim.WmiInstance) (newInstance *ds_rpcserverelement, err error) {
	tmp, err := Newads_rpcserverelementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_rpcserverelement{
		ads_rpcserverelement: tmp,
	}
	return
}

func Newds_rpcserverelementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_rpcserverelement, err error) {
	tmp, err := Newads_rpcserverelementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_rpcserverelement{
		ads_rpcserverelement: tmp,
	}
	return
}
