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

// ds_msauthz_centralaccessrules struct
type ds_msauthz_centralaccessrules struct {
	*ads_msauthz_centralaccessrules
}

func Newds_msauthz_centralaccessrulesEx1(instance *cim.WmiInstance) (newInstance *ds_msauthz_centralaccessrules, err error) {
	tmp, err := Newads_msauthz_centralaccessrulesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msauthz_centralaccessrules{
		ads_msauthz_centralaccessrules: tmp,
	}
	return
}

func Newds_msauthz_centralaccessrulesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msauthz_centralaccessrules, err error) {
	tmp, err := Newads_msauthz_centralaccessrulesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msauthz_centralaccessrules{
		ads_msauthz_centralaccessrules: tmp,
	}
	return
}
