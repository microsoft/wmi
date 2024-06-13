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

// ds_ms_sql_olapdatabase struct
type ds_ms_sql_olapdatabase struct {
	*ads_ms_sql_olapdatabase
}

func Newds_ms_sql_olapdatabaseEx1(instance *cim.WmiInstance) (newInstance *ds_ms_sql_olapdatabase, err error) {
	tmp, err := Newads_ms_sql_olapdatabaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ms_sql_olapdatabase{
		ads_ms_sql_olapdatabase: tmp,
	}
	return
}

func Newds_ms_sql_olapdatabaseEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ms_sql_olapdatabase, err error) {
	tmp, err := Newads_ms_sql_olapdatabaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ms_sql_olapdatabase{
		ads_ms_sql_olapdatabase: tmp,
	}
	return
}
