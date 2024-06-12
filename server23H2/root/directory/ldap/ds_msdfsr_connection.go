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

// ds_msdfsr_connection struct
type ds_msdfsr_connection struct {
	*ads_msdfsr_connection
}

func Newds_msdfsr_connectionEx1(instance *cim.WmiInstance) (newInstance *ds_msdfsr_connection, err error) {
	tmp, err := Newads_msdfsr_connectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_connection{
		ads_msdfsr_connection: tmp,
	}
	return
}

func Newds_msdfsr_connectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msdfsr_connection, err error) {
	tmp, err := Newads_msdfsr_connectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_connection{
		ads_msdfsr_connection: tmp,
	}
	return
}
