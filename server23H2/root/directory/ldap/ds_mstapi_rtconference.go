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

// ds_mstapi_rtconference struct
type ds_mstapi_rtconference struct {
	*ads_mstapi_rtconference
}

func Newds_mstapi_rtconferenceEx1(instance *cim.WmiInstance) (newInstance *ds_mstapi_rtconference, err error) {
	tmp, err := Newads_mstapi_rtconferenceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_mstapi_rtconference{
		ads_mstapi_rtconference: tmp,
	}
	return
}

func Newds_mstapi_rtconferenceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_mstapi_rtconference, err error) {
	tmp, err := Newads_mstapi_rtconferenceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_mstapi_rtconference{
		ads_mstapi_rtconference: tmp,
	}
	return
}
