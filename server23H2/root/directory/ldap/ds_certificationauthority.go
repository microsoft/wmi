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

// ds_certificationauthority struct
type ds_certificationauthority struct {
	*ads_certificationauthority
}

func Newds_certificationauthorityEx1(instance *cim.WmiInstance) (newInstance *ds_certificationauthority, err error) {
	tmp, err := Newads_certificationauthorityEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_certificationauthority{
		ads_certificationauthority: tmp,
	}
	return
}

func Newds_certificationauthorityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_certificationauthority, err error) {
	tmp, err := Newads_certificationauthorityEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_certificationauthority{
		ads_certificationauthority: tmp,
	}
	return
}
