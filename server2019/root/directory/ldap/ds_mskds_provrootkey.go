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

// ds_mskds_provrootkey struct
type ds_mskds_provrootkey struct {
	*ads_mskds_provrootkey
}

func Newds_mskds_provrootkeyEx1(instance *cim.WmiInstance) (newInstance *ds_mskds_provrootkey, err error) {
	tmp, err := Newads_mskds_provrootkeyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_mskds_provrootkey{
		ads_mskds_provrootkey: tmp,
	}
	return
}

func Newds_mskds_provrootkeyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_mskds_provrootkey, err error) {
	tmp, err := Newads_mskds_provrootkeyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_mskds_provrootkey{
		ads_mskds_provrootkey: tmp,
	}
	return
}
