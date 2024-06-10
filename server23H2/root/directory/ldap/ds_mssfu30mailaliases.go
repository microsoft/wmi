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

// ds_mssfu30mailaliases struct
type ds_mssfu30mailaliases struct {
	*ads_mssfu30mailaliases
}

func Newds_mssfu30mailaliasesEx1(instance *cim.WmiInstance) (newInstance *ds_mssfu30mailaliases, err error) {
	tmp, err := Newads_mssfu30mailaliasesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_mssfu30mailaliases{
		ads_mssfu30mailaliases: tmp,
	}
	return
}

func Newds_mssfu30mailaliasesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_mssfu30mailaliases, err error) {
	tmp, err := Newads_mssfu30mailaliasesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_mssfu30mailaliases{
		ads_mssfu30mailaliases: tmp,
	}
	return
}
