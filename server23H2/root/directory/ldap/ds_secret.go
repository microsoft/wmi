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

// ds_secret struct
type ds_secret struct {
	*ads_secret
}

func Newds_secretEx1(instance *cim.WmiInstance) (newInstance *ds_secret, err error) {
	tmp, err := Newads_secretEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_secret{
		ads_secret: tmp,
	}
	return
}

func Newds_secretEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_secret, err error) {
	tmp, err := Newads_secretEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_secret{
		ads_secret: tmp,
	}
	return
}
