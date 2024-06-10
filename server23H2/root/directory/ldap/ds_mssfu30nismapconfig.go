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

// ds_mssfu30nismapconfig struct
type ds_mssfu30nismapconfig struct {
	*ads_mssfu30nismapconfig
}

func Newds_mssfu30nismapconfigEx1(instance *cim.WmiInstance) (newInstance *ds_mssfu30nismapconfig, err error) {
	tmp, err := Newads_mssfu30nismapconfigEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_mssfu30nismapconfig{
		ads_mssfu30nismapconfig: tmp,
	}
	return
}

func Newds_mssfu30nismapconfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_mssfu30nismapconfig, err error) {
	tmp, err := Newads_mssfu30nismapconfigEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_mssfu30nismapconfig{
		ads_mssfu30nismapconfig: tmp,
	}
	return
}
