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

// ds_msmq_group struct
type ds_msmq_group struct {
	*ads_msmq_group
}

func Newds_msmq_groupEx1(instance *cim.WmiInstance) (newInstance *ds_msmq_group, err error) {
	tmp, err := Newads_msmq_groupEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msmq_group{
		ads_msmq_group: tmp,
	}
	return
}

func Newds_msmq_groupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msmq_group, err error) {
	tmp, err := Newads_msmq_groupEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msmq_group{
		ads_msmq_group: tmp,
	}
	return
}
