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

// ds_intellimirrorscp struct
type ds_intellimirrorscp struct {
	*ads_intellimirrorscp
}

func Newds_intellimirrorscpEx1(instance *cim.WmiInstance) (newInstance *ds_intellimirrorscp, err error) {
	tmp, err := Newads_intellimirrorscpEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_intellimirrorscp{
		ads_intellimirrorscp: tmp,
	}
	return
}

func Newds_intellimirrorscpEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_intellimirrorscp, err error) {
	tmp, err := Newads_intellimirrorscpEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_intellimirrorscp{
		ads_intellimirrorscp: tmp,
	}
	return
}
