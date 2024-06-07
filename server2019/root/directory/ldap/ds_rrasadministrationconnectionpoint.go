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

// ds_rrasadministrationconnectionpoint struct
type ds_rrasadministrationconnectionpoint struct { 
	*ads_rrasadministrationconnectionpoint
}

	func Newds_rrasadministrationconnectionpointEx1(instance *cim.WmiInstance) (newInstance *ds_rrasadministrationconnectionpoint, err error) {tmp, err := Newads_rrasadministrationconnectionpointEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_rrasadministrationconnectionpoint {
	ads_rrasadministrationconnectionpoint: tmp,
	}
	return
	}
	

	func Newds_rrasadministrationconnectionpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_rrasadministrationconnectionpoint, err error) {tmp, err := Newads_rrasadministrationconnectionpointEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_rrasadministrationconnectionpoint {
	ads_rrasadministrationconnectionpoint: tmp,
	}
	return
	}
	

