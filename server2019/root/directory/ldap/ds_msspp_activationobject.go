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

// ds_msspp_activationobject struct
type ds_msspp_activationobject struct { 
	*ads_msspp_activationobject
}

	func Newds_msspp_activationobjectEx1(instance *cim.WmiInstance) (newInstance *ds_msspp_activationobject, err error) {tmp, err := Newads_msspp_activationobjectEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_msspp_activationobject {
	ads_msspp_activationobject: tmp,
	}
	return
	}
	

	func Newds_msspp_activationobjectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_msspp_activationobject, err error) {tmp, err := Newads_msspp_activationobjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_msspp_activationobject {
	ads_msspp_activationobject: tmp,
	}
	return
	}
	

