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

// ds_leaf struct
type ds_leaf struct { 
	*ds_top
}

	func Newds_leafEx1(instance *cim.WmiInstance) (newInstance *ds_leaf, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_leaf {
	ds_top: tmp,
	}
	return
	}
	

	func Newds_leafEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_leaf, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_leaf {
	ds_top: tmp,
	}
	return
	}
	

