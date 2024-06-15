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

// ads_mswmi_mergeablepolicytemplate struct
type ads_mswmi_mergeablepolicytemplate struct {
	*ads_mswmi_policytemplate
}

func Newads_mswmi_mergeablepolicytemplateEx1(instance *cim.WmiInstance) (newInstance *ads_mswmi_mergeablepolicytemplate, err error) {
	tmp, err := Newads_mswmi_policytemplateEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_mergeablepolicytemplate{
		ads_mswmi_policytemplate: tmp,
	}
	return
}

func Newads_mswmi_mergeablepolicytemplateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mswmi_mergeablepolicytemplate, err error) {
	tmp, err := Newads_mswmi_policytemplateEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_mergeablepolicytemplate{
		ads_mswmi_policytemplate: tmp,
	}
	return
}