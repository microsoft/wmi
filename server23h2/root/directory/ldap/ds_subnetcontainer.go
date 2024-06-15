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

// ds_subnetcontainer struct
type ds_subnetcontainer struct {
	*ads_subnetcontainer
}

func Newds_subnetcontainerEx1(instance *cim.WmiInstance) (newInstance *ds_subnetcontainer, err error) {
	tmp, err := Newads_subnetcontainerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_subnetcontainer{
		ads_subnetcontainer: tmp,
	}
	return
}

func Newds_subnetcontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_subnetcontainer, err error) {
	tmp, err := Newads_subnetcontainerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_subnetcontainer{
		ads_subnetcontainer: tmp,
	}
	return
}
