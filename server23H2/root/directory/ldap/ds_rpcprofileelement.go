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

// ds_rpcprofileelement struct
type ds_rpcprofileelement struct {
	*ads_rpcprofileelement
}

func Newds_rpcprofileelementEx1(instance *cim.WmiInstance) (newInstance *ds_rpcprofileelement, err error) {
	tmp, err := Newads_rpcprofileelementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_rpcprofileelement{
		ads_rpcprofileelement: tmp,
	}
	return
}

func Newds_rpcprofileelementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_rpcprofileelement, err error) {
	tmp, err := Newads_rpcprofileelementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_rpcprofileelement{
		ads_rpcprofileelement: tmp,
	}
	return
}
