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

// ds_indexservercatalog struct
type ds_indexservercatalog struct {
	*ads_indexservercatalog
}

func Newds_indexservercatalogEx1(instance *cim.WmiInstance) (newInstance *ds_indexservercatalog, err error) {
	tmp, err := Newads_indexservercatalogEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_indexservercatalog{
		ads_indexservercatalog: tmp,
	}
	return
}

func Newds_indexservercatalogEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_indexservercatalog, err error) {
	tmp, err := Newads_indexservercatalogEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_indexservercatalog{
		ads_indexservercatalog: tmp,
	}
	return
}
