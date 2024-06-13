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

// ds_filelinktracking struct
type ds_filelinktracking struct {
	*ads_filelinktracking
}

func Newds_filelinktrackingEx1(instance *cim.WmiInstance) (newInstance *ds_filelinktracking, err error) {
	tmp, err := Newads_filelinktrackingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_filelinktracking{
		ads_filelinktracking: tmp,
	}
	return
}

func Newds_filelinktrackingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_filelinktracking, err error) {
	tmp, err := Newads_filelinktrackingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_filelinktracking{
		ads_filelinktracking: tmp,
	}
	return
}
