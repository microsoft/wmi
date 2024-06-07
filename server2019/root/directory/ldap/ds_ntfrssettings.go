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

// ds_ntfrssettings struct
type ds_ntfrssettings struct { 
	*ads_ntfrssettings
}

	func Newds_ntfrssettingsEx1(instance *cim.WmiInstance) (newInstance *ds_ntfrssettings, err error) {tmp, err := Newads_ntfrssettingsEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_ntfrssettings {
	ads_ntfrssettings: tmp,
	}
	return
	}
	

	func Newds_ntfrssettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_ntfrssettings, err error) {tmp, err := Newads_ntfrssettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_ntfrssettings {
	ads_ntfrssettings: tmp,
	}
	return
	}
	

