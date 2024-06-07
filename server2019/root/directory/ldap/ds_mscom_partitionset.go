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

// ds_mscom_partitionset struct
type ds_mscom_partitionset struct { 
	*ads_mscom_partitionset
}

	func Newds_mscom_partitionsetEx1(instance *cim.WmiInstance) (newInstance *ds_mscom_partitionset, err error) {tmp, err := Newads_mscom_partitionsetEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_mscom_partitionset {
	ads_mscom_partitionset: tmp,
	}
	return
	}
	

	func Newds_mscom_partitionsetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_mscom_partitionset, err error) {tmp, err := Newads_mscom_partitionsetEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_mscom_partitionset {
	ads_mscom_partitionset: tmp,
	}
	return
	}
	

