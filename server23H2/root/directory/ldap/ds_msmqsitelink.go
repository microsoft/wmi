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

// ds_msmqsitelink struct
type ds_msmqsitelink struct {
	*ads_msmqsitelink
}

func Newds_msmqsitelinkEx1(instance *cim.WmiInstance) (newInstance *ds_msmqsitelink, err error) {
	tmp, err := Newads_msmqsitelinkEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msmqsitelink{
		ads_msmqsitelink: tmp,
	}
	return
}

func Newds_msmqsitelinkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msmqsitelink, err error) {
	tmp, err := Newads_msmqsitelinkEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msmqsitelink{
		ads_msmqsitelink: tmp,
	}
	return
}
