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

// ds_addresstemplate struct
type ds_addresstemplate struct {
	*ads_addresstemplate
}

func Newds_addresstemplateEx1(instance *cim.WmiInstance) (newInstance *ds_addresstemplate, err error) {
	tmp, err := Newads_addresstemplateEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_addresstemplate{
		ads_addresstemplate: tmp,
	}
	return
}

func Newds_addresstemplateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_addresstemplate, err error) {
	tmp, err := Newads_addresstemplateEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_addresstemplate{
		ads_addresstemplate: tmp,
	}
	return
}
