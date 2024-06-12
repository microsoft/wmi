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

// ds_applicationentity struct
type ds_applicationentity struct {
	*ads_applicationentity
}

func Newds_applicationentityEx1(instance *cim.WmiInstance) (newInstance *ds_applicationentity, err error) {
	tmp, err := Newads_applicationentityEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_applicationentity{
		ads_applicationentity: tmp,
	}
	return
}

func Newds_applicationentityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_applicationentity, err error) {
	tmp, err := Newads_applicationentityEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_applicationentity{
		ads_applicationentity: tmp,
	}
	return
}
