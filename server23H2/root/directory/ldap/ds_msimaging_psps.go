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

// ds_msimaging_psps struct
type ds_msimaging_psps struct {
	*ads_msimaging_psps
}

func Newds_msimaging_pspsEx1(instance *cim.WmiInstance) (newInstance *ds_msimaging_psps, err error) {
	tmp, err := Newads_msimaging_pspsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msimaging_psps{
		ads_msimaging_psps: tmp,
	}
	return
}

func Newds_msimaging_pspsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msimaging_psps, err error) {
	tmp, err := Newads_msimaging_pspsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msimaging_psps{
		ads_msimaging_psps: tmp,
	}
	return
}
