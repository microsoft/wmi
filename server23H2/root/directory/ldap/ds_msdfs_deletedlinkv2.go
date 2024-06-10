// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_msdfs_deletedlinkv2 struct
type ds_msdfs_deletedlinkv2 struct {
	*ads_msdfs_deletedlinkv2
}

func Newds_msdfs_deletedlinkv2Ex1(instance *cim.WmiInstance) (newInstance *ds_msdfs_deletedlinkv2, err error) {
	tmp, err := Newads_msdfs_deletedlinkv2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msdfs_deletedlinkv2{
		ads_msdfs_deletedlinkv2: tmp,
	}
	return
}

func Newds_msdfs_deletedlinkv2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msdfs_deletedlinkv2, err error) {
	tmp, err := Newads_msdfs_deletedlinkv2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msdfs_deletedlinkv2{
		ads_msdfs_deletedlinkv2: tmp,
	}
	return
}
