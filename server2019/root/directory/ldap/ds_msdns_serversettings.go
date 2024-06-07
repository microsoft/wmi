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

// ds_msdns_serversettings struct
type ds_msdns_serversettings struct { 
	*ads_msdns_serversettings
}

	func Newds_msdns_serversettingsEx1(instance *cim.WmiInstance) (newInstance *ds_msdns_serversettings, err error) {tmp, err := Newads_msdns_serversettingsEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_msdns_serversettings {
	ads_msdns_serversettings: tmp,
	}
	return
	}
	

	func Newds_msdns_serversettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_msdns_serversettings, err error) {tmp, err := Newads_msdns_serversettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_msdns_serversettings {
	ads_msdns_serversettings: tmp,
	}
	return
	}
	

