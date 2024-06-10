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

// ds_device struct
type ds_device struct {
	*ads_device
}

func Newds_deviceEx1(instance *cim.WmiInstance) (newInstance *ds_device, err error) {
	tmp, err := Newads_deviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_device{
		ads_device: tmp,
	}
	return
}

func Newds_deviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_device, err error) {
	tmp, err := Newads_deviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_device{
		ads_device: tmp,
	}
	return
}
