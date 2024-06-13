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

// ds_ridmanager struct
type ds_ridmanager struct {
	*ads_ridmanager
}

func Newds_ridmanagerEx1(instance *cim.WmiInstance) (newInstance *ds_ridmanager, err error) {
	tmp, err := Newads_ridmanagerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ridmanager{
		ads_ridmanager: tmp,
	}
	return
}

func Newds_ridmanagerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ridmanager, err error) {
	tmp, err := Newads_ridmanagerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ridmanager{
		ads_ridmanager: tmp,
	}
	return
}
