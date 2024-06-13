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

// ds_organization struct
type ds_organization struct {
	*ads_organization
}

func Newds_organizationEx1(instance *cim.WmiInstance) (newInstance *ds_organization, err error) {
	tmp, err := Newads_organizationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_organization{
		ads_organization: tmp,
	}
	return
}

func Newds_organizationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_organization, err error) {
	tmp, err := Newads_organizationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_organization{
		ads_organization: tmp,
	}
	return
}
