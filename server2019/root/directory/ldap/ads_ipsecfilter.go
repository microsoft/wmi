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

// ads_ipsecfilter struct
type ads_ipsecfilter struct { 
	*ds_ipsecbase
}

	func Newads_ipsecfilterEx1(instance *cim.WmiInstance) (newInstance *ads_ipsecfilter, err error) {tmp, err := Newds_ipsecbaseEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_ipsecfilter {
	ds_ipsecbase: tmp,
	}
	return
	}
	

	func Newads_ipsecfilterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_ipsecfilter, err error) {tmp, err := Newds_ipsecbaseEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_ipsecfilter {
	ds_ipsecbase: tmp,
	}
	return
	}
	

