// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_msmqqueue struct
type ds_msmqqueue struct {
	*ads_msmqqueue
}

func Newds_msmqqueueEx1(instance *cim.WmiInstance) (newInstance *ds_msmqqueue, err error) {
	tmp, err := Newads_msmqqueueEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msmqqueue{
		ads_msmqqueue: tmp,
	}
	return
}

func Newds_msmqqueueEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msmqqueue, err error) {
	tmp, err := Newads_msmqqueueEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msmqqueue{
		ads_msmqqueue: tmp,
	}
	return
}
