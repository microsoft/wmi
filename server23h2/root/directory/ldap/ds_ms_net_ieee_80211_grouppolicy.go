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

// ds_ms_net_ieee_80211_grouppolicy struct
type ds_ms_net_ieee_80211_grouppolicy struct {
	*ads_ms_net_ieee_80211_grouppolicy
}

func Newds_ms_net_ieee_80211_grouppolicyEx1(instance *cim.WmiInstance) (newInstance *ds_ms_net_ieee_80211_grouppolicy, err error) {
	tmp, err := Newads_ms_net_ieee_80211_grouppolicyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ms_net_ieee_80211_grouppolicy{
		ads_ms_net_ieee_80211_grouppolicy: tmp,
	}
	return
}

func Newds_ms_net_ieee_80211_grouppolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ms_net_ieee_80211_grouppolicy, err error) {
	tmp, err := Newads_ms_net_ieee_80211_grouppolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ms_net_ieee_80211_grouppolicy{
		ads_ms_net_ieee_80211_grouppolicy: tmp,
	}
	return
}
