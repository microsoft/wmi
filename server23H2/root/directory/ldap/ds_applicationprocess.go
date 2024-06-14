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

// ds_applicationprocess struct
type ds_applicationprocess struct {
	*ads_applicationprocess
}

func Newds_applicationprocessEx1(instance *cim.WmiInstance) (newInstance *ds_applicationprocess, err error) {
	tmp, err := Newads_applicationprocessEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_applicationprocess{
		ads_applicationprocess: tmp,
	}
	return
}

func Newds_applicationprocessEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_applicationprocess, err error) {
	tmp, err := Newads_applicationprocessEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_applicationprocess{
		ads_applicationprocess: tmp,
	}
	return
}
