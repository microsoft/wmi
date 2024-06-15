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

// ds_inetorgperson struct
type ds_inetorgperson struct {
	*ads_inetorgperson
}

func Newds_inetorgpersonEx1(instance *cim.WmiInstance) (newInstance *ds_inetorgperson, err error) {
	tmp, err := Newads_inetorgpersonEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_inetorgperson{
		ads_inetorgperson: tmp,
	}
	return
}

func Newds_inetorgpersonEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_inetorgperson, err error) {
	tmp, err := Newads_inetorgpersonEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_inetorgperson{
		ads_inetorgperson: tmp,
	}
	return
}
