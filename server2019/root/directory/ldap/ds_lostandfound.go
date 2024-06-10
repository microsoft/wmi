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

// ds_lostandfound struct
type ds_lostandfound struct {
	*ads_lostandfound
}

func Newds_lostandfoundEx1(instance *cim.WmiInstance) (newInstance *ds_lostandfound, err error) {
	tmp, err := Newads_lostandfoundEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_lostandfound{
		ads_lostandfound: tmp,
	}
	return
}

func Newds_lostandfoundEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_lostandfound, err error) {
	tmp, err := Newads_lostandfoundEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_lostandfound{
		ads_lostandfound: tmp,
	}
	return
}
