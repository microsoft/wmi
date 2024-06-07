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

// ds_dsuisettings struct
type ds_dsuisettings struct { 
	*ads_dsuisettings
}

	func Newds_dsuisettingsEx1(instance *cim.WmiInstance) (newInstance *ds_dsuisettings, err error) {tmp, err := Newads_dsuisettingsEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_dsuisettings {
	ads_dsuisettings: tmp,
	}
	return
	}
	

	func Newds_dsuisettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_dsuisettings, err error) {tmp, err := Newads_dsuisettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_dsuisettings {
	ads_dsuisettings: tmp,
	}
	return
	}
	

