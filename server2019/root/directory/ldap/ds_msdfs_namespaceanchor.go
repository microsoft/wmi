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

// ds_msdfs_namespaceanchor struct
type ds_msdfs_namespaceanchor struct { 
	*ads_msdfs_namespaceanchor
}

	func Newds_msdfs_namespaceanchorEx1(instance *cim.WmiInstance) (newInstance *ds_msdfs_namespaceanchor, err error) {tmp, err := Newads_msdfs_namespaceanchorEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_msdfs_namespaceanchor {
	ads_msdfs_namespaceanchor: tmp,
	}
	return
	}
	

	func Newds_msdfs_namespaceanchorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_msdfs_namespaceanchor, err error) {tmp, err := Newads_msdfs_namespaceanchorEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_msdfs_namespaceanchor {
	ads_msdfs_namespaceanchor: tmp,
	}
	return
	}
	

