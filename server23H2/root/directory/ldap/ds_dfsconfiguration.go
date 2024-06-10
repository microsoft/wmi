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

// ds_dfsconfiguration struct
type ds_dfsconfiguration struct {
	*ads_dfsconfiguration
}

func Newds_dfsconfigurationEx1(instance *cim.WmiInstance) (newInstance *ds_dfsconfiguration, err error) {
	tmp, err := Newads_dfsconfigurationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_dfsconfiguration{
		ads_dfsconfiguration: tmp,
	}
	return
}

func Newds_dfsconfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_dfsconfiguration, err error) {
	tmp, err := Newads_dfsconfigurationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_dfsconfiguration{
		ads_dfsconfiguration: tmp,
	}
	return
}
