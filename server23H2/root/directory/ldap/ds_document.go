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

// ds_document struct
type ds_document struct {
	*ads_document
}

func Newds_documentEx1(instance *cim.WmiInstance) (newInstance *ds_document, err error) {
	tmp, err := Newads_documentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_document{
		ads_document: tmp,
	}
	return
}

func Newds_documentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_document, err error) {
	tmp, err := Newads_documentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_document{
		ads_document: tmp,
	}
	return
}
