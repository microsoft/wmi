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

// ds_linktrackvolumetable struct
type ds_linktrackvolumetable struct {
	*ads_linktrackvolumetable
}

func Newds_linktrackvolumetableEx1(instance *cim.WmiInstance) (newInstance *ds_linktrackvolumetable, err error) {
	tmp, err := Newads_linktrackvolumetableEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_linktrackvolumetable{
		ads_linktrackvolumetable: tmp,
	}
	return
}

func Newds_linktrackvolumetableEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_linktrackvolumetable, err error) {
	tmp, err := Newads_linktrackvolumetableEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_linktrackvolumetable{
		ads_linktrackvolumetable: tmp,
	}
	return
}
