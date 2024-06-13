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

// ds_msspp_activationobjectscontainer struct
type ds_msspp_activationobjectscontainer struct {
	*ads_msspp_activationobjectscontainer
}

func Newds_msspp_activationobjectscontainerEx1(instance *cim.WmiInstance) (newInstance *ds_msspp_activationobjectscontainer, err error) {
	tmp, err := Newads_msspp_activationobjectscontainerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msspp_activationobjectscontainer{
		ads_msspp_activationobjectscontainer: tmp,
	}
	return
}

func Newds_msspp_activationobjectscontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msspp_activationobjectscontainer, err error) {
	tmp, err := Newads_msspp_activationobjectscontainerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msspp_activationobjectscontainer{
		ads_msspp_activationobjectscontainer: tmp,
	}
	return
}
