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

// ds_msfve_recoveryinformation struct
type ds_msfve_recoveryinformation struct { 
	*ads_msfve_recoveryinformation
}

	func Newds_msfve_recoveryinformationEx1(instance *cim.WmiInstance) (newInstance *ds_msfve_recoveryinformation, err error) {tmp, err := Newads_msfve_recoveryinformationEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_msfve_recoveryinformation {
	ads_msfve_recoveryinformation: tmp,
	}
	return
	}
	

	func Newds_msfve_recoveryinformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_msfve_recoveryinformation, err error) {tmp, err := Newads_msfve_recoveryinformationEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_msfve_recoveryinformation {
	ads_msfve_recoveryinformation: tmp,
	}
	return
	}
	

