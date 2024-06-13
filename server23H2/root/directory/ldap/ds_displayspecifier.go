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

// ds_displayspecifier struct
type ds_displayspecifier struct {
	*ads_displayspecifier
}

func Newds_displayspecifierEx1(instance *cim.WmiInstance) (newInstance *ds_displayspecifier, err error) {
	tmp, err := Newads_displayspecifierEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_displayspecifier{
		ads_displayspecifier: tmp,
	}
	return
}

func Newds_displayspecifierEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_displayspecifier, err error) {
	tmp, err := Newads_displayspecifierEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_displayspecifier{
		ads_displayspecifier: tmp,
	}
	return
}
