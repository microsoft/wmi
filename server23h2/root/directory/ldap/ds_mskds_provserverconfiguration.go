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

// ds_mskds_provserverconfiguration struct
type ds_mskds_provserverconfiguration struct {
	*ads_mskds_provserverconfiguration
}

func Newds_mskds_provserverconfigurationEx1(instance *cim.WmiInstance) (newInstance *ds_mskds_provserverconfiguration, err error) {
	tmp, err := Newads_mskds_provserverconfigurationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_mskds_provserverconfiguration{
		ads_mskds_provserverconfiguration: tmp,
	}
	return
}

func Newds_mskds_provserverconfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_mskds_provserverconfiguration, err error) {
	tmp, err := Newads_mskds_provserverconfigurationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_mskds_provserverconfiguration{
		ads_mskds_provserverconfiguration: tmp,
	}
	return
}
