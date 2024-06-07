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

// ds_ipsecnfa struct
type ds_ipsecnfa struct { 
	*ads_ipsecnfa
}

	func Newds_ipsecnfaEx1(instance *cim.WmiInstance) (newInstance *ds_ipsecnfa, err error) {tmp, err := Newads_ipsecnfaEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_ipsecnfa {
	ads_ipsecnfa: tmp,
	}
	return
	}
	

	func Newds_ipsecnfaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_ipsecnfa, err error) {tmp, err := Newads_ipsecnfaEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_ipsecnfa {
	ads_ipsecnfa: tmp,
	}
	return
	}
	

