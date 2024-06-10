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

// ds_groupofnames struct
type ds_groupofnames struct {
	*ads_groupofnames
}

func Newds_groupofnamesEx1(instance *cim.WmiInstance) (newInstance *ds_groupofnames, err error) {
	tmp, err := Newads_groupofnamesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_groupofnames{
		ads_groupofnames: tmp,
	}
	return
}

func Newds_groupofnamesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_groupofnames, err error) {
	tmp, err := Newads_groupofnamesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_groupofnames{
		ads_groupofnames: tmp,
	}
	return
}
