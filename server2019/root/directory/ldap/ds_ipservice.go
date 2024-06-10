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

// ds_ipservice struct
type ds_ipservice struct {
	*ads_ipservice
}

func Newds_ipserviceEx1(instance *cim.WmiInstance) (newInstance *ds_ipservice, err error) {
	tmp, err := Newads_ipserviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ipservice{
		ads_ipservice: tmp,
	}
	return
}

func Newds_ipserviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ipservice, err error) {
	tmp, err := Newads_ipserviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ipservice{
		ads_ipservice: tmp,
	}
	return
}
