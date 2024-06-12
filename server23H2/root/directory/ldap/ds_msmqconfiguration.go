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

// ds_msmqconfiguration struct
type ds_msmqconfiguration struct {
	*ads_msmqconfiguration
}

func Newds_msmqconfigurationEx1(instance *cim.WmiInstance) (newInstance *ds_msmqconfiguration, err error) {
	tmp, err := Newads_msmqconfigurationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msmqconfiguration{
		ads_msmqconfiguration: tmp,
	}
	return
}

func Newds_msmqconfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msmqconfiguration, err error) {
	tmp, err := Newads_msmqconfigurationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msmqconfiguration{
		ads_msmqconfiguration: tmp,
	}
	return
}
