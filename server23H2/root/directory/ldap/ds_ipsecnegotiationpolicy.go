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

// ds_ipsecnegotiationpolicy struct
type ds_ipsecnegotiationpolicy struct {
	*ads_ipsecnegotiationpolicy
}

func Newds_ipsecnegotiationpolicyEx1(instance *cim.WmiInstance) (newInstance *ds_ipsecnegotiationpolicy, err error) {
	tmp, err := Newads_ipsecnegotiationpolicyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ipsecnegotiationpolicy{
		ads_ipsecnegotiationpolicy: tmp,
	}
	return
}

func Newds_ipsecnegotiationpolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ipsecnegotiationpolicy, err error) {
	tmp, err := Newads_ipsecnegotiationpolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ipsecnegotiationpolicy{
		ads_ipsecnegotiationpolicy: tmp,
	}
	return
}
