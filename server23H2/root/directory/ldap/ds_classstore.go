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

// ds_classstore struct
type ds_classstore struct {
	*ads_classstore
}

func Newds_classstoreEx1(instance *cim.WmiInstance) (newInstance *ds_classstore, err error) {
	tmp, err := Newads_classstoreEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_classstore{
		ads_classstore: tmp,
	}
	return
}

func Newds_classstoreEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_classstore, err error) {
	tmp, err := Newads_classstoreEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_classstore{
		ads_classstore: tmp,
	}
	return
}
