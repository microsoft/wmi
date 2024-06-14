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

// ds_nismap struct
type ds_nismap struct {
	*ads_nismap
}

func Newds_nismapEx1(instance *cim.WmiInstance) (newInstance *ds_nismap, err error) {
	tmp, err := Newads_nismapEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_nismap{
		ads_nismap: tmp,
	}
	return
}

func Newds_nismapEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_nismap, err error) {
	tmp, err := Newads_nismapEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_nismap{
		ads_nismap: tmp,
	}
	return
}
