// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_ntfrsreplicaset struct
type ds_ntfrsreplicaset struct {
	*ads_ntfrsreplicaset
}

func Newds_ntfrsreplicasetEx1(instance *cim.WmiInstance) (newInstance *ds_ntfrsreplicaset, err error) {
	tmp, err := Newads_ntfrsreplicasetEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ntfrsreplicaset{
		ads_ntfrsreplicaset: tmp,
	}
	return
}

func Newds_ntfrsreplicasetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ntfrsreplicaset, err error) {
	tmp, err := Newads_ntfrsreplicasetEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ntfrsreplicaset{
		ads_ntfrsreplicaset: tmp,
	}
	return
}
