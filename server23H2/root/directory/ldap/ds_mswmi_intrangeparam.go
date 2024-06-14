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

// ds_mswmi_intrangeparam struct
type ds_mswmi_intrangeparam struct {
	*ads_mswmi_intrangeparam
}

func Newds_mswmi_intrangeparamEx1(instance *cim.WmiInstance) (newInstance *ds_mswmi_intrangeparam, err error) {
	tmp, err := Newads_mswmi_intrangeparamEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_mswmi_intrangeparam{
		ads_mswmi_intrangeparam: tmp,
	}
	return
}

func Newds_mswmi_intrangeparamEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_mswmi_intrangeparam, err error) {
	tmp, err := Newads_mswmi_intrangeparamEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_mswmi_intrangeparam{
		ads_mswmi_intrangeparam: tmp,
	}
	return
}
