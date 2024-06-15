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

// ds_ntdsservice struct
type ds_ntdsservice struct {
	*ads_ntdsservice
}

func Newds_ntdsserviceEx1(instance *cim.WmiInstance) (newInstance *ds_ntdsservice, err error) {
	tmp, err := Newads_ntdsserviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ntdsservice{
		ads_ntdsservice: tmp,
	}
	return
}

func Newds_ntdsserviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ntdsservice, err error) {
	tmp, err := Newads_ntdsserviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ntdsservice{
		ads_ntdsservice: tmp,
	}
	return
}
