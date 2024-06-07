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

// ds_pkicertificatetemplate struct
type ds_pkicertificatetemplate struct { 
	*ads_pkicertificatetemplate
}

	func Newds_pkicertificatetemplateEx1(instance *cim.WmiInstance) (newInstance *ds_pkicertificatetemplate, err error) {tmp, err := Newads_pkicertificatetemplateEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_pkicertificatetemplate {
	ads_pkicertificatetemplate: tmp,
	}
	return
	}
	

	func Newds_pkicertificatetemplateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_pkicertificatetemplate, err error) {tmp, err := Newads_pkicertificatetemplateEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_pkicertificatetemplate {
	ads_pkicertificatetemplate: tmp,
	}
	return
	}
	

