// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_ipsecisakmppolicy struct
type ds_ipsecisakmppolicy struct {
	*ads_ipsecisakmppolicy
}

func Newds_ipsecisakmppolicyEx1(instance *cim.WmiInstance) (newInstance *ds_ipsecisakmppolicy, err error) {
	tmp, err := Newads_ipsecisakmppolicyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ipsecisakmppolicy{
		ads_ipsecisakmppolicy: tmp,
	}
	return
}

func Newds_ipsecisakmppolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ipsecisakmppolicy, err error) {
	tmp, err := Newads_ipsecisakmppolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ipsecisakmppolicy{
		ads_ipsecisakmppolicy: tmp,
	}
	return
}
