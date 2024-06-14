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

// ds_mssfu30netid struct
type ds_mssfu30netid struct {
	*ads_mssfu30netid
}

func Newds_mssfu30netidEx1(instance *cim.WmiInstance) (newInstance *ds_mssfu30netid, err error) {
	tmp, err := Newads_mssfu30netidEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_mssfu30netid{
		ads_mssfu30netid: tmp,
	}
	return
}

func Newds_mssfu30netidEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_mssfu30netid, err error) {
	tmp, err := Newads_mssfu30netidEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_mssfu30netid{
		ads_mssfu30netid: tmp,
	}
	return
}
