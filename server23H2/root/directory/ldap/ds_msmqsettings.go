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

// ds_msmqsettings struct
type ds_msmqsettings struct {
	*ads_msmqsettings
}

func Newds_msmqsettingsEx1(instance *cim.WmiInstance) (newInstance *ds_msmqsettings, err error) {
	tmp, err := Newads_msmqsettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msmqsettings{
		ads_msmqsettings: tmp,
	}
	return
}

func Newds_msmqsettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msmqsettings, err error) {
	tmp, err := Newads_msmqsettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msmqsettings{
		ads_msmqsettings: tmp,
	}
	return
}
