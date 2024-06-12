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

// ds_acssubnet struct
type ds_acssubnet struct {
	*ads_acssubnet
}

func Newds_acssubnetEx1(instance *cim.WmiInstance) (newInstance *ds_acssubnet, err error) {
	tmp, err := Newads_acssubnetEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_acssubnet{
		ads_acssubnet: tmp,
	}
	return
}

func Newds_acssubnetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_acssubnet, err error) {
	tmp, err := Newads_acssubnetEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_acssubnet{
		ads_acssubnet: tmp,
	}
	return
}
