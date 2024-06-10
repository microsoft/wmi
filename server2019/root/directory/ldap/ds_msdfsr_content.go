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

// ds_msdfsr_content struct
type ds_msdfsr_content struct {
	*ads_msdfsr_content
}

func Newds_msdfsr_contentEx1(instance *cim.WmiInstance) (newInstance *ds_msdfsr_content, err error) {
	tmp, err := Newads_msdfsr_contentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_content{
		ads_msdfsr_content: tmp,
	}
	return
}

func Newds_msdfsr_contentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msdfsr_content, err error) {
	tmp, err := Newads_msdfsr_contentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msdfsr_content{
		ads_msdfsr_content: tmp,
	}
	return
}
