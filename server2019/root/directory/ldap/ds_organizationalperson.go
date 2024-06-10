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

// ds_organizationalperson struct
type ds_organizationalperson struct {
	*ads_organizationalperson
}

func Newds_organizationalpersonEx1(instance *cim.WmiInstance) (newInstance *ds_organizationalperson, err error) {
	tmp, err := Newads_organizationalpersonEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_organizationalperson{
		ads_organizationalperson: tmp,
	}
	return
}

func Newds_organizationalpersonEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_organizationalperson, err error) {
	tmp, err := Newads_organizationalpersonEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_organizationalperson{
		ads_organizationalperson: tmp,
	}
	return
}
