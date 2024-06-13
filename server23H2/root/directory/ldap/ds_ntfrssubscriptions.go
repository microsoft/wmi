// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_ntfrssubscriptions struct
type ds_ntfrssubscriptions struct {
	*ads_ntfrssubscriptions
}

func Newds_ntfrssubscriptionsEx1(instance *cim.WmiInstance) (newInstance *ds_ntfrssubscriptions, err error) {
	tmp, err := Newads_ntfrssubscriptionsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ntfrssubscriptions{
		ads_ntfrssubscriptions: tmp,
	}
	return
}

func Newds_ntfrssubscriptionsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ntfrssubscriptions, err error) {
	tmp, err := Newads_ntfrssubscriptionsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ntfrssubscriptions{
		ads_ntfrssubscriptions: tmp,
	}
	return
}
