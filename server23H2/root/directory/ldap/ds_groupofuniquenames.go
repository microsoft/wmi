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

// ds_groupofuniquenames struct
type ds_groupofuniquenames struct {
	*ads_groupofuniquenames
}

func Newds_groupofuniquenamesEx1(instance *cim.WmiInstance) (newInstance *ds_groupofuniquenames, err error) {
	tmp, err := Newads_groupofuniquenamesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_groupofuniquenames{
		ads_groupofuniquenames: tmp,
	}
	return
}

func Newds_groupofuniquenamesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_groupofuniquenames, err error) {
	tmp, err := Newads_groupofuniquenamesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_groupofuniquenames{
		ads_groupofuniquenames: tmp,
	}
	return
}
