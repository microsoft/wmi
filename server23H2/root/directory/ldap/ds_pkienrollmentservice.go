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

// ds_pkienrollmentservice struct
type ds_pkienrollmentservice struct {
	*ads_pkienrollmentservice
}

func Newds_pkienrollmentserviceEx1(instance *cim.WmiInstance) (newInstance *ds_pkienrollmentservice, err error) {
	tmp, err := Newads_pkienrollmentserviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_pkienrollmentservice{
		ads_pkienrollmentservice: tmp,
	}
	return
}

func Newds_pkienrollmentserviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_pkienrollmentservice, err error) {
	tmp, err := Newads_pkienrollmentserviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_pkienrollmentservice{
		ads_pkienrollmentservice: tmp,
	}
	return
}
