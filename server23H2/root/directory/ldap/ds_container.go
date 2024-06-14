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

// ds_container struct
type ds_container struct {
	*ads_container
}

func Newds_containerEx1(instance *cim.WmiInstance) (newInstance *ds_container, err error) {
	tmp, err := Newads_containerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_container{
		ads_container: tmp,
	}
	return
}

func Newds_containerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_container, err error) {
	tmp, err := Newads_containerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_container{
		ads_container: tmp,
	}
	return
}
