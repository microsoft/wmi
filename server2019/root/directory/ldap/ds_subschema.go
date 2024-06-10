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

// ds_subschema struct
type ds_subschema struct {
	*ads_subschema
}

func Newds_subschemaEx1(instance *cim.WmiInstance) (newInstance *ds_subschema, err error) {
	tmp, err := Newads_subschemaEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_subschema{
		ads_subschema: tmp,
	}
	return
}

func Newds_subschemaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_subschema, err error) {
	tmp, err := Newads_subschemaEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_subschema{
		ads_subschema: tmp,
	}
	return
}
