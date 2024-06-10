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

// ds_rpcprofile struct
type ds_rpcprofile struct {
	*ads_rpcprofile
}

func Newds_rpcprofileEx1(instance *cim.WmiInstance) (newInstance *ds_rpcprofile, err error) {
	tmp, err := Newads_rpcprofileEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_rpcprofile{
		ads_rpcprofile: tmp,
	}
	return
}

func Newds_rpcprofileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_rpcprofile, err error) {
	tmp, err := Newads_rpcprofileEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_rpcprofile{
		ads_rpcprofile: tmp,
	}
	return
}
