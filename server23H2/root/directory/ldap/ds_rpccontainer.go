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

// ds_rpccontainer struct
type ds_rpccontainer struct {
	*ads_rpccontainer
}

func Newds_rpccontainerEx1(instance *cim.WmiInstance) (newInstance *ds_rpccontainer, err error) {
	tmp, err := Newads_rpccontainerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_rpccontainer{
		ads_rpccontainer: tmp,
	}
	return
}

func Newds_rpccontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_rpccontainer, err error) {
	tmp, err := Newads_rpccontainerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_rpccontainer{
		ads_rpccontainer: tmp,
	}
	return
}
