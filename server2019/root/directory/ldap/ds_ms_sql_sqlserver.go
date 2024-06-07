// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_ms_sql_sqlserver struct
type ds_ms_sql_sqlserver struct { 
	*ads_ms_sql_sqlserver
}

	func Newds_ms_sql_sqlserverEx1(instance *cim.WmiInstance) (newInstance *ds_ms_sql_sqlserver, err error) {tmp, err := Newads_ms_sql_sqlserverEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_ms_sql_sqlserver {
	ads_ms_sql_sqlserver: tmp,
	}
	return
	}
	

	func Newds_ms_sql_sqlserverEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_ms_sql_sqlserver, err error) {tmp, err := Newads_ms_sql_sqlserverEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_ms_sql_sqlserver {
	ads_ms_sql_sqlserver: tmp,
	}
	return
	}
	

