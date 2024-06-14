// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_filelinktrackingentry struct
type ds_filelinktrackingentry struct {
	*ads_filelinktrackingentry
}

func Newds_filelinktrackingentryEx1(instance *cim.WmiInstance) (newInstance *ds_filelinktrackingentry, err error) {
	tmp, err := Newads_filelinktrackingentryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_filelinktrackingentry{
		ads_filelinktrackingentry: tmp,
	}
	return
}

func Newds_filelinktrackingentryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_filelinktrackingentry, err error) {
	tmp, err := Newads_filelinktrackingentryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_filelinktrackingentry{
		ads_filelinktrackingentry: tmp,
	}
	return
}
