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

// ds_msdfsr_member struct
type ds_msdfsr_member struct {
	*ads_msdfsr_member
}

func Newds_msdfsr_memberEx1(instance *cim.WmiInstance) (newInstance *ds_msdfsr_member, err error) {
	tmp, err := Newads_msdfsr_memberEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_member{
		ads_msdfsr_member: tmp,
	}
	return
}

func Newds_msdfsr_memberEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msdfsr_member, err error) {
	tmp, err := Newads_msdfsr_memberEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_member{
		ads_msdfsr_member: tmp,
	}
	return
}
