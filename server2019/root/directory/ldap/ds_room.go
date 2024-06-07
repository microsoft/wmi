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

// ds_room struct
type ds_room struct { 
	*ads_room
}

	func Newds_roomEx1(instance *cim.WmiInstance) (newInstance *ds_room, err error) {tmp, err := Newads_roomEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_room {
	ads_room: tmp,
	}
	return
	}
	

	func Newds_roomEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_room, err error) {tmp, err := Newads_roomEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_room {
	ads_room: tmp,
	}
	return
	}
	

