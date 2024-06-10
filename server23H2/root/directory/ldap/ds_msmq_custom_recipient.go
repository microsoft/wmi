// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_msmq_custom_recipient struct
type ds_msmq_custom_recipient struct {
	*ads_msmq_custom_recipient
}

func Newds_msmq_custom_recipientEx1(instance *cim.WmiInstance) (newInstance *ds_msmq_custom_recipient, err error) {
	tmp, err := Newads_msmq_custom_recipientEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msmq_custom_recipient{
		ads_msmq_custom_recipient: tmp,
	}
	return
}

func Newds_msmq_custom_recipientEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msmq_custom_recipient, err error) {
	tmp, err := Newads_msmq_custom_recipientEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msmq_custom_recipient{
		ads_msmq_custom_recipient: tmp,
	}
	return
}
