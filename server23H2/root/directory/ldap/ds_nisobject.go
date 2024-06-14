// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_nisobject struct
type ds_nisobject struct {
	*ads_nisobject
}

func Newds_nisobjectEx1(instance *cim.WmiInstance) (newInstance *ds_nisobject, err error) {
	tmp, err := Newads_nisobjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_nisobject{
		ads_nisobject: tmp,
	}
	return
}

func Newds_nisobjectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_nisobject, err error) {
	tmp, err := Newads_nisobjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_nisobject{
		ads_nisobject: tmp,
	}
	return
}
