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

// ds_ipnetwork struct
type ds_ipnetwork struct { 
	*ads_ipnetwork
}

	func Newds_ipnetworkEx1(instance *cim.WmiInstance) (newInstance *ds_ipnetwork, err error) {tmp, err := Newads_ipnetworkEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_ipnetwork {
	ads_ipnetwork: tmp,
	}
	return
	}
	

	func Newds_ipnetworkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_ipnetwork, err error) {tmp, err := Newads_ipnetworkEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_ipnetwork {
	ads_ipnetwork: tmp,
	}
	return
	}
	

