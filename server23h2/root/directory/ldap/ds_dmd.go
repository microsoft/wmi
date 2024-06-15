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

// ds_dmd struct
type ds_dmd struct {
	*ads_dmd
}

func Newds_dmdEx1(instance *cim.WmiInstance) (newInstance *ds_dmd, err error) {
	tmp, err := Newads_dmdEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_dmd{
		ads_dmd: tmp,
	}
	return
}

func Newds_dmdEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_dmd, err error) {
	tmp, err := Newads_dmdEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_dmd{
		ads_dmd: tmp,
	}
	return
}
