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

// ds_mswmi_stringsetparam struct
type ds_mswmi_stringsetparam struct {
	*ads_mswmi_stringsetparam
}

func Newds_mswmi_stringsetparamEx1(instance *cim.WmiInstance) (newInstance *ds_mswmi_stringsetparam, err error) {
	tmp, err := Newads_mswmi_stringsetparamEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_mswmi_stringsetparam{
		ads_mswmi_stringsetparam: tmp,
	}
	return
}

func Newds_mswmi_stringsetparamEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_mswmi_stringsetparam, err error) {
	tmp, err := Newads_mswmi_stringsetparamEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_mswmi_stringsetparam{
		ads_mswmi_stringsetparam: tmp,
	}
	return
}
