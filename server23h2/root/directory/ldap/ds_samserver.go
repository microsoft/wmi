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

// ds_samserver struct
type ds_samserver struct {
	*ads_samserver
}

func Newds_samserverEx1(instance *cim.WmiInstance) (newInstance *ds_samserver, err error) {
	tmp, err := Newads_samserverEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_samserver{
		ads_samserver: tmp,
	}
	return
}

func Newds_samserverEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_samserver, err error) {
	tmp, err := Newads_samserverEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_samserver{
		ads_samserver: tmp,
	}
	return
}
