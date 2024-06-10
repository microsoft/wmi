// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_applicationversion struct
type ds_applicationversion struct {
	*ads_applicationversion
}

func Newds_applicationversionEx1(instance *cim.WmiInstance) (newInstance *ds_applicationversion, err error) {
	tmp, err := Newads_applicationversionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_applicationversion{
		ads_applicationversion: tmp,
	}
	return
}

func Newds_applicationversionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_applicationversion, err error) {
	tmp, err := Newads_applicationversionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_applicationversion{
		ads_applicationversion: tmp,
	}
	return
}
