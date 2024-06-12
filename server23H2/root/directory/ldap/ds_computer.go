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

// ds_computer struct
type ds_computer struct {
	*ads_computer
}

func Newds_computerEx1(instance *cim.WmiInstance) (newInstance *ds_computer, err error) {
	tmp, err := Newads_computerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_computer{
		ads_computer: tmp,
	}
	return
}

func Newds_computerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_computer, err error) {
	tmp, err := Newads_computerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_computer{
		ads_computer: tmp,
	}
	return
}
