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

// ds_crldistributionpoint struct
type ds_crldistributionpoint struct {
	*ads_crldistributionpoint
}

func Newds_crldistributionpointEx1(instance *cim.WmiInstance) (newInstance *ds_crldistributionpoint, err error) {
	tmp, err := Newads_crldistributionpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_crldistributionpoint{
		ads_crldistributionpoint: tmp,
	}
	return
}

func Newds_crldistributionpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_crldistributionpoint, err error) {
	tmp, err := Newads_crldistributionpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_crldistributionpoint{
		ads_crldistributionpoint: tmp,
	}
	return
}
