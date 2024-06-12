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

// ds_sitelinkbridge struct
type ds_sitelinkbridge struct {
	*ads_sitelinkbridge
}

func Newds_sitelinkbridgeEx1(instance *cim.WmiInstance) (newInstance *ds_sitelinkbridge, err error) {
	tmp, err := Newads_sitelinkbridgeEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_sitelinkbridge{
		ads_sitelinkbridge: tmp,
	}
	return
}

func Newds_sitelinkbridgeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_sitelinkbridge, err error) {
	tmp, err := Newads_sitelinkbridgeEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_sitelinkbridge{
		ads_sitelinkbridge: tmp,
	}
	return
}
