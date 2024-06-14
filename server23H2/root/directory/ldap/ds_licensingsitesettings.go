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

// ds_licensingsitesettings struct
type ds_licensingsitesettings struct {
	*ads_licensingsitesettings
}

func Newds_licensingsitesettingsEx1(instance *cim.WmiInstance) (newInstance *ds_licensingsitesettings, err error) {
	tmp, err := Newads_licensingsitesettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_licensingsitesettings{
		ads_licensingsitesettings: tmp,
	}
	return
}

func Newds_licensingsitesettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_licensingsitesettings, err error) {
	tmp, err := Newads_licensingsitesettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_licensingsitesettings{
		ads_licensingsitesettings: tmp,
	}
	return
}
