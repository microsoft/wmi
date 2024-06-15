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

// ads_mspki_key_recovery_agent struct
type ads_mspki_key_recovery_agent struct {
	*ads_user
}

func Newads_mspki_key_recovery_agentEx1(instance *cim.WmiInstance) (newInstance *ads_mspki_key_recovery_agent, err error) {
	tmp, err := Newads_userEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mspki_key_recovery_agent{
		ads_user: tmp,
	}
	return
}

func Newads_mspki_key_recovery_agentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mspki_key_recovery_agent, err error) {
	tmp, err := Newads_userEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mspki_key_recovery_agent{
		ads_user: tmp,
	}
	return
}
