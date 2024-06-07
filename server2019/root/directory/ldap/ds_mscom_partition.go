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

// ds_mscom_partition struct
type ds_mscom_partition struct { 
	*ads_mscom_partition
}

	func Newds_mscom_partitionEx1(instance *cim.WmiInstance) (newInstance *ds_mscom_partition, err error) {tmp, err := Newads_mscom_partitionEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_mscom_partition {
	ads_mscom_partition: tmp,
	}
	return
	}
	

	func Newds_mscom_partitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_mscom_partition, err error) {tmp, err := Newads_mscom_partitionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_mscom_partition {
	ads_mscom_partition: tmp,
	}
	return
	}
	

