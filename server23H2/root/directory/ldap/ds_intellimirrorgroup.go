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

// ds_intellimirrorgroup struct
type ds_intellimirrorgroup struct {
	*ads_intellimirrorgroup
}

func Newds_intellimirrorgroupEx1(instance *cim.WmiInstance) (newInstance *ds_intellimirrorgroup, err error) {
	tmp, err := Newads_intellimirrorgroupEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_intellimirrorgroup{
		ads_intellimirrorgroup: tmp,
	}
	return
}

func Newds_intellimirrorgroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_intellimirrorgroup, err error) {
	tmp, err := Newads_intellimirrorgroupEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_intellimirrorgroup{
		ads_intellimirrorgroup: tmp,
	}
	return
}
