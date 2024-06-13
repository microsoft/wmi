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

// ds_msdfsr_replicationgroup struct
type ds_msdfsr_replicationgroup struct {
	*ads_msdfsr_replicationgroup
}

func Newds_msdfsr_replicationgroupEx1(instance *cim.WmiInstance) (newInstance *ds_msdfsr_replicationgroup, err error) {
	tmp, err := Newads_msdfsr_replicationgroupEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_replicationgroup{
		ads_msdfsr_replicationgroup: tmp,
	}
	return
}

func Newds_msdfsr_replicationgroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msdfsr_replicationgroup, err error) {
	tmp, err := Newads_msdfsr_replicationgroupEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_replicationgroup{
		ads_msdfsr_replicationgroup: tmp,
	}
	return
}
