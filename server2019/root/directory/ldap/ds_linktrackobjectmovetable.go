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

// ds_linktrackobjectmovetable struct
type ds_linktrackobjectmovetable struct { 
	*ads_linktrackobjectmovetable
}

	func Newds_linktrackobjectmovetableEx1(instance *cim.WmiInstance) (newInstance *ds_linktrackobjectmovetable, err error) {tmp, err := Newads_linktrackobjectmovetableEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_linktrackobjectmovetable {
	ads_linktrackobjectmovetable: tmp,
	}
	return
	}
	

	func Newds_linktrackobjectmovetableEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_linktrackobjectmovetable, err error) {tmp, err := Newads_linktrackobjectmovetableEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_linktrackobjectmovetable {
	ads_linktrackobjectmovetable: tmp,
	}
	return
	}
	

