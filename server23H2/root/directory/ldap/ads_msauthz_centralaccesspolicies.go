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

// ads_msauthz_centralaccesspolicies struct
type ads_msauthz_centralaccesspolicies struct {
	*ds_top
}

func Newads_msauthz_centralaccesspoliciesEx1(instance *cim.WmiInstance) (newInstance *ads_msauthz_centralaccesspolicies, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msauthz_centralaccesspolicies{
		ds_top: tmp,
	}
	return
}

func Newads_msauthz_centralaccesspoliciesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msauthz_centralaccesspolicies, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msauthz_centralaccesspolicies{
		ds_top: tmp,
	}
	return
}
