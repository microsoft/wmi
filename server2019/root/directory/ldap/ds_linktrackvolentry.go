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

// ds_linktrackvolentry struct
type ds_linktrackvolentry struct { 
	*ads_linktrackvolentry
}

	func Newds_linktrackvolentryEx1(instance *cim.WmiInstance) (newInstance *ds_linktrackvolentry, err error) {tmp, err := Newads_linktrackvolentryEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_linktrackvolentry {
	ads_linktrackvolentry: tmp,
	}
	return
	}
	

	func Newds_linktrackvolentryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_linktrackvolentry, err error) {tmp, err := Newads_linktrackvolentryEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_linktrackvolentry {
	ads_linktrackvolentry: tmp,
	}
	return
	}
	

