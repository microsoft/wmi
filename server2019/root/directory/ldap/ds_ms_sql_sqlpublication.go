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

// ds_ms_sql_sqlpublication struct
type ds_ms_sql_sqlpublication struct {
	*ads_ms_sql_sqlpublication
}

func Newds_ms_sql_sqlpublicationEx1(instance *cim.WmiInstance) (newInstance *ds_ms_sql_sqlpublication, err error) {
	tmp, err := Newads_ms_sql_sqlpublicationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ms_sql_sqlpublication{
		ads_ms_sql_sqlpublication: tmp,
	}
	return
}

func Newds_ms_sql_sqlpublicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ms_sql_sqlpublication, err error) {
	tmp, err := Newads_ms_sql_sqlpublicationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ms_sql_sqlpublication{
		ads_ms_sql_sqlpublication: tmp,
	}
	return
}
