// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_locality struct
type ds_locality struct {
	*ads_locality
}

func Newds_localityEx1(instance *cim.WmiInstance) (newInstance *ds_locality, err error) {
	tmp, err := Newads_localityEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_locality{
		ads_locality: tmp,
	}
	return
}

func Newds_localityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_locality, err error) {
	tmp, err := Newads_localityEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_locality{
		ads_locality: tmp,
	}
	return
}
