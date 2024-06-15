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

// ds_ntfrssubscriber struct
type ds_ntfrssubscriber struct {
	*ads_ntfrssubscriber
}

func Newds_ntfrssubscriberEx1(instance *cim.WmiInstance) (newInstance *ds_ntfrssubscriber, err error) {
	tmp, err := Newads_ntfrssubscriberEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ntfrssubscriber{
		ads_ntfrssubscriber: tmp,
	}
	return
}

func Newds_ntfrssubscriberEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ntfrssubscriber, err error) {
	tmp, err := Newads_ntfrssubscriberEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ntfrssubscriber{
		ads_ntfrssubscriber: tmp,
	}
	return
}
