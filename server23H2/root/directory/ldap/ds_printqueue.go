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

// ds_printqueue struct
type ds_printqueue struct {
	*ads_printqueue
}

func Newds_printqueueEx1(instance *cim.WmiInstance) (newInstance *ds_printqueue, err error) {
	tmp, err := Newads_printqueueEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_printqueue{
		ads_printqueue: tmp,
	}
	return
}

func Newds_printqueueEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_printqueue, err error) {
	tmp, err := Newads_printqueueEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_printqueue{
		ads_printqueue: tmp,
	}
	return
}
