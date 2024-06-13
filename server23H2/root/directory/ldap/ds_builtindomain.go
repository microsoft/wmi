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

// ds_builtindomain struct
type ds_builtindomain struct {
	*ads_builtindomain
}

func Newds_builtindomainEx1(instance *cim.WmiInstance) (newInstance *ds_builtindomain, err error) {
	tmp, err := Newads_builtindomainEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_builtindomain{
		ads_builtindomain: tmp,
	}
	return
}

func Newds_builtindomainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_builtindomain, err error) {
	tmp, err := Newads_builtindomainEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_builtindomain{
		ads_builtindomain: tmp,
	}
	return
}
