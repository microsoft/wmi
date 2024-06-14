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

// ds_dnszonescope struct
type ds_dnszonescope struct {
	*ads_dnszonescope
}

func Newds_dnszonescopeEx1(instance *cim.WmiInstance) (newInstance *ds_dnszonescope, err error) {
	tmp, err := Newads_dnszonescopeEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_dnszonescope{
		ads_dnszonescope: tmp,
	}
	return
}

func Newds_dnszonescopeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_dnszonescope, err error) {
	tmp, err := Newads_dnszonescopeEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_dnszonescope{
		ads_dnszonescope: tmp,
	}
	return
}
