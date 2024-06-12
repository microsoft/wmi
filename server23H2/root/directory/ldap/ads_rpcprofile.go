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

// ads_rpcprofile struct
type ads_rpcprofile struct {
	*ds_rpcentry
}

func Newads_rpcprofileEx1(instance *cim.WmiInstance) (newInstance *ads_rpcprofile, err error) {
	tmp, err := Newds_rpcentryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_rpcprofile{
		ds_rpcentry: tmp,
	}
	return
}

func Newads_rpcprofileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_rpcprofile, err error) {
	tmp, err := Newds_rpcentryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_rpcprofile{
		ds_rpcentry: tmp,
	}
	return
}
