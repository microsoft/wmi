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

// ds_ntdsdsa struct
type ds_ntdsdsa struct { 
	*ads_ntdsdsa
}

	func Newds_ntdsdsaEx1(instance *cim.WmiInstance) (newInstance *ds_ntdsdsa, err error) {tmp, err := Newads_ntdsdsaEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_ntdsdsa {
	ads_ntdsdsa: tmp,
	}
	return
	}
	

	func Newds_ntdsdsaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_ntdsdsa, err error) {tmp, err := Newads_ntdsdsaEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_ntdsdsa {
	ads_ntdsdsa: tmp,
	}
	return
	}
	

