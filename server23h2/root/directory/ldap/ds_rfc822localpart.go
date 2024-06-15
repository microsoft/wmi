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

// ds_rfc822localpart struct
type ds_rfc822localpart struct {
	*ads_rfc822localpart
}

func Newds_rfc822localpartEx1(instance *cim.WmiInstance) (newInstance *ds_rfc822localpart, err error) {
	tmp, err := Newads_rfc822localpartEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_rfc822localpart{
		ads_rfc822localpart: tmp,
	}
	return
}

func Newds_rfc822localpartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_rfc822localpart, err error) {
	tmp, err := Newads_rfc822localpartEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_rfc822localpart{
		ads_rfc822localpart: tmp,
	}
	return
}
