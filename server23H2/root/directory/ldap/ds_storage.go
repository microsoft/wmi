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

// ds_storage struct
type ds_storage struct {
	*ads_storage
}

func Newds_storageEx1(instance *cim.WmiInstance) (newInstance *ds_storage, err error) {
	tmp, err := Newads_storageEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_storage{
		ads_storage: tmp,
	}
	return
}

func Newds_storageEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_storage, err error) {
	tmp, err := Newads_storageEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_storage{
		ads_storage: tmp,
	}
	return
}