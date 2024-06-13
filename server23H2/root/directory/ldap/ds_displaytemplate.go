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

// ds_displaytemplate struct
type ds_displaytemplate struct {
	*ads_displaytemplate
}

func Newds_displaytemplateEx1(instance *cim.WmiInstance) (newInstance *ds_displaytemplate, err error) {
	tmp, err := Newads_displaytemplateEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_displaytemplate{
		ads_displaytemplate: tmp,
	}
	return
}

func Newds_displaytemplateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_displaytemplate, err error) {
	tmp, err := Newads_displaytemplateEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_displaytemplate{
		ads_displaytemplate: tmp,
	}
	return
}
