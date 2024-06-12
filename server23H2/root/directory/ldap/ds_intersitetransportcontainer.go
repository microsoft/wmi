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

// ds_intersitetransportcontainer struct
type ds_intersitetransportcontainer struct {
	*ads_intersitetransportcontainer
}

func Newds_intersitetransportcontainerEx1(instance *cim.WmiInstance) (newInstance *ds_intersitetransportcontainer, err error) {
	tmp, err := Newads_intersitetransportcontainerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_intersitetransportcontainer{
		ads_intersitetransportcontainer: tmp,
	}
	return
}

func Newds_intersitetransportcontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_intersitetransportcontainer, err error) {
	tmp, err := Newads_intersitetransportcontainerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_intersitetransportcontainer{
		ads_intersitetransportcontainer: tmp,
	}
	return
}
