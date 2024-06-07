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

// ds_mstpm_informationobject struct
type ds_mstpm_informationobject struct { 
	*ads_mstpm_informationobject
}

	func Newds_mstpm_informationobjectEx1(instance *cim.WmiInstance) (newInstance *ds_mstpm_informationobject, err error) {tmp, err := Newads_mstpm_informationobjectEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_mstpm_informationobject {
	ads_mstpm_informationobject: tmp,
	}
	return
	}
	

	func Newds_mstpm_informationobjectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_mstpm_informationobject, err error) {tmp, err := Newads_mstpm_informationobjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_mstpm_informationobject {
	ads_mstpm_informationobject: tmp,
	}
	return
	}
	

