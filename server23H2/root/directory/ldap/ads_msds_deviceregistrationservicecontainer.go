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

// ads_msds_deviceregistrationservicecontainer struct
type ads_msds_deviceregistrationservicecontainer struct {
	*ds_top
}

func Newads_msds_deviceregistrationservicecontainerEx1(instance *cim.WmiInstance) (newInstance *ads_msds_deviceregistrationservicecontainer, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_deviceregistrationservicecontainer{
		ds_top: tmp,
	}
	return
}

func Newads_msds_deviceregistrationservicecontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_deviceregistrationservicecontainer, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_deviceregistrationservicecontainer{
		ds_top: tmp,
	}
	return
}