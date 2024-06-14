// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_msdfsr_globalsettings struct
type ds_msdfsr_globalsettings struct {
	*ads_msdfsr_globalsettings
}

func Newds_msdfsr_globalsettingsEx1(instance *cim.WmiInstance) (newInstance *ds_msdfsr_globalsettings, err error) {
	tmp, err := Newads_msdfsr_globalsettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_globalsettings{
		ads_msdfsr_globalsettings: tmp,
	}
	return
}

func Newds_msdfsr_globalsettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msdfsr_globalsettings, err error) {
	tmp, err := Newads_msdfsr_globalsettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_globalsettings{
		ads_msdfsr_globalsettings: tmp,
	}
	return
}
