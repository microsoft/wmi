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

// ds_grouppolicycontainer struct
type ds_grouppolicycontainer struct { 
	*ads_grouppolicycontainer
}

	func Newds_grouppolicycontainerEx1(instance *cim.WmiInstance) (newInstance *ds_grouppolicycontainer, err error) {tmp, err := Newads_grouppolicycontainerEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_grouppolicycontainer {
	ads_grouppolicycontainer: tmp,
	}
	return
	}
	

	func Newds_grouppolicycontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_grouppolicycontainer, err error) {tmp, err := Newads_grouppolicycontainerEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_grouppolicycontainer {
	ads_grouppolicycontainer: tmp,
	}
	return
	}
	

