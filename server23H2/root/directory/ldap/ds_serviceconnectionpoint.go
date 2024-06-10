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

// ds_serviceconnectionpoint struct
type ds_serviceconnectionpoint struct {
	*ads_serviceconnectionpoint
}

func Newds_serviceconnectionpointEx1(instance *cim.WmiInstance) (newInstance *ds_serviceconnectionpoint, err error) {
	tmp, err := Newads_serviceconnectionpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_serviceconnectionpoint{
		ads_serviceconnectionpoint: tmp,
	}
	return
}

func Newds_serviceconnectionpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_serviceconnectionpoint, err error) {
	tmp, err := Newads_serviceconnectionpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_serviceconnectionpoint{
		ads_serviceconnectionpoint: tmp,
	}
	return
}
