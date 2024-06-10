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

// ds_serviceadministrationpoint struct
type ds_serviceadministrationpoint struct {
	*ads_serviceadministrationpoint
}

func Newds_serviceadministrationpointEx1(instance *cim.WmiInstance) (newInstance *ds_serviceadministrationpoint, err error) {
	tmp, err := Newads_serviceadministrationpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_serviceadministrationpoint{
		ads_serviceadministrationpoint: tmp,
	}
	return
}

func Newds_serviceadministrationpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_serviceadministrationpoint, err error) {
	tmp, err := Newads_serviceadministrationpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_serviceadministrationpoint{
		ads_serviceadministrationpoint: tmp,
	}
	return
}
