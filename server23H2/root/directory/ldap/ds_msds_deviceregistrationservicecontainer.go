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

// ds_msds_deviceregistrationservicecontainer struct
type ds_msds_deviceregistrationservicecontainer struct {
	*ads_msds_deviceregistrationservicecontainer
}

func Newds_msds_deviceregistrationservicecontainerEx1(instance *cim.WmiInstance) (newInstance *ds_msds_deviceregistrationservicecontainer, err error) {
	tmp, err := Newads_msds_deviceregistrationservicecontainerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msds_deviceregistrationservicecontainer{
		ads_msds_deviceregistrationservicecontainer: tmp,
	}
	return
}

func Newds_msds_deviceregistrationservicecontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msds_deviceregistrationservicecontainer, err error) {
	tmp, err := Newads_msds_deviceregistrationservicecontainerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msds_deviceregistrationservicecontainer{
		ads_msds_deviceregistrationservicecontainer: tmp,
	}
	return
}
