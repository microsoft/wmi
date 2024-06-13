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

// ds_mspki_key_recovery_agent struct
type ds_mspki_key_recovery_agent struct {
	*ads_mspki_key_recovery_agent
}

func Newds_mspki_key_recovery_agentEx1(instance *cim.WmiInstance) (newInstance *ds_mspki_key_recovery_agent, err error) {
	tmp, err := Newads_mspki_key_recovery_agentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_mspki_key_recovery_agent{
		ads_mspki_key_recovery_agent: tmp,
	}
	return
}

func Newds_mspki_key_recovery_agentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_mspki_key_recovery_agent, err error) {
	tmp, err := Newads_mspki_key_recovery_agentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_mspki_key_recovery_agent{
		ads_mspki_key_recovery_agent: tmp,
	}
	return
}
