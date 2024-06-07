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

// ds_ntfrsmember struct
type ds_ntfrsmember struct { 
	*ads_ntfrsmember
}

	func Newds_ntfrsmemberEx1(instance *cim.WmiInstance) (newInstance *ds_ntfrsmember, err error) {tmp, err := Newads_ntfrsmemberEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_ntfrsmember {
	ads_ntfrsmember: tmp,
	}
	return
	}
	

	func Newds_ntfrsmemberEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_ntfrsmember, err error) {tmp, err := Newads_ntfrsmemberEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_ntfrsmember {
	ads_ntfrsmember: tmp,
	}
	return
	}
	

