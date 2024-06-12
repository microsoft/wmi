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

// ds_domainpolicy struct
type ds_domainpolicy struct {
	*ads_domainpolicy
}

func Newds_domainpolicyEx1(instance *cim.WmiInstance) (newInstance *ds_domainpolicy, err error) {
	tmp, err := Newads_domainpolicyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_domainpolicy{
		ads_domainpolicy: tmp,
	}
	return
}

func Newds_domainpolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_domainpolicy, err error) {
	tmp, err := Newads_domainpolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_domainpolicy{
		ads_domainpolicy: tmp,
	}
	return
}
