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

// ds_comconnectionpoint struct
type ds_comconnectionpoint struct {
	*ads_comconnectionpoint
}

func Newds_comconnectionpointEx1(instance *cim.WmiInstance) (newInstance *ds_comconnectionpoint, err error) {
	tmp, err := Newads_comconnectionpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_comconnectionpoint{
		ads_comconnectionpoint: tmp,
	}
	return
}

func Newds_comconnectionpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_comconnectionpoint, err error) {
	tmp, err := Newads_comconnectionpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_comconnectionpoint{
		ads_comconnectionpoint: tmp,
	}
	return
}
