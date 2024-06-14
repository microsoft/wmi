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

// ds_ftdfs struct
type ds_ftdfs struct {
	*ads_ftdfs
}

func Newds_ftdfsEx1(instance *cim.WmiInstance) (newInstance *ds_ftdfs, err error) {
	tmp, err := Newads_ftdfsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ftdfs{
		ads_ftdfs: tmp,
	}
	return
}

func Newds_ftdfsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ftdfs, err error) {
	tmp, err := Newads_ftdfsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ftdfs{
		ads_ftdfs: tmp,
	}
	return
}
