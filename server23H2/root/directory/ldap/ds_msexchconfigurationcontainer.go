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

// ds_msexchconfigurationcontainer struct
type ds_msexchconfigurationcontainer struct {
	*ads_msexchconfigurationcontainer
}

func Newds_msexchconfigurationcontainerEx1(instance *cim.WmiInstance) (newInstance *ds_msexchconfigurationcontainer, err error) {
	tmp, err := Newads_msexchconfigurationcontainerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msexchconfigurationcontainer{
		ads_msexchconfigurationcontainer: tmp,
	}
	return
}

func Newds_msexchconfigurationcontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msexchconfigurationcontainer, err error) {
	tmp, err := Newads_msexchconfigurationcontainerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msexchconfigurationcontainer{
		ads_msexchconfigurationcontainer: tmp,
	}
	return
}
