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

// ds_packageregistration struct
type ds_packageregistration struct {
	*ads_packageregistration
}

func Newds_packageregistrationEx1(instance *cim.WmiInstance) (newInstance *ds_packageregistration, err error) {
	tmp, err := Newads_packageregistrationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_packageregistration{
		ads_packageregistration: tmp,
	}
	return
}

func Newds_packageregistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_packageregistration, err error) {
	tmp, err := Newads_packageregistrationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_packageregistration{
		ads_packageregistration: tmp,
	}
	return
}
