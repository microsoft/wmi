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

// ds_querypolicy struct
type ds_querypolicy struct {
	*ads_querypolicy
}

func Newds_querypolicyEx1(instance *cim.WmiInstance) (newInstance *ds_querypolicy, err error) {
	tmp, err := Newads_querypolicyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_querypolicy{
		ads_querypolicy: tmp,
	}
	return
}

func Newds_querypolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_querypolicy, err error) {
	tmp, err := Newads_querypolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_querypolicy{
		ads_querypolicy: tmp,
	}
	return
}
