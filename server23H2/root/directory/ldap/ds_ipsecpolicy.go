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

// ds_ipsecpolicy struct
type ds_ipsecpolicy struct {
	*ads_ipsecpolicy
}

func Newds_ipsecpolicyEx1(instance *cim.WmiInstance) (newInstance *ds_ipsecpolicy, err error) {
	tmp, err := Newads_ipsecpolicyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ipsecpolicy{
		ads_ipsecpolicy: tmp,
	}
	return
}

func Newds_ipsecpolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ipsecpolicy, err error) {
	tmp, err := Newads_ipsecpolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ipsecpolicy{
		ads_ipsecpolicy: tmp,
	}
	return
}
