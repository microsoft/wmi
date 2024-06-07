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

// ds_serviceclass struct
type ds_serviceclass struct { 
	*ads_serviceclass
}

	func Newds_serviceclassEx1(instance *cim.WmiInstance) (newInstance *ds_serviceclass, err error) {tmp, err := Newads_serviceclassEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_serviceclass {
	ads_serviceclass: tmp,
	}
	return
	}
	

	func Newds_serviceclassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_serviceclass, err error) {tmp, err := Newads_serviceclassEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_serviceclass {
	ads_serviceclass: tmp,
	}
	return
	}
	

