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

// ds_acsresourcelimits struct
type ds_acsresourcelimits struct { 
	*ads_acsresourcelimits
}

	func Newds_acsresourcelimitsEx1(instance *cim.WmiInstance) (newInstance *ds_acsresourcelimits, err error) {tmp, err := Newads_acsresourcelimitsEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_acsresourcelimits {
	ads_acsresourcelimits: tmp,
	}
	return
	}
	

	func Newds_acsresourcelimitsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_acsresourcelimits, err error) {tmp, err := Newads_acsresourcelimitsEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_acsresourcelimits {
	ads_acsresourcelimits: tmp,
	}
	return
	}
	

