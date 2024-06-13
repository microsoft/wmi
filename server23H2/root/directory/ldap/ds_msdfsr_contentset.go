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

// ds_msdfsr_contentset struct
type ds_msdfsr_contentset struct {
	*ads_msdfsr_contentset
}

func Newds_msdfsr_contentsetEx1(instance *cim.WmiInstance) (newInstance *ds_msdfsr_contentset, err error) {
	tmp, err := Newads_msdfsr_contentsetEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_contentset{
		ads_msdfsr_contentset: tmp,
	}
	return
}

func Newds_msdfsr_contentsetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msdfsr_contentset, err error) {
	tmp, err := Newads_msdfsr_contentsetEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_contentset{
		ads_msdfsr_contentset: tmp,
	}
	return
}
