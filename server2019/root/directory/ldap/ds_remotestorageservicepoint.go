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

// ds_remotestorageservicepoint struct
type ds_remotestorageservicepoint struct { 
	*ads_remotestorageservicepoint
}

	func Newds_remotestorageservicepointEx1(instance *cim.WmiInstance) (newInstance *ds_remotestorageservicepoint, err error) {tmp, err := Newads_remotestorageservicepointEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_remotestorageservicepoint {
	ads_remotestorageservicepoint: tmp,
	}
	return
	}
	

	func Newds_remotestorageservicepointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_remotestorageservicepoint, err error) {tmp, err := Newads_remotestorageservicepointEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_remotestorageservicepoint {
	ads_remotestorageservicepoint: tmp,
	}
	return
	}
	

