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

// ds_sitelink struct
type ds_sitelink struct { 
	*ads_sitelink
}

	func Newds_sitelinkEx1(instance *cim.WmiInstance) (newInstance *ds_sitelink, err error) {tmp, err := Newads_sitelinkEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_sitelink {
	ads_sitelink: tmp,
	}
	return
	}
	

	func Newds_sitelinkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_sitelink, err error) {tmp, err := Newads_sitelinkEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_sitelink {
	ads_sitelink: tmp,
	}
	return
	}
	

