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

// ds_mswmi_unknownrangeparam struct
type ds_mswmi_unknownrangeparam struct {
	*ads_mswmi_unknownrangeparam
}

func Newds_mswmi_unknownrangeparamEx1(instance *cim.WmiInstance) (newInstance *ds_mswmi_unknownrangeparam, err error) {
	tmp, err := Newads_mswmi_unknownrangeparamEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_mswmi_unknownrangeparam{
		ads_mswmi_unknownrangeparam: tmp,
	}
	return
}

func Newds_mswmi_unknownrangeparamEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_mswmi_unknownrangeparam, err error) {
	tmp, err := Newads_mswmi_unknownrangeparamEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_mswmi_unknownrangeparam{
		ads_mswmi_unknownrangeparam: tmp,
	}
	return
}
