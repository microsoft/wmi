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

// ds_typelibrary struct
type ds_typelibrary struct {
	*ads_typelibrary
}

func Newds_typelibraryEx1(instance *cim.WmiInstance) (newInstance *ds_typelibrary, err error) {
	tmp, err := Newads_typelibraryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_typelibrary{
		ads_typelibrary: tmp,
	}
	return
}

func Newds_typelibraryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_typelibrary, err error) {
	tmp, err := Newads_typelibraryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_typelibrary{
		ads_typelibrary: tmp,
	}
	return
}
