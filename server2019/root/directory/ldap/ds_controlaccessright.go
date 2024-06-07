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

// ds_controlaccessright struct
type ds_controlaccessright struct { 
	*ads_controlaccessright
}

	func Newds_controlaccessrightEx1(instance *cim.WmiInstance) (newInstance *ds_controlaccessright, err error) {tmp, err := Newads_controlaccessrightEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_controlaccessright {
	ads_controlaccessright: tmp,
	}
	return
	}
	

	func Newds_controlaccessrightEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_controlaccessright, err error) {tmp, err := Newads_controlaccessrightEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_controlaccessright {
	ads_controlaccessright: tmp,
	}
	return
	}
	

