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

// ds_crossrefcontainer struct
type ds_crossrefcontainer struct {
	*ads_crossrefcontainer
}

func Newds_crossrefcontainerEx1(instance *cim.WmiInstance) (newInstance *ds_crossrefcontainer, err error) {
	tmp, err := Newads_crossrefcontainerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_crossrefcontainer{
		ads_crossrefcontainer: tmp,
	}
	return
}

func Newds_crossrefcontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_crossrefcontainer, err error) {
	tmp, err := Newads_crossrefcontainerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_crossrefcontainer{
		ads_crossrefcontainer: tmp,
	}
	return
}
