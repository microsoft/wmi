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

// ds_remotemailrecipient struct
type ds_remotemailrecipient struct {
	*ads_remotemailrecipient
}

func Newds_remotemailrecipientEx1(instance *cim.WmiInstance) (newInstance *ds_remotemailrecipient, err error) {
	tmp, err := Newads_remotemailrecipientEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_remotemailrecipient{
		ads_remotemailrecipient: tmp,
	}
	return
}

func Newds_remotemailrecipientEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_remotemailrecipient, err error) {
	tmp, err := Newads_remotemailrecipientEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_remotemailrecipient{
		ads_remotemailrecipient: tmp,
	}
	return
}
