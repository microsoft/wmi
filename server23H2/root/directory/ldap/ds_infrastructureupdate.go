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

// ds_infrastructureupdate struct
type ds_infrastructureupdate struct {
	*ads_infrastructureupdate
}

func Newds_infrastructureupdateEx1(instance *cim.WmiInstance) (newInstance *ds_infrastructureupdate, err error) {
	tmp, err := Newads_infrastructureupdateEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_infrastructureupdate{
		ads_infrastructureupdate: tmp,
	}
	return
}

func Newds_infrastructureupdateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_infrastructureupdate, err error) {
	tmp, err := Newads_infrastructureupdateEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_infrastructureupdate{
		ads_infrastructureupdate: tmp,
	}
	return
}
