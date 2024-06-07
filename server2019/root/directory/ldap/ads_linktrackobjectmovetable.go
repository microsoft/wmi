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

// ads_linktrackobjectmovetable struct
type ads_linktrackobjectmovetable struct { 
	*ads_filelinktracking
}

	func Newads_linktrackobjectmovetableEx1(instance *cim.WmiInstance) (newInstance *ads_linktrackobjectmovetable, err error) {tmp, err := Newads_filelinktrackingEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_linktrackobjectmovetable {
	ads_filelinktracking: tmp,
	}
	return
	}
	

	func Newads_linktrackobjectmovetableEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_linktrackobjectmovetable, err error) {tmp, err := Newads_filelinktrackingEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_linktrackobjectmovetable {
	ads_filelinktracking: tmp,
	}
	return
	}
	

