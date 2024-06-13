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

// ds_dhcpclass struct
type ds_dhcpclass struct {
	*ads_dhcpclass
}

func Newds_dhcpclassEx1(instance *cim.WmiInstance) (newInstance *ds_dhcpclass, err error) {
	tmp, err := Newads_dhcpclassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_dhcpclass{
		ads_dhcpclass: tmp,
	}
	return
}

func Newds_dhcpclassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_dhcpclass, err error) {
	tmp, err := Newads_dhcpclassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_dhcpclass{
		ads_dhcpclass: tmp,
	}
	return
}
