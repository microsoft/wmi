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

// ds_oncrpc struct
type ds_oncrpc struct {
	*ads_oncrpc
}

func Newds_oncrpcEx1(instance *cim.WmiInstance) (newInstance *ds_oncrpc, err error) {
	tmp, err := Newads_oncrpcEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_oncrpc{
		ads_oncrpc: tmp,
	}
	return
}

func Newds_oncrpcEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_oncrpc, err error) {
	tmp, err := Newads_oncrpcEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_oncrpc{
		ads_oncrpc: tmp,
	}
	return
}
