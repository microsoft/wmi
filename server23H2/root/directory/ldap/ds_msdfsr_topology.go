// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_msdfsr_topology struct
type ds_msdfsr_topology struct {
	*ads_msdfsr_topology
}

func Newds_msdfsr_topologyEx1(instance *cim.WmiInstance) (newInstance *ds_msdfsr_topology, err error) {
	tmp, err := Newads_msdfsr_topologyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_topology{
		ads_msdfsr_topology: tmp,
	}
	return
}

func Newds_msdfsr_topologyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msdfsr_topology, err error) {
	tmp, err := Newads_msdfsr_topologyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_topology{
		ads_msdfsr_topology: tmp,
	}
	return
}
