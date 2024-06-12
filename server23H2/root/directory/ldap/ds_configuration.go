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

// ds_configuration struct
type ds_configuration struct {
	*ads_configuration
}

func Newds_configurationEx1(instance *cim.WmiInstance) (newInstance *ds_configuration, err error) {
	tmp, err := Newads_configurationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_configuration{
		ads_configuration: tmp,
	}
	return
}

func Newds_configurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_configuration, err error) {
	tmp, err := Newads_configurationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_configuration{
		ads_configuration: tmp,
	}
	return
}
