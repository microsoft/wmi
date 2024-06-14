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

// ads_ipsecisakmppolicy struct
type ads_ipsecisakmppolicy struct {
	*ds_ipsecbase
}

func Newads_ipsecisakmppolicyEx1(instance *cim.WmiInstance) (newInstance *ads_ipsecisakmppolicy, err error) {
	tmp, err := Newds_ipsecbaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ipsecisakmppolicy{
		ds_ipsecbase: tmp,
	}
	return
}

func Newads_ipsecisakmppolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ipsecisakmppolicy, err error) {
	tmp, err := Newds_ipsecbaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ipsecisakmppolicy{
		ds_ipsecbase: tmp,
	}
	return
}
