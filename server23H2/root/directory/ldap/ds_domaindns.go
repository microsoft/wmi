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

// ds_domaindns struct
type ds_domaindns struct {
	*ads_domaindns
}

func Newds_domaindnsEx1(instance *cim.WmiInstance) (newInstance *ds_domaindns, err error) {
	tmp, err := Newads_domaindnsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_domaindns{
		ads_domaindns: tmp,
	}
	return
}

func Newds_domaindnsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_domaindns, err error) {
	tmp, err := Newads_domaindnsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_domaindns{
		ads_domaindns: tmp,
	}
	return
}
