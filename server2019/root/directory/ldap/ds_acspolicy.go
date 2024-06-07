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

// ds_acspolicy struct
type ds_acspolicy struct { 
	*ads_acspolicy
}

	func Newds_acspolicyEx1(instance *cim.WmiInstance) (newInstance *ds_acspolicy, err error) {tmp, err := Newads_acspolicyEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_acspolicy {
	ads_acspolicy: tmp,
	}
	return
	}
	

	func Newds_acspolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_acspolicy, err error) {tmp, err := Newads_acspolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_acspolicy {
	ads_acspolicy: tmp,
	}
	return
	}
	

