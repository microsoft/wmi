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

// ds_msprint_connectionpolicy struct
type ds_msprint_connectionpolicy struct {
	*ads_msprint_connectionpolicy
}

func Newds_msprint_connectionpolicyEx1(instance *cim.WmiInstance) (newInstance *ds_msprint_connectionpolicy, err error) {
	tmp, err := Newads_msprint_connectionpolicyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msprint_connectionpolicy{
		ads_msprint_connectionpolicy: tmp,
	}
	return
}

func Newds_msprint_connectionpolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msprint_connectionpolicy, err error) {
	tmp, err := Newads_msprint_connectionpolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msprint_connectionpolicy{
		ads_msprint_connectionpolicy: tmp,
	}
	return
}
