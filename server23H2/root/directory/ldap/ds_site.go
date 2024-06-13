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

// ds_site struct
type ds_site struct {
	*ads_site
}

func Newds_siteEx1(instance *cim.WmiInstance) (newInstance *ds_site, err error) {
	tmp, err := Newads_siteEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_site{
		ads_site: tmp,
	}
	return
}

func Newds_siteEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_site, err error) {
	tmp, err := Newads_siteEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_site{
		ads_site: tmp,
	}
	return
}
