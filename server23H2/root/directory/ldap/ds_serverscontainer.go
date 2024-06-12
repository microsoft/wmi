// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_serverscontainer struct
type ds_serverscontainer struct {
	*ads_serverscontainer
}

func Newds_serverscontainerEx1(instance *cim.WmiInstance) (newInstance *ds_serverscontainer, err error) {
	tmp, err := Newads_serverscontainerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_serverscontainer{
		ads_serverscontainer: tmp,
	}
	return
}

func Newds_serverscontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_serverscontainer, err error) {
	tmp, err := Newads_serverscontainerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_serverscontainer{
		ads_serverscontainer: tmp,
	}
	return
}
