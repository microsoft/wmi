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

// ds_addressbookcontainer struct
type ds_addressbookcontainer struct {
	*ads_addressbookcontainer
}

func Newds_addressbookcontainerEx1(instance *cim.WmiInstance) (newInstance *ds_addressbookcontainer, err error) {
	tmp, err := Newads_addressbookcontainerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_addressbookcontainer{
		ads_addressbookcontainer: tmp,
	}
	return
}

func Newds_addressbookcontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_addressbookcontainer, err error) {
	tmp, err := Newads_addressbookcontainerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_addressbookcontainer{
		ads_addressbookcontainer: tmp,
	}
	return
}
