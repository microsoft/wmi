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

// ds_ntdssitesettings struct
type ds_ntdssitesettings struct {
	*ads_ntdssitesettings
}

func Newds_ntdssitesettingsEx1(instance *cim.WmiInstance) (newInstance *ds_ntdssitesettings, err error) {
	tmp, err := Newads_ntdssitesettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ntdssitesettings{
		ads_ntdssitesettings: tmp,
	}
	return
}

func Newds_ntdssitesettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ntdssitesettings, err error) {
	tmp, err := Newads_ntdssitesettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ntdssitesettings{
		ads_ntdssitesettings: tmp,
	}
	return
}
