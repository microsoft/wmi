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

// ds_server struct
type ds_server struct { 
	*ads_server
}

	func Newds_serverEx1(instance *cim.WmiInstance) (newInstance *ds_server, err error) {tmp, err := Newads_serverEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_server {
	ads_server: tmp,
	}
	return
	}
	

	func Newds_serverEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_server, err error) {tmp, err := Newads_serverEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_server {
	ads_server: tmp,
	}
	return
	}
	

