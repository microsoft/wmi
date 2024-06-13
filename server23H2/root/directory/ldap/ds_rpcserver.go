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

// ds_rpcserver struct
type ds_rpcserver struct {
	*ads_rpcserver
}

func Newds_rpcserverEx1(instance *cim.WmiInstance) (newInstance *ds_rpcserver, err error) {
	tmp, err := Newads_rpcserverEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_rpcserver{
		ads_rpcserver: tmp,
	}
	return
}

func Newds_rpcserverEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_rpcserver, err error) {
	tmp, err := Newads_rpcserverEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_rpcserver{
		ads_rpcserver: tmp,
	}
	return
}
