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

// ds_mstpm_informationobjectscontainer struct
type ds_mstpm_informationobjectscontainer struct { 
	*ads_mstpm_informationobjectscontainer
}

	func Newds_mstpm_informationobjectscontainerEx1(instance *cim.WmiInstance) (newInstance *ds_mstpm_informationobjectscontainer, err error) {tmp, err := Newads_mstpm_informationobjectscontainerEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_mstpm_informationobjectscontainer {
	ads_mstpm_informationobjectscontainer: tmp,
	}
	return
	}
	

	func Newds_mstpm_informationobjectscontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_mstpm_informationobjectscontainer, err error) {tmp, err := Newads_mstpm_informationobjectscontainerEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_mstpm_informationobjectscontainer {
	ads_mstpm_informationobjectscontainer: tmp,
	}
	return
	}
	

