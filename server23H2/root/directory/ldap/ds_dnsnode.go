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

// ds_dnsnode struct
type ds_dnsnode struct {
	*ads_dnsnode
}

func Newds_dnsnodeEx1(instance *cim.WmiInstance) (newInstance *ds_dnsnode, err error) {
	tmp, err := Newads_dnsnodeEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_dnsnode{
		ads_dnsnode: tmp,
	}
	return
}

func Newds_dnsnodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_dnsnode, err error) {
	tmp, err := Newads_dnsnodeEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_dnsnode{
		ads_dnsnode: tmp,
	}
	return
}
