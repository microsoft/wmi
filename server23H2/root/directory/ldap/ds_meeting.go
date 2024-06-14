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

// ds_meeting struct
type ds_meeting struct {
	*ads_meeting
}

func Newds_meetingEx1(instance *cim.WmiInstance) (newInstance *ds_meeting, err error) {
	tmp, err := Newads_meetingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_meeting{
		ads_meeting: tmp,
	}
	return
}

func Newds_meetingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_meeting, err error) {
	tmp, err := Newads_meetingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_meeting{
		ads_meeting: tmp,
	}
	return
}
