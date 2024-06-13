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

// ads_msds_managedserviceaccount struct
type ads_msds_managedserviceaccount struct {
	*ads_computer
}

func Newads_msds_managedserviceaccountEx1(instance *cim.WmiInstance) (newInstance *ads_msds_managedserviceaccount, err error) {
	tmp, err := Newads_computerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_managedserviceaccount{
		ads_computer: tmp,
	}
	return
}

func Newads_msds_managedserviceaccountEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_managedserviceaccount, err error) {
	tmp, err := Newads_computerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_managedserviceaccount{
		ads_computer: tmp,
	}
	return
}
