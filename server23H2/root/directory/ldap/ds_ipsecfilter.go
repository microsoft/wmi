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

// ds_ipsecfilter struct
type ds_ipsecfilter struct {
	*ads_ipsecfilter
}

func Newds_ipsecfilterEx1(instance *cim.WmiInstance) (newInstance *ds_ipsecfilter, err error) {
	tmp, err := Newads_ipsecfilterEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ipsecfilter{
		ads_ipsecfilter: tmp,
	}
	return
}

func Newds_ipsecfilterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ipsecfilter, err error) {
	tmp, err := Newads_ipsecfilterEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ipsecfilter{
		ads_ipsecfilter: tmp,
	}
	return
}
