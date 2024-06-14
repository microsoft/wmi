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

// ds_msds_authnpolicysilo struct
type ds_msds_authnpolicysilo struct {
	*ads_msds_authnpolicysilo
}

func Newds_msds_authnpolicysiloEx1(instance *cim.WmiInstance) (newInstance *ds_msds_authnpolicysilo, err error) {
	tmp, err := Newads_msds_authnpolicysiloEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msds_authnpolicysilo{
		ads_msds_authnpolicysilo: tmp,
	}
	return
}

func Newds_msds_authnpolicysiloEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msds_authnpolicysilo, err error) {
	tmp, err := Newads_msds_authnpolicysiloEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msds_authnpolicysilo{
		ads_msds_authnpolicysilo: tmp,
	}
	return
}