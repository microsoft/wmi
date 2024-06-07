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

// ds_sitescontainer struct
type ds_sitescontainer struct { 
	*ads_sitescontainer
}

	func Newds_sitescontainerEx1(instance *cim.WmiInstance) (newInstance *ds_sitescontainer, err error) {tmp, err := Newads_sitescontainerEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_sitescontainer {
	ads_sitescontainer: tmp,
	}
	return
	}
	

	func Newds_sitescontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_sitescontainer, err error) {tmp, err := Newads_sitescontainerEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_sitescontainer {
	ads_sitescontainer: tmp,
	}
	return
	}
	

