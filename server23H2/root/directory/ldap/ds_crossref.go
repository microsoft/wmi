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

// ds_crossref struct
type ds_crossref struct {
	*ads_crossref
}

func Newds_crossrefEx1(instance *cim.WmiInstance) (newInstance *ds_crossref, err error) {
	tmp, err := Newads_crossrefEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_crossref{
		ads_crossref: tmp,
	}
	return
}

func Newds_crossrefEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_crossref, err error) {
	tmp, err := Newads_crossrefEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_crossref{
		ads_crossref: tmp,
	}
	return
}
