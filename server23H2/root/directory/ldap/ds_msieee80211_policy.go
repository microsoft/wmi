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

// ds_msieee80211_policy struct
type ds_msieee80211_policy struct {
	*ads_msieee80211_policy
}

func Newds_msieee80211_policyEx1(instance *cim.WmiInstance) (newInstance *ds_msieee80211_policy, err error) {
	tmp, err := Newads_msieee80211_policyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msieee80211_policy{
		ads_msieee80211_policy: tmp,
	}
	return
}

func Newds_msieee80211_policyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msieee80211_policy, err error) {
	tmp, err := Newads_msieee80211_policyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msieee80211_policy{
		ads_msieee80211_policy: tmp,
	}
	return
}
