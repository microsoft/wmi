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

// ds_dnszone struct
type ds_dnszone struct {
	*ads_dnszone
}

func Newds_dnszoneEx1(instance *cim.WmiInstance) (newInstance *ds_dnszone, err error) {
	tmp, err := Newads_dnszoneEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_dnszone{
		ads_dnszone: tmp,
	}
	return
}

func Newds_dnszoneEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_dnszone, err error) {
	tmp, err := Newads_dnszoneEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_dnszone{
		ads_dnszone: tmp,
	}
	return
}
