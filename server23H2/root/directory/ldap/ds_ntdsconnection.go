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

// ds_ntdsconnection struct
type ds_ntdsconnection struct {
	*ads_ntdsconnection
}

func Newds_ntdsconnectionEx1(instance *cim.WmiInstance) (newInstance *ds_ntdsconnection, err error) {
	tmp, err := Newads_ntdsconnectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ntdsconnection{
		ads_ntdsconnection: tmp,
	}
	return
}

func Newds_ntdsconnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ntdsconnection, err error) {
	tmp, err := Newads_ntdsconnectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ntdsconnection{
		ads_ntdsconnection: tmp,
	}
	return
}
