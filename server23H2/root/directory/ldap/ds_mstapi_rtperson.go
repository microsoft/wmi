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

// ds_mstapi_rtperson struct
type ds_mstapi_rtperson struct {
	*ads_mstapi_rtperson
}

func Newds_mstapi_rtpersonEx1(instance *cim.WmiInstance) (newInstance *ds_mstapi_rtperson, err error) {
	tmp, err := Newads_mstapi_rtpersonEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_mstapi_rtperson{
		ads_mstapi_rtperson: tmp,
	}
	return
}

func Newds_mstapi_rtpersonEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_mstapi_rtperson, err error) {
	tmp, err := Newads_mstapi_rtpersonEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_mstapi_rtperson{
		ads_mstapi_rtperson: tmp,
	}
	return
}
