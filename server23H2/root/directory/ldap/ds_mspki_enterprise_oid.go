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

// ds_mspki_enterprise_oid struct
type ds_mspki_enterprise_oid struct {
	*ads_mspki_enterprise_oid
}

func Newds_mspki_enterprise_oidEx1(instance *cim.WmiInstance) (newInstance *ds_mspki_enterprise_oid, err error) {
	tmp, err := Newads_mspki_enterprise_oidEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_mspki_enterprise_oid{
		ads_mspki_enterprise_oid: tmp,
	}
	return
}

func Newds_mspki_enterprise_oidEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_mspki_enterprise_oid, err error) {
	tmp, err := Newads_mspki_enterprise_oidEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_mspki_enterprise_oid{
		ads_mspki_enterprise_oid: tmp,
	}
	return
}
