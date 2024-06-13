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

// ds_msauthz_centralaccesspolicy struct
type ds_msauthz_centralaccesspolicy struct {
	*ads_msauthz_centralaccesspolicy
}

func Newds_msauthz_centralaccesspolicyEx1(instance *cim.WmiInstance) (newInstance *ds_msauthz_centralaccesspolicy, err error) {
	tmp, err := Newads_msauthz_centralaccesspolicyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msauthz_centralaccesspolicy{
		ads_msauthz_centralaccesspolicy: tmp,
	}
	return
}

func Newds_msauthz_centralaccesspolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msauthz_centralaccesspolicy, err error) {
	tmp, err := Newads_msauthz_centralaccesspolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msauthz_centralaccesspolicy{
		ads_msauthz_centralaccesspolicy: tmp,
	}
	return
}
