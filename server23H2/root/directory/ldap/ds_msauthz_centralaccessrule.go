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

// ds_msauthz_centralaccessrule struct
type ds_msauthz_centralaccessrule struct {
	*ads_msauthz_centralaccessrule
}

func Newds_msauthz_centralaccessruleEx1(instance *cim.WmiInstance) (newInstance *ds_msauthz_centralaccessrule, err error) {
	tmp, err := Newads_msauthz_centralaccessruleEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msauthz_centralaccessrule{
		ads_msauthz_centralaccessrule: tmp,
	}
	return
}

func Newds_msauthz_centralaccessruleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msauthz_centralaccessrule, err error) {
	tmp, err := Newads_msauthz_centralaccessruleEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msauthz_centralaccessrule{
		ads_msauthz_centralaccessrule: tmp,
	}
	return
}
