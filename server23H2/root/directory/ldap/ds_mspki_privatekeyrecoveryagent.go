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

// ds_mspki_privatekeyrecoveryagent struct
type ds_mspki_privatekeyrecoveryagent struct {
	*ads_mspki_privatekeyrecoveryagent
}

func Newds_mspki_privatekeyrecoveryagentEx1(instance *cim.WmiInstance) (newInstance *ds_mspki_privatekeyrecoveryagent, err error) {
	tmp, err := Newads_mspki_privatekeyrecoveryagentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_mspki_privatekeyrecoveryagent{
		ads_mspki_privatekeyrecoveryagent: tmp,
	}
	return
}

func Newds_mspki_privatekeyrecoveryagentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_mspki_privatekeyrecoveryagent, err error) {
	tmp, err := Newads_mspki_privatekeyrecoveryagentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_mspki_privatekeyrecoveryagent{
		ads_mspki_privatekeyrecoveryagent: tmp,
	}
	return
}