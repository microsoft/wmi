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

// ds_msds_claimtypes struct
type ds_msds_claimtypes struct {
	*ads_msds_claimtypes
}

func Newds_msds_claimtypesEx1(instance *cim.WmiInstance) (newInstance *ds_msds_claimtypes, err error) {
	tmp, err := Newads_msds_claimtypesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msds_claimtypes{
		ads_msds_claimtypes: tmp,
	}
	return
}

func Newds_msds_claimtypesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msds_claimtypes, err error) {
	tmp, err := Newads_msds_claimtypesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msds_claimtypes{
		ads_msds_claimtypes: tmp,
	}
	return
}
