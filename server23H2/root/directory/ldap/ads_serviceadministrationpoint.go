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

// ads_serviceadministrationpoint struct
type ads_serviceadministrationpoint struct {
	*ads_serviceconnectionpoint
}

func Newads_serviceadministrationpointEx1(instance *cim.WmiInstance) (newInstance *ads_serviceadministrationpoint, err error) {
	tmp, err := Newads_serviceconnectionpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_serviceadministrationpoint{
		ads_serviceconnectionpoint: tmp,
	}
	return
}

func Newads_serviceadministrationpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_serviceadministrationpoint, err error) {
	tmp, err := Newads_serviceconnectionpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_serviceadministrationpoint{
		ads_serviceconnectionpoint: tmp,
	}
	return
}
