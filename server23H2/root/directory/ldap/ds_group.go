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

// ds_group struct
type ds_group struct {
	*ads_group
}

func Newds_groupEx1(instance *cim.WmiInstance) (newInstance *ds_group, err error) {
	tmp, err := Newads_groupEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_group{
		ads_group: tmp,
	}
	return
}

func Newds_groupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_group, err error) {
	tmp, err := Newads_groupEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_group{
		ads_group: tmp,
	}
	return
}
