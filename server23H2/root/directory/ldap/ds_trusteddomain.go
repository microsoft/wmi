// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_trusteddomain struct
type ds_trusteddomain struct {
	*ads_trusteddomain
}

func Newds_trusteddomainEx1(instance *cim.WmiInstance) (newInstance *ds_trusteddomain, err error) {
	tmp, err := Newads_trusteddomainEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_trusteddomain{
		ads_trusteddomain: tmp,
	}
	return
}

func Newds_trusteddomainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_trusteddomain, err error) {
	tmp, err := Newads_trusteddomainEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_trusteddomain{
		ads_trusteddomain: tmp,
	}
	return
}
