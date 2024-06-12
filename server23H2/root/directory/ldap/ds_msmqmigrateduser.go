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

// ds_msmqmigrateduser struct
type ds_msmqmigrateduser struct {
	*ads_msmqmigrateduser
}

func Newds_msmqmigrateduserEx1(instance *cim.WmiInstance) (newInstance *ds_msmqmigrateduser, err error) {
	tmp, err := Newads_msmqmigrateduserEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msmqmigrateduser{
		ads_msmqmigrateduser: tmp,
	}
	return
}

func Newds_msmqmigrateduserEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msmqmigrateduser, err error) {
	tmp, err := Newads_msmqmigrateduserEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msmqmigrateduser{
		ads_msmqmigrateduser: tmp,
	}
	return
}
