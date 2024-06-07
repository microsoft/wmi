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

// ds_documentseries struct
type ds_documentseries struct { 
	*ads_documentseries
}

	func Newds_documentseriesEx1(instance *cim.WmiInstance) (newInstance *ds_documentseries, err error) {tmp, err := Newads_documentseriesEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_documentseries {
	ads_documentseries: tmp,
	}
	return
	}
	

	func Newds_documentseriesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_documentseries, err error) {tmp, err := Newads_documentseriesEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_documentseries {
	ads_documentseries: tmp,
	}
	return
	}
	

