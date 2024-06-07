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

// ds_residentialperson struct
type ds_residentialperson struct { 
	*ads_residentialperson
}

	func Newds_residentialpersonEx1(instance *cim.WmiInstance) (newInstance *ds_residentialperson, err error) {tmp, err := Newads_residentialpersonEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_residentialperson {
	ads_residentialperson: tmp,
	}
	return
	}
	

	func Newds_residentialpersonEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_residentialperson, err error) {tmp, err := Newads_residentialpersonEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_residentialperson {
	ads_residentialperson: tmp,
	}
	return
	}
	

