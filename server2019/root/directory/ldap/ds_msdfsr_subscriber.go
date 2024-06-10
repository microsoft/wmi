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

// ds_msdfsr_subscriber struct
type ds_msdfsr_subscriber struct {
	*ads_msdfsr_subscriber
}

func Newds_msdfsr_subscriberEx1(instance *cim.WmiInstance) (newInstance *ds_msdfsr_subscriber, err error) {
	tmp, err := Newads_msdfsr_subscriberEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_subscriber{
		ads_msdfsr_subscriber: tmp,
	}
	return
}

func Newds_msdfsr_subscriberEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msdfsr_subscriber, err error) {
	tmp, err := Newads_msdfsr_subscriberEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_subscriber{
		ads_msdfsr_subscriber: tmp,
	}
	return
}
