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

// ds_ipprotocol struct
type ds_ipprotocol struct { 
	*ads_ipprotocol
}

	func Newds_ipprotocolEx1(instance *cim.WmiInstance) (newInstance *ds_ipprotocol, err error) {tmp, err := Newads_ipprotocolEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_ipprotocol {
	ads_ipprotocol: tmp,
	}
	return
	}
	

	func Newds_ipprotocolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_ipprotocol, err error) {tmp, err := Newads_ipprotocolEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_ipprotocol {
	ads_ipprotocol: tmp,
	}
	return
	}
	

