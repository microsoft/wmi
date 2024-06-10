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

// ds_person struct
type ds_person struct {
	*ads_person
}

func Newds_personEx1(instance *cim.WmiInstance) (newInstance *ds_person, err error) {
	tmp, err := Newads_personEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_person{
		ads_person: tmp,
	}
	return
}

func Newds_personEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_person, err error) {
	tmp, err := Newads_personEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_person{
		ads_person: tmp,
	}
	return
}
