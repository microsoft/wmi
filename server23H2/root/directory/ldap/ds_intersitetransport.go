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

// ds_intersitetransport struct
type ds_intersitetransport struct {
	*ads_intersitetransport
}

func Newds_intersitetransportEx1(instance *cim.WmiInstance) (newInstance *ds_intersitetransport, err error) {
	tmp, err := Newads_intersitetransportEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_intersitetransport{
		ads_intersitetransport: tmp,
	}
	return
}

func Newds_intersitetransportEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_intersitetransport, err error) {
	tmp, err := Newads_intersitetransportEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_intersitetransport{
		ads_intersitetransport: tmp,
	}
	return
}
