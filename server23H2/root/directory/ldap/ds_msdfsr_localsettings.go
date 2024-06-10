// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_msdfsr_localsettings struct
type ds_msdfsr_localsettings struct {
	*ads_msdfsr_localsettings
}

func Newds_msdfsr_localsettingsEx1(instance *cim.WmiInstance) (newInstance *ds_msdfsr_localsettings, err error) {
	tmp, err := Newads_msdfsr_localsettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_localsettings{
		ads_msdfsr_localsettings: tmp,
	}
	return
}

func Newds_msdfsr_localsettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msdfsr_localsettings, err error) {
	tmp, err := Newads_msdfsr_localsettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_localsettings{
		ads_msdfsr_localsettings: tmp,
	}
	return
}
