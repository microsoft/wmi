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

// ds_contact struct
type ds_contact struct {
	*ads_contact
}

func Newds_contactEx1(instance *cim.WmiInstance) (newInstance *ds_contact, err error) {
	tmp, err := Newads_contactEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_contact{
		ads_contact: tmp,
	}
	return
}

func Newds_contactEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_contact, err error) {
	tmp, err := Newads_contactEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_contact{
		ads_contact: tmp,
	}
	return
}
