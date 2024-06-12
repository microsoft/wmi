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

// ds_serviceinstance struct
type ds_serviceinstance struct {
	*ads_serviceinstance
}

func Newds_serviceinstanceEx1(instance *cim.WmiInstance) (newInstance *ds_serviceinstance, err error) {
	tmp, err := Newads_serviceinstanceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_serviceinstance{
		ads_serviceinstance: tmp,
	}
	return
}

func Newds_serviceinstanceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_serviceinstance, err error) {
	tmp, err := Newads_serviceinstanceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_serviceinstance{
		ads_serviceinstance: tmp,
	}
	return
}
