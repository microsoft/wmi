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

// ds_classschema struct
type ds_classschema struct { 
	*ads_classschema
}

	func Newds_classschemaEx1(instance *cim.WmiInstance) (newInstance *ds_classschema, err error) {tmp, err := Newads_classschemaEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_classschema {
	ads_classschema: tmp,
	}
	return
	}
	

	func Newds_classschemaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_classschema, err error) {tmp, err := Newads_classschemaEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_classschema {
	ads_classschema: tmp,
	}
	return
	}
	

