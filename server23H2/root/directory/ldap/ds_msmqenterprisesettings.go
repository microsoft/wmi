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

// ds_msmqenterprisesettings struct
type ds_msmqenterprisesettings struct {
	*ads_msmqenterprisesettings
}

func Newds_msmqenterprisesettingsEx1(instance *cim.WmiInstance) (newInstance *ds_msmqenterprisesettings, err error) {
	tmp, err := Newads_msmqenterprisesettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msmqenterprisesettings{
		ads_msmqenterprisesettings: tmp,
	}
	return
}

func Newds_msmqenterprisesettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msmqenterprisesettings, err error) {
	tmp, err := Newads_msmqenterprisesettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msmqenterprisesettings{
		ads_msmqenterprisesettings: tmp,
	}
	return
}