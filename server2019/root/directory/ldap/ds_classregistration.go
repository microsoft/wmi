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

// ds_classregistration struct
type ds_classregistration struct {
	*ads_classregistration
}

func Newds_classregistrationEx1(instance *cim.WmiInstance) (newInstance *ds_classregistration, err error) {
	tmp, err := Newads_classregistrationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_classregistration{
		ads_classregistration: tmp,
	}
	return
}

func Newds_classregistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_classregistration, err error) {
	tmp, err := Newads_classregistrationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_classregistration{
		ads_classregistration: tmp,
	}
	return
}
