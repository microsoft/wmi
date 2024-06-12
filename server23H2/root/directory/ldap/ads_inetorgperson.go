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

// ads_inetorgperson struct
type ads_inetorgperson struct {
	*ads_user
}

func Newads_inetorgpersonEx1(instance *cim.WmiInstance) (newInstance *ads_inetorgperson, err error) {
	tmp, err := Newads_userEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_inetorgperson{
		ads_user: tmp,
	}
	return
}

func Newads_inetorgpersonEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_inetorgperson, err error) {
	tmp, err := Newads_userEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_inetorgperson{
		ads_user: tmp,
	}
	return
}
