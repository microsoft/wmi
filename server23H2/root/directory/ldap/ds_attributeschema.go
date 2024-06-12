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

// ds_attributeschema struct
type ds_attributeschema struct {
	*ads_attributeschema
}

func Newds_attributeschemaEx1(instance *cim.WmiInstance) (newInstance *ds_attributeschema, err error) {
	tmp, err := Newads_attributeschemaEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_attributeschema{
		ads_attributeschema: tmp,
	}
	return
}

func Newds_attributeschemaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_attributeschema, err error) {
	tmp, err := Newads_attributeschemaEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_attributeschema{
		ads_attributeschema: tmp,
	}
	return
}
